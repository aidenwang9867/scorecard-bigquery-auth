package query

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

func Authenticate() (BigQueryAuth, error) {
	ctx := context.Background()
	cli, err := bigquery.NewClient(ctx, "ossf-malware-analysis")
	if err != nil {
		return BigQueryAuth{}, fmt.Errorf("%w when creating the big query client", err)
	}
	auth := BigQueryAuth{
		Client:  cli,
		Context: ctx,
	}
	return auth, nil
}
