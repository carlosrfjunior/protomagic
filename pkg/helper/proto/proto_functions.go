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
	"ToUpperCase":              strings.ToUpper,
	"ToLowerCase":              strings.ToLower,
	"ToTitleCase":              strings.ToTitle,
	"ToCapitalize":             cases.Title(language.English, cases.Compact).String,
	"ToPascalCase":             ToPascalCase,
	"ToTranslateType":          ToTranslateType,
	"ToCapitalWithUnderscores": ToCapitalWithUnderscores,
	"FieldBehavior":            FieldBehavior,
	"contains":                 strings.Contains,
	"hasPrefix":                strings.HasPrefix,
	"hasSuffix":                strings.HasSuffix,
	"sum":                      SumFunc,
	"dict":                     GetAllVars,
}

type Data struct {
	Schemas database.InformationSchema
	Column  database.ColumnEnum
	Options any
}

var t = template.Must(template.New("proto").Funcs(funcMap).Parse(templateProto))

// Renders the template file by creating a protobuf file
func RenderProto(infoSchema *database.InformationSchema, columnsEnum *database.ColumnEnum) {

	path := viper.GetString("protobuf.output.path")

	log.Debugf("The path output for proto files: %s", path)

	var options = map[string]string{
		"java":   "true",
		"objc":   "true",
		"csharp": "true",
		"php":    "true",
		"ruby":   "true",
	}
	customOptions := viper.GetStringMapString("protobuf.customized.options")

	if len(customOptions) < 1 {
		customOptions = options
	}

	for tableName, columns := range infoSchema.Tables {

		var inSchema = &database.InformationSchema{
			DataBaseName: infoSchema.DataBaseName,
			Tables:       make(map[string][]database.Table),
		}
		inSchema.Tables[tableName] = columns

		var columnsEnums = &database.ColumnEnum{
			DataBaseName: infoSchema.DataBaseName,
			Enums:        make(map[string][]database.Enum),
		}

		columnsEnums.Enums[tableName] = columnsEnum.Enums[tableName]

		var dataTemplate = &Data{
			Schemas: *inSchema,
			Column:  *columnsEnums,
			Options: make(map[string]string),
		}

		dataTemplate.Options = customOptions

		log.Debugln("Tabela:", tableName, "Columns:", columns, "Options:", customOptions)

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

		log.Printf("%v", dataTemplate)

		if err := t.Execute(f, dataTemplate); err != nil {
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
func ToPascalCase(t string) string {
	title := strings.ReplaceAll(t, "_", " ")
	title = cases.Title(language.English, cases.Compact).String(title)
	return strings.ReplaceAll(title, " ", "")
}

func ToCapitalWithUnderscores(t string) string {
	title := strings.ReplaceAll(t, "_", " ")
	title = cases.Upper(language.English, cases.Compact).String(title)
	return strings.ReplaceAll(title, " ", "_")
}

func ToTranslateType(t string) string {

	customMapTypes := viper.GetStringMapString("protobuf.customized.mapsTypes")

	log.Debugln("Customized Map Types [TranslateType]: ", customMapTypes)

	customMap, okc := customMapTypes[strings.ToUpper(t)]

	if okc {
		return customMap
	}

	mapTypes, ok := MapTypes[strings.ToUpper(t)]

	if ok {
		return mapTypes
	}

	return ToPascalCase(t)
}

// Adds the SUM function for the template file
func SumFunc(c, i int) int {
	return c + i
}

// func GetAllVars(key string) any {

// 	var allVars = vars{
// 		"DataBaseName": "",
// 		"Syntax":       viper.GetString("protobuf.syntax"),
// 		"ApiVersion":   viper.GetString("protobuf.apiVersion"),
// 	}

// 	if len(key) > 0 {

// 		oneVar := allVars[key]

// 		return oneVar

// 	}

// 	return ""
// }

// Returns all the variables that has be defined for the template
func GetAllVars() *vars {

	var allVars = vars{
		"DataBaseName": "",
		"Syntax":       viper.GetString("protobuf.syntax"),
		"ApiVersion":   viper.GetString("protobuf.apiVersion"),
	}

	return &allVars
}

// Returns the template file template used by protomegic
func GetTemplateFileExample() string {
	return templateProto
}

// Returns the proper field behavior type
func FieldBehavior(f string) string {
	customMapFieldBehavior := viper.GetStringMapString("protobuf.customized.fieldBehavior")
	customMap, okc := customMapFieldBehavior[strings.ToLower(f)]

	// Customization of Field Behavior that it was defined in .protomagic.yaml file
	if okc {
		return customMap
	}

	mapFieldBehavior, ok := MapFieldBehavior[strings.ToLower(f)]

	if ok {

		return mapFieldBehavior
	}

	return ""
}
