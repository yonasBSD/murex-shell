{{ if .Documents }}{{ range $i,$a := .Documents }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}
