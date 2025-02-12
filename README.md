# Stoic-Go: Overview and Guide  
  
## Overview  
  
Stoic-Go is a modular project designed for handling backend services, database migrations, and runtime utility tasks. It includes tools for routing, logging, migrations, and dynamic recompilation.  
  
### Key Components  
- **Command-line Tools (`cmd/bin`)**:  
  - **StoicMigration**: Manages database migrations.
    - Build: `cd ./cmd/src/StoicMigration ; go build -o StoicMigration.exe ; move ./StoicMigration.exe ../../bin/StoicMigration.exe ; cd ../../..`
    - Usage: `./cmd/bin/StoicMigration.exe`

  - **StoicModelBuilder**: Builds models dynamically.
    - Build: `cd ./cmd/src/StoicModelBuilder ; go build -o StoicModelBuilder.exe ; move ./StoicModelBuilder.exe ../../bin/StoicModelBuilder.exe ; cd ../../..`
    - Usage: `./cmd/bin/StoicModelBuilder.exe`

  - **wgo**: Hot recompilation tool for Go code.  
    - Build: `cd ./cmd/src/wgo-main ; go build -o wgo.exe ; move ./wgo.exe ../../bin/wgo.exe ; cd ../../..`
    - Usage: `./cmd/bin/wgo.exe run main.go -w "*.go"`
  
## Features and Usage  
  
### `Exec.ps1` Commands  
1. **init**: Sets up the development environment.  
   - **Dockerized Services**:  
     - SMTP for emails.  
     - Vite/Vue3.js frontend.  
     - Databases: MySQL, SQLServer, Postgres, SQLite.  
     - Stoic-Go backend.  
   - **Migration Control**:
     - Custom parsing for migrations:  
       - `-- StoicMigration Up`  
       - `-- StoicMigration Down`
  
2. **test**: Runs unit and integration tests.  
  
## Goals  
  
### Core Functionality  
  
#### Package `utils.go`  
- [ ] Functions:  
  - `AssertOnError(err error, format string, args ...any)`  
  - `LoggerInit()`  
  - Logging: `LogInfo()`, `LogWarn()`, `LogDebug()`, `LogError()`, `LogFatal()`  
  - `LogOnError(err error, format string, args ...any)`  
  - Generic casting: `castAny[T any](v any) T`  
  
### Runtime Tools  
  
- Install `wgo`:  
  ```  
  cd ./cmd/src/wgo-main  
  go build ; mv ./wgo ../../bin/
  ```  
- Run with hot recompilation:  
  ```  
  ./cmd/bin/wgo run main.go -w "*.go"  
  ```  
  
### TODO  
- Add models from Stoic-PHP.  
- Consolidate into a single package:  
  - `StoicCore.Router`  
  - `StoicCore.ORM`  
  - `StoicCore.Utility`  
  
## Services  
  
- **Router Package** (Uses Gorilla Mux):  
  - `router.go`:  
    - `NewRouter()`  
    - `RegisterPrefix(newPrefix string)`  
    - `RegisterApiEndpoint(path, handler, method, middlewares ...StoicMiddleware)`  
  - `middleware.go`:  
    - Public middleware:  
      - `RegisterCommonMiddleware(middlewares ...StoicMiddleware)`  
      - `MiddlewareCORS()`  
      - `MiddlewareValidParams(requiredParams ...string)`  
      - `MiddlewareLogger()`  
    - Private functions:  
      - `isMiddlewareRegistered()`  
      - `chainMiddleware()`  
  - `request.go`:  
    - `type StoicRequest`  
    - `SetError(msg string)`  
    - `SetData(data any)`  
  - `response.go`:  
    - `type StoicResponse`  
    - Functions for parameter handling (`Has`, `GetStringParam`, `GetIntParam`, etc.).  
  
## Future Work  
- ORM, Stripe
- Pre-computation for `FromX` methods to simplify templates.  
- Replace panics with robust error handling. 
- Track successful migrations in the database Migration table.

## Notes for myself:
- I'm going to use go-generate to do some metaprogramming with the tools
- dsn := ORM.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
- ORM.ConnectLocal("../SQLite.db)
- I need to figure out how setData should work with returning something like a User for example