package test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/flosch/pongo2/v6"
)

func hasUser(user string) bool {
	println("hasUser ", user)
	return user != ""
}

func TestUser(t *testing.T) {
	tmpl := template.New("example")
	tmpl = tmpl.Funcs(template.FuncMap{"hasUser": hasUser})
	tmpl, _ = tmpl.Parse(`{{if hasUser .user}}{{else}};{{end}} User={{.user}}`)

	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{"user": "John"})
	fmt.Println("1", buf.String()) // 输出 User=John

	buf.Reset()
	tmpl.Execute(&buf, map[string]string{})
	fmt.Println("2", buf.String()) // 输出 ; User=
}

func TestPgo(t *testing.T) {
	// Compile the template first (i. e. creating the AST)
	tpl, err := pongo2.FromString("Hello {{ name|capfirst }}! {% if name != null %} ! User={{name}} {% endif %}")
	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(pongo2.Context{"name": "florian"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!

	out, err = tpl.Execute(pongo2.Context{"name1": "florian"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!
}
