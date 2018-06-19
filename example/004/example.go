// +build go1.7

package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
)

func main() {
	mux := bone.New()
	mux.CaseSensitive = true
	mux.RegisterValidator("isNum", func(s string) bool {
		if _, err := strconv.Atoi(s); err == nil {
			return true
		}
		return false
	})

	mux.GetFunc("/ctx/:age|isNum/name/:name", rootHandler)

	http.ListenAndServe(":8080", mux)
}

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), "var", bone.GetValue(req, "var"))
	subHandler(rw, req.WithContext(ctx))
}

func subHandler(rw http.ResponseWriter, req *http.Request) {
	vars := bone.GetAllValues(req)
	age := vars["age"]
	name := vars["name"]
	rw.Write([]byte(age + " " + name))
}
