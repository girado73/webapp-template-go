package Testmodule

import "net/http"

func createSingleTestmodule(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateSingleTestmodule"))
}
