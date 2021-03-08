// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"fmt"

	github4 "github.com/gopl.io/ch4/github"

	// "io/ioutil"
	"log"
	"os"
	"text/template"
)

//var issueList = template.Must(template.New("issuelist").Parse(`
//<h1>{{.TotalCount}} issues</h1>
//<table>
//<tr style='text-align: left'>
//  <th>#</th>
//  <th>State</th>
//  <th>User</th>
//  <th>Title</th>
//</tr>
//{{range .Items}}
//<tr>
//  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
//  <td>{{.State}}</td>
//  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
//  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
//</tr>
//{{end}}
//</table>
//`))

//!-template

func main() {
	//result, err := github4.SearchIssues(os.Args[1:])
	result := github4.IssuesSearchResult{
		TotalCount: 1113,
		Items:      nil,
	}
	issueList, _ := template.ParseFiles("/Volumes/860EVO/go-path/src/github.com/gopl.io/ch4/issueshtml/aa.yaml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("===============================================%_v", result)
	println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	//dir, _ := ioutil.TempDir("","yaml")
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//!-
