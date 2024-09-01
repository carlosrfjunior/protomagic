package postgresql

import (
	"database/sql"
	"strings"

	"github.com/toolsascode/protomagic/internal/database"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	driverName     string
	dataSourceName string
}

func DB() database.DBInterface {
	return &PostgreSQL{
		driverName:     "postgres",
		dataSourceName: viper.GetString("databases.postgresql.dataSourceName"),
	}
}

func (p *PostgreSQL) Open() (*sql.DB, error) {

	if strings.Count(p.dataSourceName, ":") <= 0 {
		return nil, nil
	}

	if !strings.Contains(p.driverName, "postgres") {
		log.Fatal("Driver name is not supported: [ postgres ]")
		return nil, nil
	}

	log.Debugf("Data Source Name: %s", p.dataSourceName)

	db, err := sql.Open(p.driverName, p.dataSourceName)

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return nil, err
	}
	// defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
		return nil, err
	}
	log.Println("Successfully connected to PostgreSQL!")

	return db, err

}
