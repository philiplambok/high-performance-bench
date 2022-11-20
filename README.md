## On High-Performance API Application

Sample scripts to compare speed between stacks.

Currently covered:
- Rails 7 (Ruby 3.1.2)
- Roda (Ruby 3.1.2)
- Go with Gin Framework (Go 1.18.4)

Database:
- Mysql 14.14

## Setup DB && Running the API servers

Setup the DB with: 

```sh
$ make install
```

Running the Rails server (listen at http:3000)

```sh
$ make rails-server
```

Running the Go server (listen at http:8080)

```sh
$ make go-server
```

Running the Roda server (listen at http:9292)

```sh
$ make roda-server
```

Running the benchmarks:

```sh
$ R=123 make on-rails
$ R=123 make on-go
$ R=123 make on-roda
```

`R` is a total request.

## Some benchmarks

Stack|Total Request|Average per request|Total time|
---------|----------|---------|-------|
Rails API | 100 | 13.6605 ms|1.366055 s|
Go API | 100 | 4.3355 ms | .433559 s|
Roda API | 100 | 3.1006 ms | .310069 s|
Rails API | 1000 | 13.9347 ms | 13.934720 s|
Go API | 1000 | 3.168585 ms | 3.168585 s|
Roda APi | 1000 | 3.1366 ms | 3.136605 s |
