package hell

import (
	"fmt"
	"net/http"
)

type HellHandler struct{}

func NewHellHandler(router *http.ServeMux) {
	hanler := &HellHandler{}
	router.HandleFunc("/hell", hanler.Hell())
}
func (handler *HellHandler) Hell() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hell")
	}
}
