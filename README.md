# High level overview of Stoic-Go

- cmd
    - bin
        - StoicMigration
        - StoicModelBuilder
        - wgo (Hot Recompiles Go code) (really good software nice job): ./cmd/bin/wgo.exe run main.go -w "*.go" 
 
## Exec.ps1 | init | stop | reset | test | migration up | migration down

### init
- Docker
    - SMTP (Emails)
    - Frontend: Vite/Vue3.js
    - Database (mysql, sqlserver, postgres, sql_lite)
    - Go Backend: Stoic-Go

- Migration Control
    - Goose: https://github.com/pressly/goose (Sql lite only -_-)
    - Custom parsing like goose
    - -- StoicMigration Up
    - -- StoicMigration Down
    - Store migration in db to know which ones have been successfully run



### test


# Goals

## Core

### Package Utils
- [ ] utils.go
    - AssertOnError(err error, format string, args ...any)
    - LoggerInit()

    - LogInfo()
    - LogWarn()
    - LogDebug()
    - LogError()
    - LogFatal()

    - LogOnError(err error, format string, args ...any)
    - castAny[T any](v any) T
    - 


cd ./tools/wgo-main
go install

i'm sorry for my sins I'm about to use a whole lot of global state purely for the name spacing if this becomes
a problem we can also consolidate it into something called StoicState and then initialize everything in teh main.go
The only issue is that the Nice Namespacing that we get will be gone!

clear ; ./tools/wgo-main/wgo run main.go -w "*.go"

Auto increment works!


I need to easily sperate runtime developer checks 
- Register Column stuff

clear ; go run ./tools/StoicMigration/stoic_migration.go up|down
clear ; ./tools/wgo-main/wgo run main.go -w "*.go"  

# TODO
Add the rest of the models from stoic-php
