
# Newsfeeder

Newsfeeder is a project to demonstrate REST API in Go using the Gin web framework.

## Running the project

`make dev`

## Testing

Run test from a particular package

`
go test ./platform/newsfeed`

Run test for entire project

`go test ./...`

Run test and check test coverage from the entire project

`go test -cover ./...`

####  **Environment variables and values**

```
DB_NAME=newsfeeder

DB_USERNAME=username

DB_PASSWORD=password

DB_HOST=localhost

DB_PORT=5432

DB_POOL_SIZE=10

BUILD_ENV=dev

SERVER_PORT=8080
```