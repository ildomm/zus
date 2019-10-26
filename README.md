
## Installation

Database configuration:
```
CREATE USER zus;
ALTER USER zus WITH ENCRYPTED PASSWORD 'y8364mry5w';
CREATE DATABASE "zus";
GRANT ALL PRIVILEGES ON DATABASE zus TO zus;

/use zus;
CREATE EXTENSION "uuid-ossp";

Install dependencies:

```