package proto

type TMapFieldBehavior map[string]string
type TMapTypes map[string]string

var MapTypes = TMapTypes{
	"VARCHAR":   "string",
	"CHAR":      "string",
	"BINARY":    "bytes",
	"VARBINARY": "bytes",
	"BLOB":      "google.protobuf.Any",
	"LONGBLOB":  "google.protobuf.Any",
	"LONGTEXT":  "google.protobuf.Any",
	"TEXT":      "google.protobuf.Any",
	"ENUM":      "uint32",
	"SET":       "string",
	"INTEGER":   "int",
	"INT":       "int",
	"SMALLINT":  "int",
	"TINYINT":   "bool",
	"MEDIUMINT": "int32",
	"BIGINT":    "int64",
	"DECIMAL":   "float",
	"NUMERIC":   "float",
	"FLOAT":     "float",
	"DOUBLE":    "double",
	"BIT":       "fixed32",
	"DATE":      "google.protobuf.date",
	"TIME":      "string",
	"DATETIME":  "google.protobuf.Timestamp",
	"TIMESTAMP": "google.protobuf.Timestamp",
	"YEAR":      "int",
}

var MapFieldBehavior = TMapFieldBehavior{
	"created_at":   " [(google.api.field_behavior) = OUTPUT_ONLY, \n(google.api.field_behavior) = IMMUTABLE]",
	"deleted_at":   " [(google.api.field_behavior) = OUTPUT_ONLY]",
	"updated_at":   " [(google.api.field_behavior) = OUTPUT_ONLY]",
	"finalized_at": " [(google.api.field_behavior) = OUTPUT_ONLY]",
}