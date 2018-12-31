package main

var staticTemplatesData = map[string]string{
	"at_end":             "\n// Check if the rest is empty\nif len(p.rest) > 0 {\nreturn false, {{ if .Serious }}fmt.Errorf(\"Rest must be empty, not `\\033[1m%s\\033[0m`\", string(p.rest)){{else}}nil{{end}}\n}\n",
	"close_option_scope": ";p.{{.Name}} = true\n{{.PRest}} = {{.Rest}}\n{{ if .WasAbandoned }}{{.ScopeLabel}}:\n{{end}}\n",
	"close_scope":        "};\n",
	"getter":             "\n// Get{{.LongName}} retrieves optional value for {{.LongName}}.Name\nfunc (p *{{.ParserName}}) Get{{.LongName}}() (res {{.Type}}) {\n\t if {{ range $index, $data := .Accesses }}{{ if $index }}||{{end}}!p.{{$data}}.Valid{{end}} {\n  \t return\n\t };\n\t return p.{{.Access}}.{{.Name}}\n}\n",
	"open_option":        "{{.Name}} struct {\n   Valid bool;\n",
	"open_option_scope":  ";{{.Rest}} = {{.PRest}};\n",
	"parser_body":        "// Extract autogenerated method of {{.ParserName}}\nfunc (p *{{.ParserName}}) Extract(Line []byte) (bool, error) {\n   {{ range .Vars }} var {{.Name}} {{.Type}};{{end}}\n   p.rest = Line\n   {{.Parser}}\n\n   return true, nil\n}\n",
	"parser_code":        "package {{ .PkgName }}\n\n{{ if .Imports }}import (\n{{ range .Imports }}{{.Name}} \"{{.Path}}\"\n{{end}}\n){{end}}\n\n{{range $name, $value := .Consts }}var {{$name}} = []byte({{$value}});{{end}}\n\n\n{{.Struct}}\n\n{{.Parser}}\n\n{{.Getters}}\n",
	"pass_n_items":       "\n// Cut out first N items from the rest\nif len({{.Rest}}) >= {{.Upper}} {\n {{.Rest}} = {{.Rest}}[{{.Upper}}:]\n} else {\n{{if .Namespace}}p.{{.Namespace}}.Valid = false; goto {{.ScopeLabel}}{{else}}return false, {{ if .Serious }}fmt.Errorf(\"Cannot cut out first {{.Upper}} symbols from the rest: it is shorter (%d)\", len({{.Rest}})){{else}}nil{{end}};{{end}}\n}\n",
	"soft_exit":          "return false, nil\n",
	"struct_body":        "// {{.ParserName}} autogenerated parser\ntype {{.ParserName}} struct {\nrest []byte\n{{.Struct}}\n}\n",
	"struct_field":       "{{.Name}} {{.Type}};\n",
}
