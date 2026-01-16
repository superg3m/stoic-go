package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
)

const STOIC_MIGRATION_UP_STR = "-- StoicMigration Up"
const STOIC_MIGRATION_DOWN_STR = "-- StoicMigration Down"

type MigrationMode int

const (
	MIGRATION_MODE_UP MigrationMode = iota
	MIGRATION_MODE_DOWN
)

func hasMigrationString(data []byte, migrationStr string) bool {
	s := string(data)
	return strings.Contains(s, migrationStr)
}

func getSqlCommandsFromFile(mode MigrationMode, filePath string) ([]string, error) {
	migrationStr := []string{"-- StoicMigration Up\n", "-- StoicMigration Down\n"}
	delimitor := ';'

	otherMode := int(mode)
	Utility.ToggleBit(&otherMode, 0)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	string_data := strings.ReplaceAll(string(data), "\r\n", "\n")

	if !strings.Contains(string_data, migrationStr[mode]) {
		return nil, fmt.Errorf("migration file doesn't contain %s", migrationStr[mode])
	}

	lines := strings.SplitAfter(string_data, "\n")

	var ret []string
	var charAccumulator strings.Builder
	insideMigration := false

	for _, line := range lines {
		if !insideMigration && line != migrationStr[mode] {
			continue
		}

		if line == migrationStr[mode] {
			insideMigration = true
			continue
		}

		if line == migrationStr[otherMode] {
			break
		}

		for _, c := range line {
			charAccumulator.WriteByte(byte(c))
			if c == delimitor {
				ret = append(ret, charAccumulator.String())
				charAccumulator.Reset()
			}
		}
	}

	return ret, nil
}

func findFilesWithExtension(root, ext string) ([]string, error) {
	info, err := os.Stat(root)
	Utility.AssertOnErrorMsg(err, "Failed to access the root directory")
	if !info.IsDir() {
		return nil, fmt.Errorf("provided root path is not a directory: %s", root)
	}

	var files []string

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			Utility.AssertOnErrorMsg(err, fmt.Sprintf("Error accessing path: %s", path))
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(d.Name()) == ext {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <program> dbName [up|down]")
		os.Exit(1)
	}

	dbNameArg := os.Args[1]

	modeArg := os.Args[1]
	mode := MIGRATION_MODE_DOWN

	switch modeArg {
	case "up":
		mode = MIGRATION_MODE_UP
	case "down":
		mode = MIGRATION_MODE_DOWN
	default:
		fmt.Printf("Invalid argument: %s\n", modeArg)
		fmt.Println("Valid options are: 'up' or 'down'")
		os.Exit(1)
	}

	siteSettings := Utility.GetSiteSettings()
	siteSettings = siteSettings["settings"].(map[string]any)
	DB_ENGINE := Utility.CastAny[string](siteSettings["dbEngine"])
	HOST := Utility.CastAny[string](siteSettings["dbHost"])
	PORT := Utility.CastAny[int](siteSettings["dbPort"])
	USER := Utility.CastAny[string](siteSettings["dbUser"])
	PASSWORD := Utility.CastAny[string](siteSettings["dbPass"])
	DBNAME := dbNameArg

	dsn := ORM.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
	db := ORM.Register(DBNAME, DB_ENGINE, dsn)
	defer ORM.Close(DBNAME)

	files, _ := findFilesWithExtension(fmt.Sprintf("./migrations/%s", DB_ENGINE), ".sql")

	for _, file := range files {
		sqlUpCommands, err := getSqlCommandsFromFile(mode, file)
		Utility.AssertOnError(err)

		if mode == MIGRATION_MODE_UP {
			Utility.LogSuccess("Migration Up: %s", file)
		} else {
			Utility.LogDebug("Migration Down: %s", file)
		}

		for _, element := range sqlUpCommands {
			_, err := db.Exec(element)
			Utility.AssertOnError(err)
		}
	}
}
