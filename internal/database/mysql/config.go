package mysql

import (
	"database/sql"
	"strings"

	"github.com/carlosrfjunior/protomagic/internal/database"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	driverName     string
	dataSourceName string
}

func DB() database.DBInterface {
	return &MySQL{
		driverName:     "mysql",
		dataSourceName: viper.GetString("databases.mysql.dataSourceName"),
	}
}

func (p *MySQL) Open() (*sql.DB, error) {

	if strings.Count(p.dataSourceName, ":") <= 0 {
		return nil, nil
	}

	if !strings.Contains(p.driverName, "mysql") {
		log.Fatal("Driver name is not supported: [ mysql]")
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

	log.Println("Successfully connected to MySQL!")

	return db, err

}
