package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type Table struct {
	ColumnName string
	IsNullable string
	DataType   string
	ColumnType string
}

// type Tables map[string][]Table

type InformationSchema struct {
	DataBaseName string
	Tables       map[string][]Table
}

func GetDataBaseName(db *sql.DB, dbQuery string, args ...any) (string, error) {

	var dataBaseName string

	rows, err := db.Query(dbQuery, args...)
	if err != nil {
		log.Fatal("GetDataBaseName DB: ", err)
		return "", err
	}
	// defer db.Close()
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dataBaseName)
		if err != nil {
			log.Fatal("GetDataBaseName Rows Next: ", err)
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("GetDataBaseName Rows: ", err)
		return "", err
	}

	return dataBaseName, nil
}

// Based on each database, it collects information from the information_schema and passes it on to the template controller
func Query(db *sql.DB, dbName string, dbQuery string, args ...any) (*InformationSchema, error) {

	var (
		t         Table
		tableName string
	)

	InfoSchema := &InformationSchema{
		DataBaseName: dbName,
		Tables:       make(map[string][]Table),
	}

	rows, err := db.Query(dbQuery, args...)
	if err != nil {
		log.Fatal("Query DB: ", err)
		return nil, err
	}
	defer db.Close()
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName, &t.ColumnName, &t.IsNullable, &t.DataType, &t.ColumnType)
		if err != nil {
			log.Fatal("Query Rows Next: ", err)
			return nil, err
		}

		InfoSchema.Tables[tableName] = append(InfoSchema.Tables[tableName], t)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Query Rows: ", err)
		return nil, err
	}

	return InfoSchema, nil
}
