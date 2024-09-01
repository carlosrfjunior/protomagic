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

	var query string

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

	defer db.Close()

	log.Debugf("PostgreSQL ::: Data Base Name: %s", dataBaseName)

	query = `SELECT 
		table_name, 
		column_name, 
		is_nullable, 
		udt_name, 
		data_type 
		FROM information_schema.columns 
		WHERE table_schema = $1 AND table_catalog = $2 ORDER BY table_name, ordinal_position ASC
	`

	infoSchema, err := database.GetColumns(db, dataBaseName, query, "public", dataBaseName)

	if err != nil {
		log.Fatal("Query Tables [infoSchema]: ", err)
		return err
	}

	query = `SELECT 
		c.table_name AS table_name,
		c.column_name AS column_name,
		t.typname AS type_name,  
		e.enumlabel as enum_Label,
		e.enumsortorder AS enum_order
	FROM pg_type t 
	JOIN pg_enum e ON t.oid = e.enumtypid  
	JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
	JOIN information_schema.columns c ON c.udt_name = t.typname
	WHERE n.nspname = $1 AND 
			c.udt_schema = $2 AND
			c.table_catalog = $3
	ORDER BY t.typname, e.enumsortorder ASC;`

	columnEnum, err := database.GetColumnEnum(db, dataBaseName, query, "public", "public", dataBaseName)

	if err != nil {
		log.Fatal("Query Tables [columnEnum]: ", err)
		return err
	}

	proto.RenderProto(infoSchema, columnEnum)

	return nil
}
