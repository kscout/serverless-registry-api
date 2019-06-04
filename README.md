# App API
API which manages applications.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)
- [Deployment](#deployment)

# Overview
See [DESIGN.md](DESIGN.md)

# Development
## Database
Start a local MongoDB server by running:

```
make db
```

## Run
Start the server by running:

```
go run .
```

## Temporary Open Shift
The `tmpk` script wraps `kubectl` with the required arguments to connect to the
48 hour Open Shift clusters.

Set the `TMPK_TOKEN` and `TMPK_N` environment variables. See the `tmpk` file 
for details about what to set these environment variables to.

Use the `tmpk` script as if it was `kubectl`:

```
./tmpk get all
```

# Deployment
Configuration is passed via environment variables:

- `APP_HTTP_ADDR` (String): Address to bind server, defaults to `:5000`
- `APP_DB_CONN_URL` (String): MongoDB connection URI, defaults to connecting to
  the local database started by `make db`

A webhook should exist for the
[app-repository](https://github.com/knative-scout/app-repository/settings/hooks/new).  
This webhook should send pull request events to the app pull request 
webhook endpoint.
