# High level overview of Stoic-Go
 
## Exec.ps1 | init | stop | reset | test | migration up | migration down

### init
- Docker
    - SMTP (Emails)
    - Frontend: Vite/Vue3.js
    - Database (mysql, sqlserver, postgres, sql_lite)
    - Go Backend: Stoic-Go

- Migration Control
    - Goose: https://github.com/pressly/goose



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

