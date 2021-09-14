package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		file, error := os.Open(path)
		if error != nil {
			//http.Error(writer, error.Error(), http.StatusInternalServerError)
			http.Error(writer, errors.New("233423423423423423423").Error(), 404)
			return
		}
		defer file.Close()
		all, err := ioutil.ReadAll(file)
		if err != nil {
			//panic(err)
			http.Error(writer, errors.New("asdfasdfasdf").Error(), 404)
		}
		writer.Write(all)
	})

	http.ListenAndServe(":10888", nil)
}
