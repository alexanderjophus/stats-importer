# Stats Importer

This project intends to be an entrypoint to a pachyderm repo for the NHL stats api.

Goals;

- Read in relative link
- Call nhl stats api
- Write out to statsapi pachyderm repo

## Prerequisites

Installed;

- Go
- pachctl

## Usage

Run;

- `make deploy-pachyderm` to deploy pachyderm on a local k8s cluster. Check pachd daemon is running before continuing
- `export ADDRESS=$(minikube ip):30650` (if using minikube), address is needed for pachyderm
- `pachctl port-forward &` to use pachyderm
- `make create-repo` to create the repo we'll be using
- `make run` this will build and start the service on port 8080

Once the API is up and running call the service with a post with a body similar to;

```json
{
  "link": "/api/v1/game/2018020002/feed/live"
}
```

For other examples of links to use try searching some [unofficial](https://gitlab.com/dword4/nhlapi/blob/master/stats-api.md) docs
