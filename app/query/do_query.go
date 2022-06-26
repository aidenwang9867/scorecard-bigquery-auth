package query

import (
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func DoQueryAndGetRowIterator(auth BigQueryAuth, queryStr string) (*bigquery.RowIterator, error) {
	client, context := auth.Client, auth.Context
	q := client.Query(queryStr)
	// Execute the query and return the result row iterator.
	iter, err := q.Read(context)
	if err != nil {
		return nil, fmt.Errorf("%w when reading the context", err)
	}
	return iter, nil
}

// GetVulnerabilityByAdvID now is only used for supplementing the vuln result obtained
// from the GitHub API with vuln data retrieved from BQ.
func GetVulnerabilityByAdvID(auth BigQueryAuth, advID string) (Vulnerability, error) {
	it, err := DoQueryAndGetRowIterator(
		auth,
		fmt.Sprintf(
			QueryVulnerabilityByAdvID,
			advID,
		),
	)
	if err != nil {
		return Vulnerability{}, err
	}
	vuln := Vulnerability{}
	err = it.Next(&vuln)
	if err != nil {
		return Vulnerability{}, err
	}
	return vuln, nil
}

func GetVulnerabilitiesBySystemNameVersion(
	auth BigQueryAuth, system string, name string, version string,
) ([]Vulnerability, error) {
	it, err := DoQueryAndGetRowIterator(
		auth,
		fmt.Sprintf(
			QueryVulnerabilitiesBySystemNameVersion,
			strings.ToUpper(system),
			name,
			version,
		),
	)
	if err != nil {
		return nil, err
	}
	vuln := []Vulnerability{}
	for {
		row := Vulnerability{}
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		vuln = append(vuln, row)
	}
	return vuln, nil
}

func GetDependenciesBySystemNameVersion(
	auth BigQueryAuth, system string, name string, version string,
) ([]Dependency, error) {
	it, err := DoQueryAndGetRowIterator(
		auth,
		fmt.Sprintf(
			QueryDependencies,
			strings.ToUpper(system),
			name,
			version,
		),
	)
	if err != nil {
		return nil, err
	}
	dep := []Dependency{}
	for {
		row := Dependency{}
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		dep = append(dep, row)
	}
	return dep, nil
}

func GetResultsByArbitraryQuery(auth BigQueryAuth, q string) (string, error) {
	it, err := DoQueryAndGetRowIterator(
		auth,
		q,
	)
	if err != nil {
		return "", err
	}
	results := ""
	for {
		values := []bigquery.Value{}
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		results = fmt.Sprintf(results+"\n"+"%v", fmt.Sprintf("%v", values))
	}
	return results, nil
}
