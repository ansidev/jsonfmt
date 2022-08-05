# CHANGELOG
{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ if .Tag.Previous }}[{{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/releases/tag/{{ .Tag.Name }}){{ else }}[{{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/releases/tag/{{ .Tag.Name }}){{ end }} - {{ datetime "2006-01-02" .Tag.Date }}
{{ range .CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
* {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }} ([{{ .Hash.Short }}]({{ $.Info.RepositoryURL }}/commit/{{ .Hash.Long }})) by [{{ .Author.Name }}](mailto:{{ .Author.Email }}) on {{ datetime "2006-01-02" .Author.Date }}.
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### {{ .Title }}
{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}
