package sprig

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"log"
	"os"
	"testing"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func TestSprig001(t *testing.T) {

	//tpl := template.Must(
	//	template.New("base").Funcs(sprig.FuncMap()).ParseGlob("*.html"),
	//)

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	//dir, _ := os.Getwd()
	//htmlFile := path.Join(dir, "base.html")
	htmlFile := "./base.html"
	fmt.Println("html file path:", htmlFile)

	stat, err := os.Stat(htmlFile)

	if err != nil {
		if os.IsExist(err) {
			fmt.Println("file exists...")
		} else {
			fmt.Println("file not exits!")
		}
	}

	fmt.Printf("file stat: %v\n", stat)

	tmpl := template.Must(template.ParseFiles(htmlFile))

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("parse template error...%v \n", err)
	}
}

func TestSprig002(t *testing.T) {
	const text = "<<.Greeting>> {{.Name}}"

	data := struct {
		Greeting string
		Name     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
	}

	tpl := template.Must(template.New("tpl").Delims("<<", ">>").Parse(text))

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSprig003(t *testing.T) {
	const text = `{{.Greeting}}, {{.Name}}
      {{ "are you ok!" | upper | repeat 5 }}
    `

	data := struct {
		Greeting string
		Name     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
	}

	tpl := template.Must(template.New("tpl").Funcs(sprig.FuncMap()).Parse(text))

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
