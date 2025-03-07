package Testmodule

import "net/http"

func getOneTestmodule(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetOneTestmodule"))

}
