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

Running the Go server (listen at http:4000)

```sh
$ make go-server
```

Running the Roda server (listen at http:9292)

```sh
$ make roda-server
```

## Some benchmarks

|Stack|Total Request|Average per request|Total time|
|---|---|---|---|---|
|Rails API|100||13.6605 ms|1.366055 s|

