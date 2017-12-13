package main

import (
	"net/http"

	"github.com/mkrcek/sql-server/server/mydevices"
	)



func main() {

	http.HandleFunc("/devices/", mydevices.HandleMain)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.ListenAndServe(":8080", nil)
}





//co chci udělat
// - refresh po uložení
// - podpora v index
// - sqlight
// - více tabulek přohlížení
// - a provázanost
