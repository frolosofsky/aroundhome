# Aroundhome assignment

## TL;DR

```shell
docker compose up -d # to build and run in Docker
make test            # to run unit and behave tests
make                 # to build native executable

# to run locally
dbconn='postgres://postgres:pass@localhost/aroundhome?sslmode=disable' bin/matcher
```

Edit [compose.yaml](compose.yaml) to seed the database. The data is identical to the one used in behavior tests (see below).

# Content

* [features/customers-and-partners.feature](./features/customers-and-partners.feature) feature (behavior) tests.
  Running `make test-behave` will create python virtual environment, install dependencies, and execute `behave`. **Warning**
  virtual environment is by no mean a security sandbox.
* [api/v1/openapi.yaml](./api/v1/openapi.yaml) OpenAPI specs.
* [sql/master.sql](./sql/master.sql) database schema (postgres).
* [compose.yaml](./compose.yaml) docker compose.
* [cmd/matcher/main.go](./cmd/matcher/main.go) service implementation entrypoint.

# Testing

You can edit `.feature` file to play with different testing scenarios. In this case, tests' backend will create all necessary data
in the database. If you prefer manually test http endpoints, refer to open api specs and make sure you seeded the database:
* either edit `compose.yaml` and (re)run `docker compose up`
* or `psql postgres://postgres:pass@localhost/aroundhome -f sql/dump.sql`
* or enter arbitrary data manually.

# Implementation details

## Database

I used PostGIS extension to the postgres database to effeciently store and access geography data. The service logic is quite trivial,
it translates http requests in SQL queries and back SQL data to JSON. Still, the service is built in extendable manner:
* [cmd/matcher/api](./cmd/matcher/api) http handlers and types
* [pkg/store](./pkg/store) data store

You can find actual SQL queries in [pkg/store/driver/pg/pg.go](./pkg/store/driver/pg/pg.go).

## Configuration

Service understands two environment variables: `dbconn` (required) and `bind` (defaults to `127.0.0.1:8080`).

# Flaws / ideas to improve

1. Implement authentication and authorization.
2. Follow the specification-first approach, generate backend code from openapi.yaml.
3. Use more high-level http framework, likely in conjunction with openapi generator.
4. Observability.
5. When running behave-tests, create temporary database and run api service on random port.
6. Behave tests should not go to the service database (Blackbox testing principle). This requires more administrative http endpoints.
7. CI.
