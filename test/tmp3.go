package main

import (
	"os"
	"text/template"
)

type User struct {
	UserName string
}

var s string = `
<script type="text/javascript">
 alert(\"hello\");
</script>
`

func main() {
	t := template.New("Test")
	template.Must(t.Parse("{{if `wujunbin`}} T1 {{end}}"))
	t.Execute(os.Stdout,nil)
}
