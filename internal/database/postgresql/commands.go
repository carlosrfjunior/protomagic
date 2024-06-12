package postgresql

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/carlosrfjunior/protomagic/internal/database"

	"github.com/carlosrfjunior/protomagic/pkg/helper/proto"
)

type Commands struct{}

func Generate() database.DBCommands {
	return &Commands{}
}

func (c *Commands) Run() error {

	db, err := DB().Open()

	if db == nil {
		return nil
	}

	if err != nil {
		log.Fatal(err)
		return err
	}

	dataBaseName, err := database.GetDataBaseName(db, "SELECT current_database()")

	if err != nil {
		log.Fatal("Data Base Name: ", err)
		return err
	}

	log.Debugf("PostgreSQL ::: Data Base Name: %s", dataBaseName)

	infoSchema, err := database.Query(db, dataBaseName, "select table_name, column_name, is_nullable, data_type, udt_name from information_schema.columns where table_schema = $1 and table_catalog = $2 ORDER BY table_name ASC", "public", dataBaseName)

	if err != nil {
		log.Fatal(err)
		return err
	}

	proto.RenderProto(infoSchema)

	log.Debugf("%v", infoSchema)

	return nil
}
