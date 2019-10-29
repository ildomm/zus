
## Installation

Database configuration:
```
CREATE USER zus;
ALTER USER zus WITH ENCRYPTED PASSWORD 'y8364mry5w';
CREATE DATABASE "zus";
GRANT ALL PRIVILEGES ON DATABASE zus TO zus;

/use zus;
CREATE EXTENSION "uuid-ossp";
```

## Validation

Testing:
```
run_tests.*
```

## Running

Locally:
```
go build cmd\zus-server\main.go
```

Deploying:
```
deploy.sh
```