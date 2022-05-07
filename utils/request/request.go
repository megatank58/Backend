package request

import (
	"io/ioutil"
	"net/http"
)

func request(route string, url string) []byte {
	response, err := http.Get(url + route)
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return responseData
}

func GitHub(route string) []byte {
	return request(route, "https://api.github.com/")
}

func RawGitHub(route string) []byte {
	return request(route, "https://raw.githubusercontent.com/Megatank58/")
}