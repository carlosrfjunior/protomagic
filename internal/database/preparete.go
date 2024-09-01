package database

import (
	"database/sql"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Table struct {
	ColumnName string
	IsNullable string
	DataType   string
	ColumnType string
}

type Enum struct {
	TableName  string
	ColumnName string
	TypeName   string
	EnumLabel  string
	EnumOrder  string
}

// type Tables map[string][]Table
type InformationSchema struct {
	DataBaseName string
	Tables       map[string][]Table
}

type ColumnEnum struct {
	DataBaseName string
	Enums        map[string][]Enum
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
func GetColumns(db *sql.DB, dbName string, dbQuery string, args ...any) (*InformationSchema, error) {

	var (
		t         Table
		tableName string
	)

	infoSchema := &InformationSchema{
		DataBaseName: dbName,
		Tables:       make(map[string][]Table),
	}

	rows, err := db.Query(dbQuery, args...)
	if err != nil {
		log.Fatal("Query DB: ", err)
		return nil, err
	}
	// defer db.Close()
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName, &t.ColumnName, &t.IsNullable, &t.DataType, &t.ColumnType)
		if err != nil {
			log.Fatal("Query Rows Next: ", err)
			return nil, err
		}

		//MySQL Implementation
		if t.DataType == "enum" {
			t.DataType = t.ColumnName
		}

		infoSchema.Tables[tableName] = append(infoSchema.Tables[tableName], t)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Query Rows: ", err)
		return nil, err
	}

	return infoSchema, nil
}

func GetColumnEnum(db *sql.DB, dbName string, dbQuery string, args ...any) (*ColumnEnum, error) {

	var (
		e         Enum
		tableName string
	)

	columnEnum := &ColumnEnum{
		DataBaseName: dbName,
		Enums:        make(map[string][]Enum),
	}

	rows, err := db.Query(dbQuery, args...)
	if err != nil {
		log.Fatal("Query DB [ColumnEnum]: ", err)
		return nil, err
	}
	// defer db.Close()
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName, &e.ColumnName, &e.TypeName, &e.EnumLabel, &e.EnumOrder)
		if err != nil {
			log.Fatal("Query Rows Next [ColumnEnum]: ", err)
			return nil, err
		}

		// MySQL Implementation
		if e.TypeName == "enum" &&
			strings.ContainsAny(e.EnumLabel, ",") {
			str := regexp.MustCompile(`[^a-zA-Z0-9,]+`).ReplaceAllString(e.EnumLabel, "")
			for _, v := range strings.Split(str, ",") {
				e.EnumLabel = v
				e.TypeName = e.ColumnName
				columnEnum.Enums[tableName] = append(columnEnum.Enums[tableName], e)
			}
		} else {
			// Postgress Implementation
			columnEnum.Enums[tableName] = append(columnEnum.Enums[tableName], e)
		}

	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Query Rows [ColumnEnum]: ", err)
		return nil, err
	}

	return columnEnum, nil
}
