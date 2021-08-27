package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	DELETE
	PUT
)

func (this HttpMethod) String() string {
	switch this {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case DELETE:
		return "DELETE"
	case PUT:
		return "PUT"
	default:
		return ""
	}
}

func PostByText(url string, body string) (string, error) {
	req, err := http.NewRequest(POST.String(), url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "text/plain;charset=UTF-8")
	return exeGet(req)
}

func PostByJson(url string, body string) (string, error) {
	req, err := http.NewRequest(POST.String(), url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	return exeGet(req)
}

func RequestWithForm(method HttpMethod, url string, param map[string]string) (string, error) {
	req, err := http.NewRequest(method.String(), url, strings.NewReader(""))
	if err != nil {
		return "", err
	}
	q := req.PostForm
	for k, v := range param {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return exeGet(req)
}

func RequestWithParam(method HttpMethod, url string, param map[string]string) (string, error) {
	req, err := http.NewRequest(method.String(), url, strings.NewReader(""))
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	for k, v := range param {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return exeGet(req)
}

func Get(url string) (string, error) {
	return RequestByMethod(GET, url)
}
func Post(url string) (string, error) {
	return RequestByMethod(POST, url)
}
func Delete(url string) (string, error) {
	return RequestByMethod(DELETE, url)
}

func RequestByMethod(method HttpMethod, url string) (string, error) {
	req, err := http.NewRequest(method.String(), url, strings.NewReader(""))
	if err != nil {
		return "", err
	}
	return exeGet(req)
}

func exeGet(req *http.Request) (string, error) {
	log.Println("url=" + req.URL.String())
	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		return "", err2
	}
	temp, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		return "", err3
	}
	s := string(temp)
	return s, nil
}
