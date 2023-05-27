# Title: ZUS

// copilot: describe the project here

## Prerequisites

* Go 1.9+
* PostgreSQL 9.6+
* GNU Make 3.81+

## Local installation

```
make deps
make setupdb
make build
```

## Configuration

Configuration:
```
cp config.yml.example config.yml
```

Database configuration:
```
make setupdb
```

## Validation

Testing:
```
make test
```

## Running

Locally:
```
make build
./build/zus
```
