package repository

import (
	"curved-crater/utils"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OpenSqliteDb(databaseFile string) (*sql.DB, error) {
	if !utils.DoesThisFileExist(databaseFile) {
		f, err := os.Create(databaseFile)

		if err != nil {
			return nil, err
		}

		f.Close()
	}

	return sql.Open("sqlite3", databaseFile)
}
