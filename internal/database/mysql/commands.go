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

	var query string

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
		log.Fatalln("Data Base Name: ", err)
		return err
	}

	log.Debugf("MySQL ::: Data Base Name: %s", dataBaseName)

	query = `SELECT 
		table_name, 
		column_name, 
		is_nullable, 
		data_type, 
		column_type 
		FROM information_schema.columns 
		WHERE table_schema = ? 
		ORDER BY table_name, ordinal_position ASC`

	infoSchema, err := database.GetColumns(db, dataBaseName, query, dataBaseName)

	if err != nil {
		log.Fatal("Query Tables [infoSchema]: ", err)
		return err
	}

	query = `SELECT 
	table_name, 
	column_name, 
	data_type AS type_name,
	trim(leading 'enum' from column_type) AS enum_Label, 
    ordinal_position as enum_order
	FROM information_schema.columns 
	WHERE table_schema = ? 
	AND data_type = 'enum' 
	ORDER BY table_name, ordinal_position ASC`

	columnEnum, err := database.GetColumnEnum(db, dataBaseName, query, dataBaseName)

	if err != nil {
		log.Fatal("Query Tables [columnEnum]: ", err)
		return err
	}

	proto.RenderProto(infoSchema, columnEnum)

	return nil
}
