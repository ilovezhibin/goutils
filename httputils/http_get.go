package httputils

import (
	"io/ioutil"
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
	case GET: return "GET"
	case POST: return "POST"
	case DELETE: return "DELETE"
	case PUT: return "PUT"
	default:
		return ""
	}
}


func Get(url string) (string,error) {
	req,err := http.NewRequest(GET.String(),url,strings.NewReader(""))
	if err != nil {
		return "",err
	}
	res,err2 :=http.DefaultClient.Do(req)
	if err2 != nil {
		return "",err2
	}

	temp,err3 :=ioutil.ReadAll(res.Body)
	if err3 != nil {
		return "",err3
	}
	s := string(temp)
	return s,nil

}
