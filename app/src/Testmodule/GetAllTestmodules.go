package Testmodule

import "net/http"

func getAllTestmodules(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllTestmodules"))
}
