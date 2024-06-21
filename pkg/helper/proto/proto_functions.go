package proto

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/carlosrfjunior/protomagic/internal/database"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type variable map[string]any
type vars = variable

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

	path := viper.GetString("protobuf.output.path")

	log.Debugf("The path output for proto files: %s", path)

	if viper.GetBool("protobuf.output.reset") {
		log.Debugf("You have chosen to recreate the directory: %s", path)
		os.RemoveAll(path)
	}

	for tableName, columns := range infoSchema.Tables {

		var inSchema = &database.InformationSchema{
			DataBaseName: infoSchema.DataBaseName,
			Tables:       make(map[string][]database.Table),
		}
		inSchema.Tables[tableName] = columns

		log.Debugln("Tabela:", tableName, "Columns:", columns)

		var f *os.File

		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, fs.ModePerm); err != nil {
				log.Panic(err)
			}
		}

		var filename = fmt.Sprintf("%s/%s.proto", path, strings.ToLower(tableName))

		f, err := os.Create(filename)
		if err != nil {
			log.Panic(err)
		}

		if err := t.Execute(f, inSchema); err != nil {
			log.Panic(err)
		}

		err = f.Close()
		if err != nil {
			log.Panic(err)
		}

	}

	cmd := exec.Command("buf", "format", path, "-w")
	if err := cmd.Run(); err != nil {
		log.Errorf("Unable to execute buf command for indentation in path %s. Please verify! \n %v", path, err)
	}

	log.Infoln("Process rendered and finalized successfully!!!")

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
func GetAllVars() *vars {
	return &vars{
		"DataBaseName": "",
		"Syntax":       viper.GetString("protobuf.syntax"),
		"ApiVersion":   viper.GetString("protobuf.apiVersion"),
	}
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
