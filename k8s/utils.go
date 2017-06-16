package k8s

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Response(result interface{}, rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(result); err != nil {
		panic(err)
	}
}

func JsonUnmarshal(rw http.ResponseWriter, req *http.Request) map[string]interface{} {
	var obj interface{}
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &obj); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(422)
		if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}
	}
	m := obj.(map[string]interface{})
	return m
}
