package testhttp

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	//println("############################", request.RequestURI)
	defer func() error {
		if r := recover(); r != nil {
			return testingUserError("panic错误发生了recover recover recover")
		} else {
			return testingUserError("panic错误发生了EEEEEEEEEEE")
		}
	}()
	panic("一个panic错误发生了！！！！")
}

func errIsNotExist(writer http.ResponseWriter, request *http.Request) error {
	err := fs.ErrNotExist
	return err
	//return testingUserError("用户自定义错误0000001")
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("用户自定义错误0000001")
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "服务器内部错误500！"},
	{errUserError, 400, "用户定义错误000000002"},
	{errIsNotExist, 404, "Not Found"},
}

func TestErrWraper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil,
		)

		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("测试错误：期待(%d, %s) ----- 得到(%d, %s)",
				tt.code,
				tt.message,
				response.Code,
				body,
			)
		}
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("测试错误：期待(%d, %s) ----- 得到(%d, %s)",
				tt.code,
				tt.message,
				resp.StatusCode,
				body,
			)
		} else {

			t.Fatalf("测试正确：期待(%d, %s) ----- 得到(%d, %s)",
				tt.code,
				tt.message,
				resp.StatusCode,
				body)
		}
	}
}
