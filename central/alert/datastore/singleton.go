package datastore

import (
	"sync"

	"github.com/stackrox/rox/central/alert/index"
	"github.com/stackrox/rox/central/alert/search"
	"github.com/stackrox/rox/central/alert/store"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/globalindex"
)

var (
	once sync.Once

	indexer  index.Indexer
	storage  store.Store
	searcher search.Searcher

	ad DataStore
)

func initialize() {
	storage = store.New(globaldb.GetGlobalDB())
	indexer = index.New(globalindex.GetGlobalIndex())

	var err error
	searcher, err = search.New(storage, indexer)
	if err != nil {
		panic("unable to load search index for alerts")
	}

	ad = New(storage, indexer, searcher)
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return ad
}
