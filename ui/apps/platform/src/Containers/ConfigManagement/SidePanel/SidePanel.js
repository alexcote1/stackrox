import React, { useContext } from 'react';
import PropTypes from 'prop-types';
import ReactRouterPropTypes from 'react-router-prop-types';
import { withRouter, Link } from 'react-router-dom';
import onClickOutside from 'react-onclickoutside';
import { ExternalLink as ExternalLinkIcon } from 'react-feather';

import Panel from 'Components/Panel';
import searchContext from 'Containers/searchContext';
import Entity from 'Containers/ConfigManagement/Entity';
import { useTheme } from 'Containers/ThemeProvider';
import workflowStateContext from 'Containers/workflowStateContext';
import parseURL from 'utils/URLParser';
import URLService from 'utils/URLService';
import BreadCrumbs from './BreadCrumbs';

const SidePanel = ({
    match,
    location,
    history,
    contextEntityType,
    contextEntityId,
    entityListType1,
    entityType1,
    entityId1,
    entityType2,
    entityListType2,
    entityId2,
    query,
}) => {
    const { isDarkMode } = useTheme();
    const workflowState = parseURL(location);
    const searchParam = useContext(searchContext);
    const isList = !entityId1 || (entityListType2 && !entityId2);

    function getCurrentEntityId() {
        return entityId2 || entityId1;
    }

    function getCurrentEntityType() {
        return (
            entityType2 ||
            (entityId2 && entityListType2) ||
            entityType1 ||
            entityListType1 ||
            contextEntityType
        );
    }

    function getListType() {
        if (!isList) {
            return null;
        }
        return entityListType2;
    }

    function getSearchParams() {
        return query[searchParam];
    }

    function onClose() {
        history.push(URLService.getURL(match, location).clearSidePanelParams().url());
    }

    SidePanel.handleClickOutside = () => {
        onClose();
    };

    const entityId = getCurrentEntityId();
    const entityType = getCurrentEntityType();
    const listType = getListType();
    const externalURL = URLService.getURL(match, location)
        .base(entityType, entityId)
        .push(listType)
        .query()
        .query(getSearchParams())
        .url();
    const externalLink = (
        <div className="flex items-center h-full hover:bg-base-300">
            <Link
                to={externalURL}
                data-testid="external-link"
                className={`${
                    !isDarkMode ? 'border-base-100' : 'border-base-400'
                } border-l h-full p-4`}
            >
                <ExternalLinkIcon className="h-6 w-6 text-base-600" />
            </Link>
        </div>
    );

    const entityContext = {};
    if (contextEntityType) {
        entityContext[contextEntityType] = contextEntityId;
    }
    if (entityId2) {
        entityContext[entityType1 || entityListType1] = entityId1;
    }
    return (
        <workflowStateContext.Provider value={workflowState}>
            <Panel
                id="side-panel"
                headerClassName={`flex w-full h-14 rounded-tl-lg overflow-y-hidden border-b ${
                    !isDarkMode
                        ? 'bg-side-panel-wave border-base-100'
                        : 'bg-primary-200 border-primary-400'
                }`}
                bodyClassName={isDarkMode ? 'bg-base-0' : 'bg-base-100'}
                headerTextComponent={
                    <BreadCrumbs
                        className="font-700 leading-normal text-base-600 tracking-wide truncate"
                        entityType1={entityType1 || entityListType1}
                        entityId1={entityId1}
                        entityType2={entityType2}
                        entityListType2={entityListType2}
                        entityId2={entityId2}
                    />
                }
                headerComponents={externalLink}
                onClose={onClose}
                closeButtonClassName={
                    isDarkMode ? 'border-l border-base-400' : 'border-l border-base-100'
                }
            >
                <Entity
                    entityContext={entityContext}
                    entityType={entityType}
                    entityId={entityId}
                    entityListType={listType}
                    query={query}
                />
            </Panel>
        </workflowStateContext.Provider>
    );
};

SidePanel.propTypes = {
    match: ReactRouterPropTypes.match.isRequired,
    location: ReactRouterPropTypes.location.isRequired,
    history: ReactRouterPropTypes.history.isRequired,
    contextEntityType: PropTypes.string,
    contextEntityId: PropTypes.string,
    entityType1: PropTypes.string,
    entityListType1: PropTypes.string,
    entityId1: PropTypes.string,
    entityType2: PropTypes.string,
    entityListType2: PropTypes.string,
    entityId2: PropTypes.string,
    query: PropTypes.shape().isRequired,
};

SidePanel.defaultProps = {
    contextEntityType: null,
    contextEntityId: null,
    entityType1: null,
    entityListType1: null,
    entityId1: null,
    entityType2: null,
    entityListType2: null,
    entityId2: null,
};

const clickOutsideConfig = {
    handleClickOutside: () => SidePanel.handleClickOutside,
};

/*
 * If more than one SidePanel is rendered, this Pure Functional Component will need to be converted to
 * a Class Component in order to work correctly. See https://github.com/stackrox/rox/pull/3090#pullrequestreview-274948849
 */
export default onClickOutside(withRouter(SidePanel), clickOutsideConfig);
