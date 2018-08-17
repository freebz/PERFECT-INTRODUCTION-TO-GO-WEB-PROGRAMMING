// 문자열 다루기

type Subject struct {
	Name  string
	Score int
}
subjects := []Subject{{"English", 86}, {"Math", 92}, {"Science", 83}}
subjectTable := template.New("subjectTable")
subjectTable.Parse(
	`<TABLE>` +
		`{{range .}}` +
		`{{printf "<TR><TD>%s</TD><TD>%d</TD></TR>" .Name .Score}}` + `{{end}}` +
		`</TABLE>`)
subjectTable.Execute(os.Stdout, subjects)
