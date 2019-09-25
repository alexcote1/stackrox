// Entity Relationships

// note: the relationships are directional!
// changing direction may change relationship type between entities!!

import entityTypes from 'constants/entityTypes';
import relationshipTypes from 'constants/relationshipTypes';
import useCaseTypes from 'constants/useCaseTypes';
import { uniq } from 'lodash';

// base k8s entities to be used across all use cases
const baseEntities = [entityTypes.CLUSTER, entityTypes.NAMESPACE, entityTypes.DEPLOYMENT];

// map of use cases to entities
const useCaseEntityMap = {
    [useCaseTypes.COMPLIANCE]: [entityTypes.CONTROL, entityTypes.NODE, ...baseEntities],
    [useCaseTypes.CONFIG_MANAGEMENT]: [
        entityTypes.CONTROL,
        entityTypes.NODE,
        entityTypes.IMAGE,
        entityTypes.ROLE,
        entityTypes.SECRET,
        entityTypes.SUBJECT,
        entityTypes.SERVICE_ACCOUNT,
        entityTypes.POLICY,
        ...baseEntities
    ],
    [useCaseTypes.VULN_MANAGEMENT]: [
        entityTypes.POLICY,
        entityTypes.IMAGE,
        entityTypes.COMPONENT,
        entityTypes.CVE,
        ...baseEntities
    ]
};

// const edgeTypes = {
//     VIOLATIONS: 'VIOLATIONS',
//     EVIDENCE: 'EVIDENCE'
// };

// map of edge types (side effects of when two entities cross)
// note: these will not be listed -- they should only show up as columns in `x findings` tables
// const relationshipEdgeMap = {
//     [edgeTypes.VIOLATIONS]: {
//         entityType1: [entityTypes.POLICY],
//         entityType2: [entityTypes.DEPLOYMENT]
//     },
//     [edgeTypes.EVIDENCE]: {
//         entityType1: [entityTypes.CONTROL],
//         entityType2: [entityTypes.NODE, entityTypes.CLUSTER, entityTypes.DEPLOYMENT]
//     }
// };

const entityRelationshipMap = {
    [entityTypes.CLUSTER]: {
        children: [entityTypes.NODE, entityTypes.NAMESPACE, entityTypes.ROLE],
        parents: [],
        matches: [entityTypes.CONTROL]
    },
    [entityTypes.NODE]: {
        children: [],
        parents: [entityTypes.CLUSTER],
        matches: [entityTypes.CONTROL]
    },
    [entityTypes.NAMESPACE]: {
        children: [entityTypes.DEPLOYMENT, entityTypes.SERVICE_ACCOUNT, entityTypes.SECRET],
        parents: [entityTypes.CLUSTER],
        matches: []
    },
    [entityTypes.DEPLOYMENT]: {
        children: [entityTypes.IMAGE],
        parents: [entityTypes.NAMESPACE, entityTypes.CLUSTER],
        matches: [
            entityTypes.SERVICE_ACCOUNT,
            entityTypes.POLICY,
            entityTypes.CONTROL,
            entityTypes.SECRET
        ]
    },
    [entityTypes.IMAGE]: {
        children: [entityTypes.COMPONENT],
        parents: [],
        matches: [entityTypes.DEPLOYMENT]
    },
    [entityTypes.COMPONENT]: {
        children: [entityTypes.CVE],
        parents: [],
        matches: [entityTypes.IMAGE]
    },
    [entityTypes.CVE]: {
        children: [],
        parents: [],
        matches: [entityTypes.COMPONENT]
    },
    [entityTypes.CONTROL]: {
        children: [],
        parents: [],
        matches: [entityTypes.NODE, entityTypes.DEPLOYMENT, entityTypes.CLUSTER]
    },
    [entityTypes.POLICY]: {
        children: [],
        parents: [],
        matches: [entityTypes.DEPLOYMENT]
    },
    [entityTypes.SECRET]: {
        children: [],
        parents: [entityTypes.NAMESPACE],
        matches: [entityTypes.DEPLOYMENT]
    },
    [entityTypes.SUBJECT]: {
        children: [],
        parents: [],
        matches: [entityTypes.ROLE]
    },
    [entityTypes.SERVICE_ACCOUNT]: {
        children: [],
        parents: [entityTypes.NAMESPACE],
        matches: [entityTypes.DEPLOYMENT, entityTypes.ROLE]
    },
    [entityTypes.ROLE]: {
        children: [],
        parents: [entityTypes.CLUSTER],
        matches: [entityTypes.SERVICE_ACCOUNT, entityTypes.SUBJECT]
    }
};

// helper functions
const getChildren = entityType => entityRelationshipMap[entityType].children;
const getParents = entityType => entityRelationshipMap[entityType].parents;
const getMatches = entityType => entityRelationshipMap[entityType].matches;

// function to recursively get inclusive 'contains' relationships (inferred)
// this includes all generations of children AND inferred (matches of children down the chain) relationships
// e.g. namespace inclusively contains policy since ns contains deployment and deployment matches policy
const getContains = entityType => {
    const relationships = [];
    const children = getChildren(entityType);
    if (children) {
        children.forEach(child => {
            const childMatches = getMatches(child);
            const childContains = getContains(child);
            relationships.push(child, ...childMatches, ...childContains);
        });
    }
    return uniq(relationships);
};

const isChild = (parent, child) => !!getChildren(parent).find(c => c === child);
const isParent = (parent, child) => !!getParents(child).find(p => p === parent);
const isMatch = (entityType1, entityType2) =>
    !!getMatches(entityType1).find(m => m === entityType2);
const isContained = (entityType1, entityType2) =>
    !!getContains(entityType1).find(c => c === entityType2);

// wrapper function returns a list of entities, given an entitytype, relationship, and context
// e.g.
// f(type, relationship, context)
// f(cluster, contains, config management), f(deployment, parents, config management)
export const getEntityTypesByRelationship = (entityType, relationship, context) => {
    let entities = [];
    if (relationship === relationshipTypes.CONTAINS) {
        entities = getContains(entityType);
    } else if (relationship === relationshipTypes.MATCHES) {
        entities = getMatches(entityType);
    } else if (relationship === relationshipTypes.PARENTS) {
        entities = getParents(entityType);
    } else if (relationship === relationshipTypes.CHILDREN) {
        entities = getChildren(entityType);
    }
    return entities.filter(entity => useCaseEntityMap[context].includes(entity));
};

export default {
    getChildren,
    getParents,
    getMatches,
    getContains,
    isChild,
    isParent,
    isMatch,
    isContained
};
