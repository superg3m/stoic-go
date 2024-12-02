package main

import (
	"errors"
	"fmt"
	Utility2 "github.com/superg3m/stoic-go/cmd/src/Utility"
	"os"
	"strings"
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

	delimitor := ';' // This might change if you see DELIMITOR ~ or something

	otherMode := int(mode)
	Utility2.ToggleBit(&otherMode, 0)

	data, err := os.ReadFile(filePath)
	Utility2.AssertOnError(err)

	if !hasMigrationString(data, migrationStr[mode]) {
		err_string := fmt.Sprintf("Migration File doesn't' have %s", migrationStr[mode])
		Utility2.AssertOnError(errors.New(err_string))
		return []string{}
	}

	strData := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	charAccumulator := []byte{}
	found := false
	for _, line := range strData {
		if line != migrationStr[mode] && !found {
			continue
		}

		if line == migrationStr[mode] {
			found = true
			continue
		}

		if line == migrationStr[otherMode] {
			break
		}

		for _, c := range line {
			charAccumulator = append(charAccumulator, byte(c))
			if c == delimitor {
				ret = append(ret, string(charAccumulator))
				charAccumulator = nil
			}
		}
	}

	return ret
}

func main() {
	sqlUpCommands := getSqlCommandsFromFile(MIGRATION_MODE_UP, "../../../Migrations/mysql/UserCreate.mysql")

	for _, element := range sqlUpCommands {
		Utility2.LogPrint(element)
	}

	Utility2.LogPrint("\n")

	sqlDownCommands := getSqlCommandsFromFile(MIGRATION_MODE_DOWN, "../../../Migrations/mysql/UserCreate.mysql")

	for _, element := range sqlDownCommands {
		Utility2.LogPrint(element)
	}
}
