package main

import (
	"html/template"
	"log"
	"github.com/leandroli/my-golang-playground/ch7/exercise7.8"
	"net/http"
	"sort"
)

var people = []exercise7_8.Person{
	{"Alice", 20},
	{"Bob", 12},
	{"Bob", 20},
	{"Alice", 12},
}

var html = template.Must(template.New("people").Parse(`
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>exercise 7.9</title>
</head>
<body>
<table>
    <tr>
        <th><a href="?sort=name">name</a> </th>
        <th><a href="?sort=age">age</a> </th>
    </tr>
{{range .}}
    <tr>
        <td>{{.Name}}</td>
        <td>{{.Age}}</td>
    </tr>
{{end}}
</table>
</body>
</html>
`))

func main() {
	peopleSort := exercise7_8.NewPeopleSort(people, 2)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.FormValue("sort") {
		case "age":
			peopleSort.Select(exercise7_8.ByAge)
		case "name":
			peopleSort.Select(exercise7_8.ByName)
		}
		sort.Sort(peopleSort)
		err := html.Execute(writer, people)
		if err != nil {
			log.Printf("template error: %v", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
