package datastore

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/cloudsources/datastore/internal/search"
	"github.com/stackrox/rox/central/cloudsources/datastore/internal/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/stringutils"
)

var _ DataStore = (*datastoreImpl)(nil)

type datastoreImpl struct {
	searcher search.Searcher
	store    store.Store
}

func (ds *datastoreImpl) CountCloudSources(ctx context.Context, query *v1.Query) (int, error) {
	count, err := ds.searcher.Count(ctx, query)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count cloud sources")
	}
	return count, nil
}

func (ds *datastoreImpl) GetCloudSource(ctx context.Context, id string) (*storage.CloudSource, error) {
	cloudSource, exists, err := ds.store.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cloud source")
	}
	if !exists {
		return nil, errox.NotFound.Newf("cloud source %q not found", id)
	}
	return cloudSource, nil
}

func (ds *datastoreImpl) ListCloudSources(ctx context.Context, query *v1.Query) ([]*storage.CloudSource, error) {
	cloudSources, err := ds.store.GetByQuery(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list cloud sources")
	}
	return cloudSources, nil
}

func (ds *datastoreImpl) UpsertCloudSource(ctx context.Context, cloudSource *storage.CloudSource) error {
	if err := validateCloudSource(cloudSource); err != nil {
		return errors.Wrap(errox.InvalidArgs, err.Error())
	}
	if err := ds.store.Upsert(ctx, cloudSource); err != nil {
		return errors.Wrapf(err, "failed to upsert cloud source %q", cloudSource.GetId())
	}
	return nil
}

func (ds *datastoreImpl) DeleteCloudSource(ctx context.Context, id string) error {
	if err := ds.store.Delete(ctx, id); err != nil {
		return errors.Wrapf(err, "failed to delete cloud source %q", id)
	}
	return nil
}

func validateCloudSource(cloudSource *storage.CloudSource) error {
	if cloudSource == nil {
		return errors.New("empty cloud source")
	}

	errorList := errorhelpers.NewErrorList("Validation")
	if stringutils.AtLeastOneEmpty(
		cloudSource.GetId(),
		cloudSource.GetName(),
		cloudSource.GetCredentials().GetSecret(),
	) {
		errorList.AddString("all required fields must be set")
	}
	if cloudSource.GetConfig() == nil {
		errorList.AddString("empty cloud source config")
	}
	return errorList.ToError()
}
