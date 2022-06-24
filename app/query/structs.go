package query

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type BigQueryAuth struct {
	client  *bigquery.Client
	context context.Context
}

// type QueryDependency struct {
// 	System  string
// 	Name    string
// 	Version string
// }
