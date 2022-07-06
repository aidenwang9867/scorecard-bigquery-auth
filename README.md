# Scorecard-BigQuery-Auth Web App
This is the initial version of the REST API server to be deployed to the GKE cluster.
It can be pulled to local and run on `127.0.0.1:8080`.

Currently supported routes (APIs) include:
1. `GET` `/query/dependencies`?system=`{pkg_system_name}`&name=`{pkg_name}`&version=`{pkg_version}`

2. `POST` `/query/dependencies`, with a request body of type `application/json`:
  ```json
  {
  "system": "pkg_system_name",
  "name": "pkg_name",
  "version": "pkg_version"
  }
  ```
3. `GET` `/query/vulnerabilities`?system=`{pkg_system_name}`&name=`{pkg_name}`&version=`{pkg_version}`

4. `POST` `/query/vulnerabilities`, with a request body of type `application/json`:
  ```json
  {
  "system": "pkg_system_name",
  "name": "pkg_name",
  "version": "pkg_version"
  }
  ```
  
5. `POST` `/query/arbitrary`, with a request body of type `text/plain`:
  ```
  SELECT some_thing FROM some_table WHERE... (put any SQL queries here)
  ```
