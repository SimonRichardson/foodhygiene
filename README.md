# Food Hygiene

## Command Hygiene

 - [Getting started](#getting-started)
 - [Introduction](#introduction)
 - [Backend](#backend-cli)
 - [Frontend](#frontend-ui)
 - [Tests](#tests)
 - [Improvements](#improvements)

### Getting started

The food hygiene CLI expects a few prerequisites to be installed for it to be
able to build. A valid `go` installation (1.8+) and a valid `$GOPATH`.

The quickest way to get up and running is to do the following:

```
make install
cd dist

./hygiene query
```

Then go to the following url [localhost:8080](localhost:8080/ui/dist/) and
start checking out the various ratings.

### Introduction

The hygiene command code is split into two distinct parts, the backend CLI and
then the frontend web site.

### Backend CLI

The backend CLI is written in `go`, using little dependencies as possible,
because most of the `stdlib` can perform most if not all the tasks required.
Including routing!

The dependency management is managed by `glide` as this seems to be the most
stable and reliable at the time of writing.

The `main.go` file (entry point) can be found in `cmd/hygiene/main.go`. This
is the entry point to the CLI and the code found with in the folder creates
the dependencies for the server. It also provides a helpful usage API to help
with starting said server.

The `pkg` directory is where all the CLI dependencies (libraries) are. This
is the [current standard practice](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
in go.

Building the whole of the backend can be done with the following command:

```
make clean clean-ui build-ui build
```

If nothing has changed with in the ui and you're just rebuilding the backend CLI
then the following command should suffice:

```
make clean build
```

#### Service

With in the `pkg` directory you'll find a `service` module that wraps the
Ratings Food gov API, so that we can normalize the data that's coming from the
API. For example we don't need all the fields found in the API json schema, so
we can reduce parsing required. Also we can normalize the naming of the
API where lowercase and uppercase conflicts with itself, to make a cleaner API
for ourselves.

Testing the `service` package becomes a lot easier as we have a `Service`
abstraction (interface), that allows us to create a mock service, a cache
service (see [README.md](pkg/service/README.md)) and a real implementation of
the service to the gov API.

#### Query

The `query` module handles all the requests and responses to and from the
server, including handling potential errors (out of bounds, malformed payloads).

The resulting ratings are returned as floats, but are rounded to 2 decimal
places to help with readability. The tests in the `query` module are tested
against the `service` mock API.

#### UI

The `ui` module handles all the UI requests. We embed the UI inside of the
binary for two reasons:

 1. We can distribute the binary and it'll run almost anywhere without
  modification
 2. We don't have any potential issues with the Ajax requests (thinking CORS)

The `Makefile` handles all the embedding of the files.

#### Usage

There are a few defaults that can be useful when developing the backend CLI and
they can be seen in the help section:

```
./hygiene query -help
USAGE
  query [flags]

FLAGS
  -api tcp://0.0.0.0:8080  listen address for ingest and store APIs
  -cache true              use cached results for better responsiveness
  -debug false             debug logging
  -ui.local false          Ignores embedded files and goes straight to the filesystem
```

### Frontend UI

The frontend is written in Javascript, using reactjs as the UI rendering agent.
I decided to use reactjs for the rendering as it's efficient at building this
sort of Ajax requests and render loop.

The code for the UI is all handled with in the `ui` folder and isn't dependant
on anything from outside of it.

To install all the packages simply type (replace `npm` with `yarn` if you don't
have it installed):

```
yarn install
```

This should then allow you to run `npm start` or `npm test`. The components
with in the UI are really simple, with the `App.jsx` handling all the Ajax
requests and the `Dropdown.jsx` and `Table.jsx` just rendering when required.

### Tests

Most of the application is tested to some degree, either via built in stdlib
testing, quick check testing (fuzz testing), mock testing and then finally
manual testing.

Of course more could be done, but with in the timeframe I personally think it's
a promising start.

#### Backend CLI Tests

Because we're using `glide` - it offers us a quick helpful command to test
everything:

```
go test -v $(glide nv)
```

#### Frontend Tests

NPM also provides us with some quick helpful command to test components

```
cd ui
npm test
```

### Improvements

 1. The first page could be rendered server side, then all subsequent requests
 could be ajax.
 2. The cache could use something like redis, memcached so we can have better
 management of the data.
