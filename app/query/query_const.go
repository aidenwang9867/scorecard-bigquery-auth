package query

const (
	QueryVulnerabilitiesBySystemNameVersion = `
		SELECT
			pv.System, pv.Name, pv.Version,
			adv.Source, adv.SourceID, adv.SourceURL, adv.Title, adv.Description,
			CAST(adv.CVSS3Score AS FLOAT64) AS Score, adv.Severity, adv.GitHubSeverity, adv.Disclosed,
			adv.ReferenceURLs
		FROM (
			SELECT 
				System, Name, Version, SourceID
			FROM
				bigquery-public-data.deps_dev_v1.PackageVersions
			INNER JOIN
				UNNEST(Advisories)
			WHERE
				System = "%s" -- Input 1: dependency system
			AND
				Name = "%s" -- Input 2: dependency name
			AND
				Version = "%s" -- Input 3: dependency version
			AND
				SnapshotAt=(SELECT Time FROM bigquery-public-data.deps_dev_v1.Snapshots ORDER BY Time DESC LIMIT 1)
		) AS pv
		INNER JOIN 
			bigquery-public-data.deps_dev_v1.Advisories AS adv
		ON
		pv.SourceID = adv.SourceID
		WHERE
			SnapshotAt=(SELECT Time FROM bigquery-public-data.deps_dev_v1.Snapshots ORDER BY Time DESC LIMIT 1)
		ORDER BY Score
		DESC
		;
	`

	QueryDependencies = `
		SELECT
			Dependency.System, Dependency.Name, Dependency.Version
		FROM
			bigquery-public-data.deps_dev_v1.Dependencies
		WHERE
			System = "%s" -- Input 1: dependency system
		AND
			Name = "%s" -- Input 2: dependency name
		AND
			Version = "%s" -- Input 3: dependency version
		AND
			SnapshotAt=(SELECT Time FROM bigquery-public-data.deps_dev_v1.Snapshots ORDER BY Time DESC LIMIT 1)
		;
	`

	QueryVulnerabilityByAdvID = `
		SELECT 
			adv.Source, adv.SourceID, adv.SourceURL, adv.Title, adv.Description,
			CAST(adv.CVSS3Score AS FLOAT64) AS Score, adv.Severity, adv.GitHubSeverity, adv.Disclosed,
			adv.ReferenceURLs
		FROM
			bigquery-public-data.deps_dev_v1.Advisories AS adv
		WHERE
			adv.SourceID = "%s"
		AND
			SnapshotAt=(SELECT Time FROM bigquery-public-data.deps_dev_v1.Snapshots ORDER BY Time DESC LIMIT 1)
		;
	`
)
