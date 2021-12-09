package templates

import "github.com/manifoldco/promptui"

var TracksTemplate = &promptui.SelectTemplates{
	Label: "	{{ .Name }}?",
	Active:   " ▶ {{ .Name | cyan }} {{\"-\" | cyan}}{{range .Artists}} {{.Name | cyan}} {{end}}",
	Inactive: "   {{ .Name | green }} {{ \"-\" | green}}{{range .Artists}} {{.Name | green}} {{end}}",
	Selected: " 🎧 {{ \"Playing:\" | blue }} {{ .Name | blue }}{{range .Artists}} {{.Name | blue}} {{end}}",
	Details: `
	--------- Song Details ----------
	{{ "Name:" | faint }}	{{ .Name }}
	{{ "Artist:" | faint }}	{{range .Artists}}{{.Name}} {{end}}
	{{ "Album:" | faint }}	{{ .Album.Name }}
	{{ "Release Date:" | faint }}	{{ .Album.ReleaseDate }}`,
}

var AlbumTemplate = &promptui.SelectTemplates{
	Label: "	{{ .Name }}?",
	Active:   " ▶ {{ .Name | cyan }} {{\"-\" | cyan}}{{range .Artists}} {{.Name | cyan}} {{end}}",
	Inactive: "   {{ .Name | green }} {{ \"-\" | green}}{{range .Artists}} {{.Name | green}} {{end}}",
	Selected: " ⎆ {{ \"Opening:\" | blue }} {{ .Name | blue }}{{range .Artists}} {{.Name | blue}} {{end}}",
	Details: `
--------- Album Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Artist:" | faint }}	{{range .Artists}}{{.Name}} {{end}}
{{ "Release Date:" | faint }}	{{ .ReleaseDate }}`,
}

var AlbumTracksTemplate = &promptui.SelectTemplates{
	Label: "	{{ .Name }}?",
	Active:   " ▶ {{ .Name | cyan }} {{\"-\" | cyan}}{{range .Artists}} {{.Name | cyan}} {{end}}",
	Inactive: "   {{ .Name | green }} {{ \"-\" | green}}{{range .Artists}} {{.Name | green}} {{end}}",
	Selected: " 🎧 {{ \"Playing:\" | blue }} {{ .Name | blue }}{{range .Artists}} {{.Name | blue}} {{end}}",
}

var PlaylistTemplate = &promptui.SelectTemplates{
	Label: "	{{ .Name }}?",
	Active:   " ▶ {{ .Name | cyan }} {{\"-\" | cyan}} {{ .Owner.DisplayName | cyan }}",
	Inactive: "  {{ .Name | green }} {{\"-\" | green}} {{ .Owner.DisplayName | green }}",
	Selected: " ⎆ {{ \"Opening:\" | blue }} {{ .Name | blue }} {{\"-\" | blue}}{{ .Owner.DisplayName | blue }}",
	Details: `
--------- Playlist Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Track Count:" | faint }}{{ .Tracks.Total }}`,
}

var PlaylistTracksTemplate = &promptui.SelectTemplates{
	Label: "	{{ .Track.Name }}?",
	Active:   " ▶ {{ .Track.Name | cyan }} {{\"-\" | cyan}}{{range .Track.Artists}} {{.Name | cyan}} {{end}}",
	Inactive: "   {{ .Track.Name | green }} {{ \"-\" | green}}{{range .Track.Artists}} {{.Name | green}} {{end}}",
	Selected: " 🎧 {{ \"Playing:\" | blue }} {{ .Track.Name | blue }}{{range .Track.Artists}} {{.Name | blue}} {{end}}",
	Details: `
	--------- Song Details ----------
	{{ "Name:" | faint }}	{{ .Track.Name }}
	{{ "Artist:" | faint }}	{{range .Track.Artists}}{{.Name}} {{end}}
	{{ "Album:" | faint }}	{{ .Track.Album.Name }}
	{{ "Added by:" | faint }}	{{ .AddedBy.DisplayName }}
	{{ "Added on:" | faint }}	{{ .AddedAt }}
	{{ "Release Date:" | faint }}	{{ .Track.Album.ReleaseDate }}`,
}

var DeviceTemplate = &promptui.SelectTemplates{
	Label: "	{{ .Name }}?",
	Active:   " ▶ {{ .Name | cyan }} {{\"-\" | cyan}} {{.Type | cyan}}",
	Inactive: "   {{ .Name | green }} {{ \"-\" | green}} {{.Type | green}}",
	Selected: " 🎧 {{ \"Playing:\" | blue }} {{ .Name | blue }} {.Type | blue}}",
}