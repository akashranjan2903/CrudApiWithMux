package utils

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

type httpMethod string

const (
	GET    httpMethod = "GET"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
	PATCH  httpMethod = "PATCH"
)

func ResponseWriter(w http.ResponseWriter, status int, data []byte, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(message))
	if len(data) > 0 {
		w.Write(data)
	}
}
func Checkmethod(method string, checkmethod httpMethod) bool {
	return method == string(checkmethod)
}

func Getidfromurl(url string) int {
	idstr := strings.Split(url, "/")
	str := idstr[len(idstr)-1]
	str = strings.Replace(str, "{", "", -1)
	str = strings.Replace(str, "}", "", -1)
	id, e := strconv.Atoi(str)
	if e != nil {
		log.Fatal("Conversion of string into int failed")
		panic(e)
	}
	return id
}

func Errorhandlefordataconversion(err error) {

	if err != nil {
		log.Fatal("error found in Marshaling the data in json")
		panic(err)
	}
}
