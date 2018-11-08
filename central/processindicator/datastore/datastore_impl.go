package datastore

import (
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/stackrox/rox/central/processindicator/index"
	"github.com/stackrox/rox/central/processindicator/search"
	"github.com/stackrox/rox/central/processindicator/store"
	"github.com/stackrox/rox/generated/api/v1"
	pkgSearch "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/set"
)

type datastoreImpl struct {
	storage  store.Store
	indexer  index.Indexer
	searcher search.Searcher
}

func (ds *datastoreImpl) Search(q *v1.Query) ([]pkgSearch.Result, error) {
	return ds.searcher.Search(q)
}

func (ds *datastoreImpl) SearchRawProcessIndicators(q *v1.Query) ([]*v1.ProcessIndicator, error) {
	return ds.searcher.SearchRawProcessIndicators(q)
}

func (ds *datastoreImpl) GetProcessIndicator(id string) (*v1.ProcessIndicator, bool, error) {
	return ds.storage.GetProcessIndicator(id)
}

func (ds *datastoreImpl) GetProcessIndicators() ([]*v1.ProcessIndicator, error) {
	return ds.storage.GetProcessIndicators()
}

func (ds *datastoreImpl) AddProcessIndicators(indicators ...*v1.ProcessIndicator) error {
	removedIndicators, err := ds.storage.AddProcessIndicators(indicators...)
	if err != nil {
		return err
	}

	// If there are no indicators to remove, short-circuit the rest of the code path.
	if len(removedIndicators) == 0 {
		return ds.indexer.AddProcessIndicators(indicators)
	}

	removedIndicatorsSet := set.NewStringSet(removedIndicators...)

	// We want to filter out indicators in the current batch which were dropped.
	filteredIndicators := indicators[:0]
	for _, indicator := range indicators {
		if removedIndicatorsSet.Contains(indicator.GetId()) {
			removedIndicatorsSet.Remove(indicator.GetId())
			continue
		}
		filteredIndicators = append(filteredIndicators, indicator)
	}

	// This removes indicators that previously existed in the index.
	if removedIndicatorsSet.Cardinality() > 0 {
		if err := ds.indexer.DeleteProcessIndicators(removedIndicatorsSet.AsSlice()...); err != nil {
			return err
		}
	}
	return ds.indexer.AddProcessIndicators(filteredIndicators)
}

func (ds *datastoreImpl) AddProcessIndicator(i *v1.ProcessIndicator) error {
	removedIndicator, err := ds.storage.AddProcessIndicator(i)
	if err != nil {
		return fmt.Errorf("adding indicator to bolt: %s", err)
	}
	if removedIndicator != "" {
		if err := ds.indexer.DeleteProcessIndicator(removedIndicator); err != nil {
			return fmt.Errorf("removing process indicator: %s", err)
		}
	}
	if err := ds.indexer.AddProcessIndicator(i); err != nil {
		return fmt.Errorf("adding indicator to index: %s", err)
	}
	return nil
}

func (ds *datastoreImpl) RemoveProcessIndicator(id string) error {
	if err := ds.storage.RemoveProcessIndicator(id); err != nil {
		return err
	}
	return ds.indexer.DeleteProcessIndicator(id)
}

func (ds *datastoreImpl) removeMatchingIndicators(results []pkgSearch.Result) error {
	idsToDelete := make([]string, 0, len(results))
	for _, r := range results {
		idsToDelete = append(idsToDelete, r.ID)
	}

	for _, id := range idsToDelete {
		if err := ds.storage.RemoveProcessIndicator(id); err != nil {
			log.Warnf("Failed to remove process indicator %q: %v", id, err)
		}
	}
	return ds.indexer.DeleteProcessIndicators(idsToDelete...)
}

func (ds *datastoreImpl) RemoveProcessIndicatorsByDeployment(id string) error {
	q := pkgSearch.NewQueryBuilder().AddExactMatches(pkgSearch.DeploymentID, id).ProtoQuery()
	results, err := ds.Search(q)
	if err != nil {
		return err
	}
	return ds.removeMatchingIndicators(results)
}

func (ds *datastoreImpl) RemoveProcessIndicatorsOfStaleContainers(deploymentID string, currentContainerIDs []string) error {
	queries := make([]*v1.Query, 0, len(currentContainerIDs)+1)
	queries = append(queries, pkgSearch.NewQueryBuilder().AddExactMatches(pkgSearch.DeploymentID, deploymentID).ProtoQuery())

	for _, containerID := range currentContainerIDs {
		queries = append(queries, pkgSearch.NewQueryBuilder().AddStrings(pkgSearch.ContainerID, pkgSearch.NegateQueryString(containerID)).ProtoQuery())
	}

	results, err := ds.Search(pkgSearch.ConjunctionQuery(queries...))
	if err != nil {
		return err
	}
	return ds.removeMatchingIndicators(results)
}
