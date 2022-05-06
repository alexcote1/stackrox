// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableSimpleaccessscopesStmt holds the create statement for table `simpleaccessscopes`.
	CreateTableSimpleaccessscopesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists simpleaccessscopes (
                   Id varchar,
                   Name varchar UNIQUE,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}

	// SimpleaccessscopesSchema is the go schema for table `simpleaccessscopes`.
	SimpleaccessscopesSchema = func() *walker.Schema {
		schema := globaldb.GetSchemaForTable("simpleaccessscopes")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.SimpleAccessScope)(nil)), "simpleaccessscopes")
		globaldb.RegisterTable(schema)
		return schema
	}()
)
