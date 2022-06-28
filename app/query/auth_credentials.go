package query

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type BigQueryAuth struct {
	Client  *bigquery.Client
	Context context.Context
}
