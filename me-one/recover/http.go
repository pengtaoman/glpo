package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	/***
		http服务发生错误后，仍然继续服务，就是使用了recover
	    net/http/server.go
		// Serve a new connection.
		func (c *conn) serve(ctx context.Context) {
			c.remoteAddr = c.rwc.RemoteAddr().String()
			ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
			defer func() {
				if err := recover(); err != nil && err != ErrAbortHandler {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
				}
				if !c.hijacked() {
					c.close()
					c.setState(c.rwc, StateClosed, runHooks)
				}
			}()

		****/
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		println("########################################", path)
		panic(errors.New("直接panic看看！！！！！"))
		file, _ := os.Open(path)
		//if error != nil {
		//	//http.Error(writer, error.Error(), http.StatusInternalServerError)
		//	http.Error(writer, errors.New("233423423423423423423").Error(), 404)
		//	return
		//}
		defer file.Close()
		all, _ := ioutil.ReadAll(file)
		//if err != nil {
		//	//panic(err)
		//	http.Error(writer, errors.New("asdfasdfasdf").Error(), 404)
		//}
		writer.Write(all)
	})

	http.ListenAndServe(":10888", nil)
}
