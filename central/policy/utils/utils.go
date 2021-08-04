package utils

import (
	mitreDS "github.com/stackrox/rox/central/mitre/datastore"
	"github.com/stackrox/rox/generated/storage"
)

// GetFullMitreAttackVectors returns MITRE ATT&CK for policy with full data.
func GetFullMitreAttackVectors(mitreStore mitreDS.MitreAttackReadOnlyDataStore, policy *storage.Policy) ([]*storage.MitreAttackVector, error) {
	if mitreStore == nil || policy == nil {
		return nil, nil
	}

	vectorsAsMap := make(map[string]map[string]struct{})
	for _, vector := range policy.GetMitreAttackVectors() {
		tacticID := vector.GetTactic()
		vectorsAsMap[tacticID] = make(map[string]struct{})
		for _, techniqueID := range vector.GetTechniques() {
			vectorsAsMap[tacticID][techniqueID] = struct{}{}
		}
	}

	resp := make([]*storage.MitreAttackVector, 0, len(vectorsAsMap))
	for tacticID, techniqueIDs := range vectorsAsMap {
		fullVector, err := mitreStore.Get(tacticID)
		if err != nil {
			return nil, err
		}

		vector := &storage.MitreAttackVector{
			Tactic: fullVector.GetTactic(),
		}
		for _, technique := range fullVector.GetTechniques() {
			if _, ok := techniqueIDs[technique.GetId()]; ok {
				vector.Techniques = append(vector.Techniques, technique)
			}
		}
		resp = append(resp, vector)
	}
	return resp, nil
}
