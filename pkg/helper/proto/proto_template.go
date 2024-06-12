package proto

var templateProto = `
{{- range $tableName, $column := .Tables -}}

syntax = "proto3";

package {{ $.DataBaseName | ToLower }}.api.v1;

import "google/api/annotations.proto";
import "google/api/client.proto";
import "google/api/field_behavior.proto";
import "google/protobuf/timestamp.proto";
import "google/protobuf/date.proto";
import "google/protobuf/any.proto";

option csharp_namespace = "{{ $.DataBaseName | ToUpper }}.Api.V1";
option go_package = "api.{{ $.DataBaseName | ToLower }}.com/go/v1/{{ $tableName | ToLower }}pb;{{ $tableName | ToLower }}pb";
option java_multiple_files = true;
option java_outer_classname = "{{ $tableName | ToCapitalize }}Proto";
option java_package = "com.{{ $.DataBaseName | ToLower }}.api.v1";
option objc_class_prefix = "{{ $tableName | ToUpper }}";
option php_namespace = "{{ $.DataBaseName | ToUpper }}\\Api\\V1";
option ruby_package = "{{ $.DataBaseName | ToUpper }}::Api::V1";

message Get{{ $tableName | ToCapitalizeTable }}Request {
	{{- range $index, $value := $column }}
	{{ $value.DataType | ToTranslateType }} {{ $value.ColumnName }} = {{ sum $index 1  }}{{ $value.ColumnName | FieldBehavior }};
	{{- end }}
}
message Create{{ $tableName | ToCapitalizeTable }}Request {
	string id = 1;
}
message Update{{ $tableName | ToCapitalizeTable }}Request {
	string id = 1;
}
message Delete{{ $tableName | ToCapitalizeTable }}Request {
	string id = 1;
}

{{- end -}}
`
