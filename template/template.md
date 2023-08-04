# My Friends

{{- range .Friends }}
- [{{ .Name }}｜{{.Description}}]({{ .URL }})
  {{- end }}

# Recent Blogs

{{- range .Posts }}
## [{{ .Title }}]({{ .PostURL }})  by [{{ .Author }}]({{ .AuthorURL }}), {{ .Date }}

{{ .Content }}

{{- end }}
