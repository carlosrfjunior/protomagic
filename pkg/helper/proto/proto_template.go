package proto

var templateProto = `
{{- range $tableName, $column := .Tables -}}

syntax = "{{ dict.Syntax }}";

package {{ $.DataBaseName | ToLower }}.api.{{ dict.ApiVersion }};

import "google/api/annotations.proto";
import "google/api/client.proto";
import "google/api/field_behavior.proto";
import "google/protobuf/timestamp.proto";
import "google/protobuf/date.proto";
import "google/protobuf/any.proto";

option csharp_namespace = "{{ $.DataBaseName | ToUpper }}.Api.{{ dict.ApiVersion | ToUpper }}";
option go_package = "api.{{ $.DataBaseName | ToLower }}.com/go/{{ dict.ApiVersion }}/{{ $tableName | ToLower }}pb;{{ $tableName | ToLower }}pb";
option java_multiple_files = true;
option java_outer_classname = "{{ $tableName | ToCapitalize }}Proto";
option java_package = "com.{{ $.DataBaseName | ToLower }}.api.{{ dict.ApiVersion }}";
option objc_class_prefix = "{{ $tableName | ToUpper }}";
option php_namespace = "{{ $.DataBaseName | ToUpper }}\\Api\\{{ dict.ApiVersion | ToUpper }}";
option ruby_package = "{{ $.DataBaseName | ToUpper }}::Api::{{ dict.ApiVersion | ToUpper }}";

service {{ $tableName | ToCapitalizeTable }}Service {
  rpc Create{{ $tableName | ToCapitalizeTable }}(Create{{ $tableName | ToCapitalizeTable }}Request) returns (Create{{ $tableName | ToCapitalizeTable }}Response) {
    option (google.api.http) = {
      post: "/{{ dict.ApiVersion }}/{{ $tableName | ToLower }}/create"
      body: "*"
    };
  }
  rpc Get{{ $tableName | ToCapitalizeTable }}(Get{{ $tableName | ToCapitalizeTable }}Request) returns (Get{{ $tableName | ToCapitalizeTable }}Response) {
    option (google.api.http) = {get: "/{{ dict.ApiVersion }}/{{ $tableName | ToLower }}/{id}/get"};
    option (google.api.method_signature) = "id";
  }
  rpc Update{{ $tableName | ToCapitalizeTable }}(Update{{ $tableName | ToCapitalizeTable }}Request) returns (Update{{ $tableName | ToCapitalizeTable }}Response) {
    option (google.api.http) = {
      patch: "/{{ dict.ApiVersion }}/{{ $tableName | ToLower }}/{id}/update"
      body: "id"
    };
    option (google.api.method_signature) = "id";
  }
  rpc Delete{{ $tableName | ToCapitalizeTable }}(Delete{{ $tableName | ToCapitalizeTable }}Request) returns (Delete{{ $tableName | ToCapitalizeTable }}Response) {
    option (google.api.http) = {delete: "/{{ dict.ApiVersion }}/{{ $tableName | ToLower }}/{id}/delete"};
    option (google.api.method_signature) = "id";
  }
}

message Get{{ $tableName | ToCapitalizeTable }}Request {
	string id = 1;
}
message Create{{ $tableName | ToCapitalizeTable }}Request {
	{{- range $index, $value := $column }}
	{{ $value.DataType | ToTranslateType }} {{ $value.ColumnName }} = {{ sum $index 1  }}{{ $value.ColumnName | FieldBehavior }};
	{{- end }}
}
message Update{{ $tableName | ToCapitalizeTable }}Request {
	string id = 1;
}
message Delete{{ $tableName | ToCapitalizeTable }}Request {
	string id = 1;
}

message Get{{ $tableName | ToCapitalizeTable }}Response {
	{{- range $index, $value := $column }}
	{{ $value.DataType | ToTranslateType }} {{ $value.ColumnName }} = {{ sum $index 1  }}{{ $value.ColumnName | FieldBehavior }};
	{{- end }}
}
message Create{{ $tableName | ToCapitalizeTable }}Response {
	string id = 1;
}
message Update{{ $tableName | ToCapitalizeTable }}Response {
	string id = 1;
}
message Delete{{ $tableName | ToCapitalizeTable }}Response {
	string id = 1;
}

{{- end -}}
`
