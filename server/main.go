package main

import (
	"net/http"
	"github.com/mkrcek/test-web-api/server-device/mydevices"
	)



func main() {

	http.HandleFunc("/mydevices/", mydevices.HandleMain)

	////http.HandleFunc("/", index)
	//http.HandleFunc("/mydevices/", mydevices.IndexApi)
	////localhost:8080/mydevices/
	//
	//http.HandleFunc("/mydevices/show/", mydevices.ShowApi)
	////localhost:8080/mydevices/show/2
	//
	//http.HandleFunc("/mydevices/create/", mydevices.CreateApi)
	////localhost:8080/mydevices/create/
	////POST: {"deviceName":"3tiny Garage Controller","deviceLocation":"3House","deviceIP":"192.168.1.45","deviceType":"Arduino","deviceBoard":"RobotDyn Wifi D1R2","deviceSwVersion":"2017-11-28","targetServer":"192.168.0.18","httpPort":"9091","note":"3Ovlada svetla v mistnosti"}
	//
	//http.HandleFunc("/mydevices/update/", mydevices.UpdateApi)
	////localhost:8080/mydevices/update/10 - kde za lomítkem je číslo záznamuu. Bere se u úvahu jen toto - ne v POST BODY
	////POST: {"deviceName":"-XX- Garage Controller","deviceLocation":"1733House","deviceIP":"192.168.0.45","deviceType":"Arduino","deviceBoard":"40 RobotDyn Wifi D1R2","deviceSwVersion":"2017-11-29","targetServer":"192.168.0.18","httpPort":"9091","note":"40super device"}
	//
	//http.HandleFunc("/mydevices/delete/", mydevices.DeleteApi)
	////localhost:8080/mydevices/delete/9
	//
	http.HandleFunc("/", mydevices.HandleRootDevice)
	//jen tak :-)

	http.ListenAndServe(":8080", nil)
}

//func index(w http.ResponseWriter, r *http.Request) {
//	http.Redirect(w, r, "/mydevices", http.StatusSeeOther)
//}


//pro kontrolu: zde je validace správnosti vygenerováného JSONu
//https://jsonlint.com/
