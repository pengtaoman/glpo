package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code:", resp.StatusCode)
		return nil,
			fmt.Errorf("Error: status code: %d", resp.StatusCode)
		//errors.New("Error: status code:" + strconv.Itoa(resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}
