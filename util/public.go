package util

import (
	"html/template"
	"os"
	"path"
)

func ParseFiles(filenames ...string) (*template.Template) {
	var pre string
	var t *template.Template
	t = template.New("t")
	if len(os.Args) >= 2 {
		pre = os.Args[1]
	} else {
		pre = ""
	}
	for i, v := range filenames {
		filenames[i] = path.Join(pre, v)
		t,_ = t.ParseFiles(filenames[i])
	}
	return t
}
