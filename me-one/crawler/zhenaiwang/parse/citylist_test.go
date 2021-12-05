package parse

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//content, _ := ioutil.ReadFile("test.html")
	resp, _ := http.Get("http://www.zhenai.com/zhenghun")
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	result := ParseCityList(content, "")

	for _, item := range result.Items {
		log.Printf("Got item  %s", item)
	}
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("Result should have %d request; but had %d resultsets!", resultSize, len(result.Requests))
	}

}
