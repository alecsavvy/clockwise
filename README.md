# clockwise

Clockwise is an entity manager implementation intended to upgrade the acdc network to something more lightweight and secure.

## running

To run the local cluster on your machine with a running loadtest, all you need to do is run the make file. Make sure to have docker running and go installed.

```
make
```

## architecture

### core

Where the embedded chain node is ran from as well as the abci for said node is implemented.

### pubsub

A generic pubsub module that uses channels to publish to multiple subscribers.

### api

Generic rest apis for clockwise, this where a `/relay` endpoint would live. Also has a

### graph

GraphQL interface for clockwise. This was initially implemented to avoid needing a frontend to test. To use go to `http://{node}:{port}/graphiql`

### loadtest (moshpit)

Separate binary that uses the graphql client to send a constant stream of entity manager events to nodes. Also known as mosh pit. Has a UI located at `http://localhost:8080`.

### infra

Docker compose, dockerfiles, and other crap to stand up the cluster locally. Also builds moshpit into it's own image.

### sdk

External go client that's used by moshpit but could also be used by external projects.
