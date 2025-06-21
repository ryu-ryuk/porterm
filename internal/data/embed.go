package data

import _ "embed"

//go:embed assets/glamour-catppuccin.json
var GlamourStyle []byte

//go:embed assets/paste.txt
var ProjectsData string

//go:embed content/resume.md
var ResumeData string
