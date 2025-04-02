# Graphql golang


```mermaid
flowchart TD
    %% Client Layer
    subgraph "Client Layer"
        Client("Client"):::client
    end

    %% API & Middleware Layer
    subgraph "API & Middleware Layer"
        API("GraphQL API Server"):::api
        Middleware("Middleware (CORS,DataLoader,Auth)"):::middleware
    end

    %% GraphQL Layer
    subgraph "GraphQL Layer"
        GraphQL("GraphQL Layer"):::graphql
    end

    %% Domain Layer
    subgraph "Domain Layer"
        Domain("Domain (Account & User)"):::domain
    end

    %% Data Access Layer
    subgraph "Data Access Layer"
        DataAccess("Data Access (sqlc Repos)"):::data
        Migrations("Migration Scripts"):::data
        DB("(PostgreSQL Database)"):::database
    end

    %% Supporting Services
    subgraph "Supporting Services"
        Observability("Observability (Logger/Tracer)"):::support
        Redis("Redis Cache"):::redis
    end

    %% Connections
    Client -->|"sendsRequest"| API
    API -->|"invokes"| Middleware
    API -->|"routesTo"| GraphQL
    Middleware -->|"processes"| GraphQL
    GraphQL -->|"resolves"| Domain
    Domain -->|"calls"| DataAccess
    DataAccess -->|"executes"| DB
    Migrations -->|"appliedTo"| DB
    API -->|"logs"| Observability
    Domain -->|"caches"| Redis

    %% Click Events
    click API "https://github.com/alextanhongpin/go-graphql-template/tree/master/cmd/server"
    click GraphQL "https://github.com/alextanhongpin/go-graphql-template/tree/master/external/graph"
    click Domain "https://github.com/alextanhongpin/go-graphql-template/tree/master/domain"
    click DataAccess "https://github.com/alextanhongpin/go-graphql-template/tree/master/internal/postgres/sql"
    click Migrations "https://github.com/alextanhongpin/go-graphql-template/tree/master/internal/postgres/migrations"
    click Middleware "https://github.com/alextanhongpin/go-graphql-template/tree/master/external/middleware"
    click Observability "https://github.com/alextanhongpin/go-graphql-template/blob/master/internal/logger.go"
    click Redis "https://github.com/alextanhongpin/go-graphql-template/blob/master/internal/redis.go"

    %% Styles
    classDef client fill:#AED6F1,stroke:#1F618D,stroke-width:2px;
    classDef api fill:#FAD7A0,stroke:#D35400,stroke-width:2px;
    classDef middleware fill:#F9E79F,stroke:#B7950B,stroke-width:2px;
    classDef graphql fill:#A9DFBF,stroke:#27AE60,stroke-width:2px;
    classDef domain fill:#F5B7B1,stroke:#C0392B,stroke-width:2px;
    classDef data fill:#D2B4DE,stroke:#8E44AD,stroke-width:2px;
    classDef database fill:#F9E79F,stroke:#D4AC0D,stroke-width:2px;
    classDef support fill:#AED6F1,stroke:#1ABC9C,stroke-width:2px;
    classDef redis fill:#F4D03F,stroke:#D68910,stroke-width:2px;
```

## Middleware
- [x] CORS
- [x] DataLoader - solving the N+1 problem
- [x] Authorization - validates JWT tokens

## Observability
- [x] Logging - logging with correlation id, and also ease of integration with Jaeger for tracing

## Reliability
- [x] Graceful shutdown - signals termination to running processes (Database connections, redis connections, background worker etc) before shutting down the server
- [x] Retry database connection - attempts to reconnect to database if fails

## Database
- [x] Transactions - operations with transactions possible
- [x] Code generation - using sqlc package to generate repository layer


## Tests Assertions
- using goconvey and testify suite/assertions

https://github.com/smartystreets/goconvey/wiki/Assertions


## Pre-commit
To enable, change the `.git/hooks/pre-commit`:
```
#!/bin/sh

sh scripts/pre-commit.sh
[ $? -ne 0 ] && exit 1;
exit 0;
```
