package controllers

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "TESTING")
}
