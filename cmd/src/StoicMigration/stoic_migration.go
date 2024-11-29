package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/superg3m/stoic-go/Core"
)

const STOIC_MIGRATION_UP_STR = "-- StoicMigration Up"
const STOIC_MIGRATION_DOWN_STR = "-- StoicMigration Down"

type MigrationMode int

const (
	MIGRATION_MODE_UP MigrationMode = iota
	MIGRATION_MODE_DOWN
)

// ./cmd/bin/migration up

// migration parse

// migration up
// migraiton down

// migration execute sql
// store migrations that have succesfully ran

// Need db connection

func hasMigrationString(data []byte, migrationStr string) bool {
	s := string(data)
	return strings.Contains(s, migrationStr)
}

func getSqlCommandsFromFile(mode MigrationMode, filePath string) []string {
	ret := []string{}
	migrationStr := []string{"-- StoicMigration Up", "-- StoicMigration Down"}

	otherMode := int(mode)
	Core.ToggleBit(&otherMode, 0)

	data, err := os.ReadFile(filePath)
	Core.AssertOnError(err)

	if !hasMigrationString(data, migrationStr[mode]) {
		err_string := fmt.Sprintf("Migratoin File doesn't' have %s", migrationStr[mode])
		Core.AssertOnError(errors.New(err_string))
		return []string{}
	}

	strData := strings.Split(string(data), "\r\n")

	for _, line := range strData {
		charAccumulator := []byte{}

		if line == migrationStr[otherMode] {
			break
		}

		for _, c := range line {
			charAccumulator = append(charAccumulator, byte(c))
			if c == '\r' || c == '\n' {
				ret = append(ret, string(charAccumulator))
				charAccumulator = nil
			}
		}
		fmt.Println(line)
	}

	return ret
}

func main() {
	sqlCommands := getSqlCommandsFromFile(MIGRATION_MODE_UP, "../../../Migrations/mysql/UserCreate.mysql")

	_ = sqlCommands
}
