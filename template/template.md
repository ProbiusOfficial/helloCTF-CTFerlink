---
comments: true

---
# CTFer档案馆
## List

{{- range .Friends }}
- [{{ .Name }}｜{{.Description}}]({{ .URL }})
  {{- end }}

## Recent Post

{{- range .Posts }}
### [{{ .Title }}]({{ .PostURL }})  
>by [{{ .Author }}]({{ .AuthorURL }}), {{ .Date }}

{{ .Content }}...

{{- end }}
