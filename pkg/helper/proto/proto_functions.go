package proto

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/carlosrfjunior/protomagic/internal/database"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type variable map[string]any
type vars = variable

var allVars = vars{
	"DataBaseName": "",
}

var funcMap = template.FuncMap{
	"ToUpper":           strings.ToUpper,
	"ToLower":           strings.ToLower,
	"ToTitle":           strings.ToTitle,
	"ToCapitalize":      cases.Title(language.English, cases.Compact).String,
	"ToCapitalizeTable": CapitalizeTable,
	"ToTranslateType":   TranslateType,
	"FieldBehavior":     FieldBehavior,
	"sum":               SumFunc,
	"dict":              GetAllVars,
}

var t = template.Must(template.New("proto").Funcs(funcMap).Parse(templateProto))

// Renders the template file by creating a protobuf file
func RenderProto(infoSchema *database.InformationSchema) {

	for tableName, columns := range infoSchema.Tables {

		// type table map[string][]database.Table
		var inSchema = &database.InformationSchema{
			DataBaseName: infoSchema.DataBaseName,
			Tables:       make(map[string][]database.Table),
		}

		log.Println(tableName, columns)
		inSchema.Tables[tableName] = columns

		var f *os.File

		path := "./proto"

		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, fs.ModePerm); err != nil {
				panic(err)
			}
		}

		f, err := os.Create(fmt.Sprintf("%s/%s.proto", path, strings.ToLower(tableName)))
		if err != nil {
			panic(err)
		}

		if err := t.Execute(f, inSchema); err != nil {
			panic(err)
		}

		err = f.Close()
		if err != nil {
			panic(err)
		}

	}

}

// Capitalize the table name for message session for proto file
func CapitalizeTable(t string) string {
	title := strings.ReplaceAll(t, "_", " ")
	title = cases.Title(language.English, cases.Compact).String(title)
	return strings.ReplaceAll(title, " ", "")
}

func TranslateType(t string) string {

	mapType, ok := MapTypes[strings.ToUpper(t)]

	if ok {
		return mapType
	}

	return t
}

// Adds the SUM function for the template file
func SumFunc(c, i int) int {
	return c + i
}

// Returns all the variables that has be defined for the template
func GetAllVars() *variable {
	return &allVars
}

// Returns the template file template used by protomegic
func GetTemplateFileExample() string {
	return templateProto
}

// Returns the proper field behavior type
func FieldBehavior(f string) string {
	mapFieldBehavior, ok := MapFieldBehavior[strings.ToLower(f)]
	customMapFieldBehavior := viper.GetStringMapString("protobuf.customFieldBehavior")
	customMap, okc := customMapFieldBehavior[strings.ToLower(f)]

	// Customization of Field Behavior that it was defined in .protomagic.yaml file
	if okc {
		return customMap
	}

	if ok {

		return mapFieldBehavior
	}

	return ""
}
