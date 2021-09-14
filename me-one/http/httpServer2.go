package main

import (
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//函数式编程，函数作为参数和返回值，把输入函数包装城输出函数
func errWrapper(appHandler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := appHandler(writer, request)
		if err != nil {
			logger.Warn("文件查找发生错误：：： %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//HandleFileList
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, error := os.Open(path)
	if error != nil {
		return error
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return error
	}
	writer.Write(all)
	return nil
}

func main() {

	http.HandleFunc("/list/", errWrapper(HandleFileList))

	http.ListenAndServe(":10888", nil)
}
