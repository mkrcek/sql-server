package main

import (
	"net/http"
	"github.com/mkrcek/sql-server/server/mydevices"
	)



func main() {

	http.HandleFunc("/devices/", mydevices.HandleMain)
	http.HandleFunc("/", mydevices.HandleRootDevice)
	//root

	http.HandleFunc("/test/", mydevices.RootTest)
	//pokusný handler pro integraci  HTML /JS do web serveru

	http.ListenAndServe(":8080", nil)
}

//func index(w http.ResponseWriter, r *http.Request) {
//	http.Redirect(w, r, "/mydevices", http.StatusSeeOther)
//}


//pro kontrolu: zde je validace správnosti vygenerováného JSONu
//https://jsonlint.com/


