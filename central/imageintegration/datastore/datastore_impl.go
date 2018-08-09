package datastore

import (
	"github.com/stackrox/rox/central/imageintegration/store"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/set"
)

type datastoreImpl struct {
	storage store.Store
}

// GetImageIntegration is pass-through to the underlying store.
func (ds *datastoreImpl) GetImageIntegration(id string) (*v1.ImageIntegration, bool, error) {
	return ds.storage.GetImageIntegration(id)
}

// GetImageIntegrations provides an in memory layer on top of the underlying DB based storage.
func (ds *datastoreImpl) GetImageIntegrations(request *v1.GetImageIntegrationsRequest) ([]*v1.ImageIntegration, error) {
	integrations, err := ds.storage.GetImageIntegrations(request)
	if err != nil {
		return nil, err
	}

	integrationSlice := integrations[:0]
	for _, integration := range integrations {
		clusterSet := set.NewSetFromStringSlice(integration.GetClusters())
		if len(request.GetCluster()) != 0 && !clusterSet.Contains(request.GetCluster()) {
			continue
		}
		if request.GetName() != "" && request.GetName() != integration.GetName() {
			continue
		}
		integrationSlice = append(integrationSlice, integration)
	}
	return integrationSlice, nil
}

// AddImageIntegration is pass-through to the underlying store.
func (ds *datastoreImpl) AddImageIntegration(integration *v1.ImageIntegration) (string, error) {
	return ds.storage.AddImageIntegration(integration)
}

// UpdateImageIntegration is pass-through to the underlying store.
func (ds *datastoreImpl) UpdateImageIntegration(integration *v1.ImageIntegration) error {
	return ds.storage.UpdateImageIntegration(integration)
}

// RemoveImageIntegration is pass-through to the underlying store.
func (ds *datastoreImpl) RemoveImageIntegration(id string) error {
	return ds.storage.RemoveImageIntegration(id)
}
