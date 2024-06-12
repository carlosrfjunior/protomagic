package mysql

import (
	_ "github.com/go-sql-driver/mysql"
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
		log.Fatal("DB Connection: ", err)
		return err
	}

	dataBaseName, err := database.GetDataBaseName(db, "SELECT DATABASE()")

	if err != nil {
		log.Fatal("Data Base Name: ", err)
		return err
	}

	log.Debugf("MySQL ::: Data Base Name: %s", dataBaseName)

	infoSchema, err := database.Query(db, dataBaseName, "select table_name, column_name, is_nullable, data_type, column_type from information_schema.columns where table_schema = ? ORDER BY table_name, ordinal_position ASC", dataBaseName)

	if err != nil {
		log.Fatal("Query Tables: ", err)
		return err
	}

	proto.RenderProto(infoSchema)

	// for k, v := range infoSchema.Tables {

	// 	type table map[string][]database.Table
	// 	log.Println(k, v)
	// 	t := make(table)
	// 	t[k] = v
	// 	helper.RenderProto(k, t)
	// }

	log.Debugf("%v", infoSchema.Tables)

	return nil
}
