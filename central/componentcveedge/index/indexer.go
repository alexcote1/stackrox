package index

import (
	"context"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

// Indexer is the component-cve edge indexer.
//
//go:generate mockgen-wrapper
type Indexer interface {
	AddComponentCVEEdge(componentcveedge *storage.ComponentCVEEdge) error
	AddComponentCVEEdges(componentcveedges []*storage.ComponentCVEEdge) error
	Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error)
	DeleteComponentCVEEdge(id string) error
	DeleteComponentCVEEdges(ids []string) error
	MarkInitialIndexingComplete() error
	NeedsInitialIndexing() (bool, error)
	Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error)
}
