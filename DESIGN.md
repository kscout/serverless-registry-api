# Design
API design.

# Table Of Contents
- [Overview](#overview)
- [Data Model](#data-model)
- [Endpoints](#endpoints)

# Overview
HTTP RESTful API.  

Requests pass data via JSON encoded bodies except for in GET requests where data
will be passed via URL and query parameters.

Responses will always return JSON.

# Data Model
## App Model
`apps` collection.

Schema:

- `app_id` (String)
- `name` (String)
- `tagline` (String)
- `description` (String)
- `screenshot_urls` (List[String])
- `logo_url` (String)
- `tags` (List[String])
- `verification_status` (String): One of `pending`, `verifying`, `good`, `bad`
- `github_url` (String)
- `version` (String)
- `author` (String)
- `maintainer` (String)

# Endpoints
Most endpoints to not require authentication.  

Those which do will be marked. Provide authentication as a bearer token in the
`Authorization` header.  

Endpoints which specify a response of `None` will return the 
JSON: `{"ok": true}`.

## App Endpoints
### Search Apps
`GET /apps?query=<query>&tags=<tags>&categories=<categories>`

Search serverless apps in hub.

If no search parameters are provided all applications will be returned.

Request:

- `query` (Optional, String): Natural language query
- `tags` (Optional, List[String]): Tags applications must have
- `categories` (Optional, List[String]): Categories applications must be part of

Response:

- `apps` (List[[App Model](#app-model)])

### Get App By ID
`GET /apps/<app_id>`

Get application by ID.

Request:

- `app_id` (String)

Response:

- `app` ([App Model](#app-model))

### App Pull Request Webhook Endpoint
`POST /apps/webhook`

GitHub will make a request to this endpoint every time a new pull request is 
made to submit an app.

Request:

- [GitHub Webhook Request](https://developer.github.com/webhooks/#payloads)
- [`PullRequestEvent'](https://developer.github.com/v3/activity/events/types/#pullrequestevent)

Response: None

### Search Tags
`GET /apps/tags?query=<query>`

Get all available tags.

Request:

- `query` (Optional, String): Search string, if empty all tags will be returned

Response:

- `tags` (List[String])

### Search Categories
`GET /apps/categories?query=<query>`

Get all available categories.

Request:

- `query` (Optional, String): Search string, if empty all categories will 
  be returned

Response:

- `categories` (List[String])

## User Endpoints
`GET /users/login`

Login via OpenShift.

Request: None

Response:

- `authentication_token` (String): Use this to authenticate with the App API in
  the future

## Cluster Endpoints
### List Clusters
`GET /clusters`

Lists available OpenShift clusters.

Authentication required.

Request: None

Response:

- `clusters` (List[Object]): Objects with keys
  - `id` (String)
  - `name` (String)

### Deploy To Cluster
`POST /clusters/<id>/deploy?app_id=<app_id>`

Deploy application to OpenShift cluster.

Request:

- `id` (String): ID of OpenShift cluster
- `app_id` (String): ID of app to deploy

Response: None

### Get Deploy Instructions
`GET /clusters/deploy_instructions?app_id=<app_id>`

Get manual deploy instructions for app.

Request:

- `app_id` (String): ID of app to return instructions for

Response:

- `instructions` (String)

## Meta Endpoints
### Health Check
`GET /health`

Used to determine if server is operating fully.

Request: None

Response: None