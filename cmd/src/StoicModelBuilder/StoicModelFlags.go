package main

import (
	"strings"
)

const (
	IS_KEY         int = 1 << iota
	NULLABLE           = 1 << iota
	IS_UNIQUE          = 1 << iota
	AUTO_INCREMENT     = 1 << iota
)

func generateFlags(isNull string, isKey string, extra string) int {
	flags := 0
	if isKey == "PRI" {
		flags |= IS_KEY
	} else if isKey == "UNI" {
		flags |= IS_UNIQUE
	}

	if isNull == "NO" {
		flags |= NULLABLE
	}

	if strings.Contains(extra, "auto_increment") {
		flags |= AUTO_INCREMENT
	}

	return flags
}

func generateStrFlags(isNull string, isKey string, extra string) string {
	var flags []string

	if isKey == "PRI" {
		flags = append(flags, "ORM.KEY")
	} else if isKey == "UNI" {
		flags = append(flags, "ORM.UNIQUE")
	}

	if isNull == "YES" {
		flags = append(flags, "ORM.NULLABLE")
	}

	if strings.Contains(extra, "auto_increment") {
		flags = append(flags, "ORM.AUTO_INCREMENT")
	}

	return strings.Join(flags, "|")
}
