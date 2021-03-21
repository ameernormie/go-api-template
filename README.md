# Backend-go-template

### Installation

1. Make sure you have [**go**](https://golang.org/doc/install) and [**docker**](https://www.docker.com/products/docker-desktop) installed.
2. Install dependencies.

```
$ go mod vendor
```

3. running docker container

```
$ docker-compose up -d
```

4. Run test

```
$ make test
```

or

```
$ go test ./internal/routes -run YourCaseName # run single case
```

5. Run server

```
$ make start
```
