package proto

var templateProto = `

{{- range $tableName, $column := .Schemas.Tables -}}

syntax = "{{ dict.Syntax }}";

package {{ $.Schemas.DataBaseName | ToLowerCase }}.api.{{ dict.ApiVersion }};

import "google/api/annotations.proto";
import "google/api/client.proto";
import "google/api/field_behavior.proto";
import "google/protobuf/timestamp.proto";
import "google/protobuf/date.proto";
import "google/protobuf/any.proto";

option go_package = "api.{{ $.Schemas.DataBaseName  | ToLowerCase }}.com/go/{{ dict.ApiVersion }}/{{ $tableName | ToLowerCase }}pb;{{ $tableName | ToLowerCase }}pb";
{{ if and ($.Options.ruby) (contains $.Options.ruby "true") }}
option ruby_package = "{{ $.Schemas.DataBaseName  | ToUpperCase }}::Api::{{ dict.ApiVersion | ToUpperCase }}";
{{ end }}
{{ if and ($.Options.java) (contains $.Options.java "true") }}
option java_package = "com.{{ $.Schemas.DataBaseName  | ToLowerCase }}.api.{{ dict.ApiVersion }}";
option java_multiple_files = true;
option java_outer_classname = "{{ $tableName | ToCapitalize }}Proto";
{{ end }}
{{ if and ($.Options.objc) (contains $.Options.objc "true") }}
option objc_class_prefix = "{{ $tableName | ToUpperCase }}";
{{ end }}
{{ if and ($.Options.csharp) (contains $.Options.csharp "true") }}
option csharp_namespace = "{{ $.Schemas.DataBaseName  | ToUpperCase }}.Api.{{ dict.ApiVersion | ToUpperCase }}";
{{ end }}
{{ if and ($.Options.php) (contains $.Options.php "true") }}
option php_namespace = "{{ $.Schemas.DataBaseName  | ToUpperCase }}\\Api\\{{ dict.ApiVersion | ToUpperCase }}";
option php_metadata_namespace = "{{ $.Schemas.DataBaseName | ToPascalCase }}";
{{ end }}

{{ if (index $.Column.Enums $tableName) }}
{{- range $table, $enums := $.Column.Enums }}
enum {{ (index $enums 0).TypeName | ToPascalCase }} {
{{- range $index, $value := $enums }} 
   {{ (index $enums 0).TypeName | ToCapitalWithUnderscores}}_{{ $value.EnumLabel | ToCapitalWithUnderscores }}={{ $index }};
{{- end }}
{{- end }}
}
{{ end }}

service {{ $tableName | ToPascalCase }}Service {
  rpc Create{{ $tableName | ToPascalCase }}(Create{{ $tableName | ToPascalCase }}Request) returns (Create{{ $tableName | ToPascalCase }}Response) {
    option (google.api.http) = {
      post: "/{{ dict.ApiVersion }}/{{ $tableName | ToLowerCase }}/create"
      body: "*"
    };
  }
  rpc Get{{ $tableName | ToPascalCase }}(Get{{ $tableName | ToPascalCase }}Request) returns (Get{{ $tableName | ToPascalCase }}Response) {
    option (google.api.http) = {get: "/{{ dict.ApiVersion }}/{{ $tableName | ToLowerCase }}/{id}/get"};
    option (google.api.method_signature) = "id";
  }
  rpc Update{{ $tableName | ToPascalCase }}(Update{{ $tableName | ToPascalCase }}Request) returns (Update{{ $tableName | ToPascalCase }}Response) {
    option (google.api.http) = {
      patch: "/{{ dict.ApiVersion }}/{{ $tableName | ToLowerCase }}/{id}/update"
      body: "id"
    };
    option (google.api.method_signature) = "id";
  }
  rpc Delete{{ $tableName | ToPascalCase }}(Delete{{ $tableName | ToPascalCase }}Request) returns (Delete{{ $tableName | ToPascalCase }}Response) {
    option (google.api.http) = {delete: "/{{ dict.ApiVersion }}/{{ $tableName | ToLowerCase }}/{id}/delete"};
    option (google.api.method_signature) = "id";
  }
}

message Get{{ $tableName | ToPascalCase }}Request {
	string id = 1;
}
message Create{{ $tableName | ToPascalCase }}Request {
	{{- range $index, $value := $column }}
	{{ $value.DataType | ToTranslateType }} {{ $value.ColumnName | ToLowerCase }} = {{ sum $index 1  }}{{ $value.ColumnName | FieldBehavior | ToLowerCase }};
	{{- end }}
}
message Update{{ $tableName | ToPascalCase }}Request {
	string id = 1;
}
message Delete{{ $tableName | ToPascalCase }}Request {
	string id = 1;
}

message Get{{ $tableName | ToPascalCase }}Response {
	{{- range $index, $value := $column }}
	{{ $value.DataType | ToTranslateType }} {{ $value.ColumnName }} = {{ sum $index 1  }}{{ $value.ColumnName | FieldBehavior }};
	{{- end }}
}
message Create{{ $tableName | ToPascalCase }}Response {
	string id = 1;
}
message Update{{ $tableName | ToPascalCase }}Response {
	string id = 1;
}
message Delete{{ $tableName | ToPascalCase }}Response {
	string id = 1;
}

{{- end -}}
`
