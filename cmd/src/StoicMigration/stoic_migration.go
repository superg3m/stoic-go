package StoicMigration

import (
	"fmt"
	"os"
	"strings"

	"github.com/superg3m/stoic-go/Core"
)

const STOIC_MIGRATION_UP_STR = "-- StoicMigration Up"
const STOIC_MIGRATION_DOWN_STR = "-- StoicMigration Down"

// ./cmd/bin/migration up

// migration parse

// migration up
// migraiton down

// migration execute sql
// store migrations that have succesfully ran

// Need db connection

func hasMigrationString(filePath string, migrationStr string) {
	b, err := os.ReadFile(filePath)
	Core.AssertOnError(err)
	s := string(b)
	// //check whether s contains substring text
	fmt.Println(strings.Contains(s, migrationStr))
}
