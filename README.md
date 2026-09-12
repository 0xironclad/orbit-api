# Orbit API

Core backend services for the Orbit. 

## Getting started

### Prerequisites

- Go 1.26+
- [direnv](https://direnv.net/)

### Configure the environment

Create a local `.envrc` file and set the server address:

```sh
export ADDRESS=":3000"
```

Allow direnv to use it:

```sh
direnv allow
```

### Run the API

Start the development server with live reload:

```sh
direnv exec . air
```

The API will be available at `http://localhost:3000`. Check that it is running at:

```sh
curl http://localhost:3000/v1/health
```

### Build

```sh
go build -o bin/api ./cmd/api
```
