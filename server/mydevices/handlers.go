package mydevices

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
	"regexp"
	"strings"
	"strconv"
	"github.com/mkrcek/sql-server/server/config"
)


func IndexApi(w http.ResponseWriter, r *http.Request) {

	mds, err := AllDevices()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	numberOfRows := len(mds)
	fmt.Println("pocet záznamů", numberOfRows)
	myDeviceSetup := make([]DeviceSetup, numberOfRows)

	//kopírování obsahu z SQL do pole
	for i:=0;i<numberOfRows ;i++  {
		myDeviceSetup[i]=DeviceSetup{
			DeviceID:   mds[i].deviceID,
			DeviceName:  mds[i].deviceName,
			DeviceLocation: mds[i].deviceLocation,
			DeviceIP:  mds[i].deviceIP,
			DeviceType:   mds[i].deviceType,
			DeviceBoard:  mds[i].deviceBoard,
			DeviceSwVersion: mds[i].deviceSwVersion,
			TargetServer:  mds[i].targetServer,
			HttpPort:   mds[i].httpPort,
			Note:  mds[i].note,
		}
	}
	fmt.Println(myDeviceSetup)
	b, err := json.Marshal(myDeviceSetup)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(bytes.NewBuffer(b)) //vypis co se nakopírovalo

	//úprava hlavičky
	//w.Header().Set("Content-Length", "0") - POZOR DELKA NEMUZE BYT NULA
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")

	//uprava těla - vrátí do HTTP GET
	w.Write(b)
}


func ShowApi(w http.ResponseWriter, r *http.Request, rowNumber int) {

	mds, err := OneDevice(r, rowNumber)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	myDeviceSetup := DeviceSetup{}


	myDeviceSetup = DeviceSetup{
		DeviceID:        mds.deviceID,
		DeviceName:      mds.deviceName,
		DeviceLocation:  mds.deviceLocation,
		DeviceIP:        mds.deviceIP,
		DeviceType:      mds.deviceType,
		DeviceBoard:     mds.deviceBoard,
		DeviceSwVersion: mds.deviceSwVersion,
		TargetServer:    mds.targetServer,
		HttpPort:        mds.httpPort,
		Note:            mds.note,
	}

	b, err := json.Marshal(myDeviceSetup)
	if err != nil {
		fmt.Println("error:", err)
	}

	//úprava hlavičky
	//w.Header().Set("Content-Length", "0") - POZOR DELKA NEMUZE BYT NULA
	w.Header().Set("Connection", "keep-alive")
//nově:
	w.Header().Set("Content-Type", "application/json") //

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")

	//uprava těla - vrátí do HTTP GET
	w.Write(b)

	fmt.Println("***************")
	fmt.Println("ALL do DB: ", bytes.NewBuffer(b))
	fmt.Println("***************")

}




func CreateApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	dv, err := PutDevicePost(r)

	//úprava hlavičky
	//w.Header().Set("Content-Length", "0") - POZOR DELKA NEMUZE BYT NULA
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")


	if err != nil {
		w.Write([]byte("ERROR"))
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	} else {
		//http.Error(w, http.StatusText(200),http.StatusCreated)
		//asi by mělo být něco jiného než ERROR - ale nemohu najít.
		//uprava těla - vrátí do HTTP GET
		w.Write([]byte("OK CREATED"))


		fmt.Println("***************")
		fmt.Println("Přidáno do DB: ", dv)
		fmt.Println("***************")
	}

}



func UpdateApi(w http.ResponseWriter, r *http.Request, rowNumber int) {

	dv, err := UpdateBookApi(r, rowNumber)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}


	//úprava hlavičky
	//w.Header().Set("Content-Length", "0") - POZOR DELKA NEMUZE BYT NULA
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")

	//zpráva do těla
	w.Write([]byte("OK UPDATED"))


	fmt.Println("***************")
	fmt.Println("Update   : ", dv)
	fmt.Println("***************")
	fmt.Println(dv.deviceBoard)
	fmt.Println(dv.deviceID)
}

func DeleteApi(w http.ResponseWriter, r *http.Request, rowNumber int) {

	err := DeleteBookApi(r, rowNumber)
	if err != nil {
		fmt.Println("chyba", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	//úprava hlavičky
	//w.Header().Set("Content-Length", "0") - POZOR DELKA NEMUZE BYT NULA
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")

	//zpráva do těla
	w.Write([]byte("OK - DELETED"))

	fmt.Println("***************")
	fmt.Println("DELETEd   : ")
	fmt.Println("***************")
}

func HandleRootDevice(w http.ResponseWriter, req *http.Request) {

	////config.TPL.ExecuteTemplate(w, "update.gohtml", bk)

	fmt.Printf("AHHHOJ")

	w.Header().Set("Content-Type", "application/text") //nebo text/json
	w.Header().Set("Connection", "keep-alive")
	//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/device")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")

	fmt.Println("This is my text in ROOT")

	//config.TPL.ExecuteTemplate(w, "/index.html", nil)

	w.Write([]byte("This is my text in ROOT"))
}

func OptionsApi(w http.ResponseWriter, req *http.Request) {
	//odpověď na volání, pokud by se místo POST klient ptal na OPTIONS
	//tato varianta je pro CORS - https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request

	w.Header().Set("Content-Length", "0")
	w.Header().Set("Connection", "keep-alive")
	//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/device")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(200)
	return
}

func HandleMain(w http.ResponseWriter, r *http.Request) {

	//nalezení, jeslti je zvolený nějaký konkrétní řádek /zázmam = ID (číslo)
	//...číslo za lomítkem, např./devices/21293 = 21293
	myURL := r.RequestURI        // req.URL vs req.RequestURI		"/devices/21293"
	re := regexp.MustCompile("[0-9]+")		//vyfiltruje všechna čísla = 21293
	ulrDeviceIDs := re.FindAllString(myURL, 1)	//vybere jen první sekvenci číslic = 21293
	ulrDeviceIDstr := strings.Join(ulrDeviceIDs,"")			//převede type []string” to string
	ulrDeviceID, _ := strconv.Atoi(ulrDeviceIDstr)			//převede na číslo

	fmt.Println("parametr: ", ulrDeviceID)
	fmt.Println("metoda: ", r.Method)		//výpis metody - jen pro info

	switch r.Method {

	case "GET":
		if ulrDeviceID == 0 {	//není žádné číslo - tak vypíšu všechny záznamy
			IndexApi(w, r)
			//GET localhost:8080/mydevices/
			//vrátí všechny záznamy
		} else {				//je konkrétní číslo - tak vrátím požadovaný řádek
			ShowApi(w,r,ulrDeviceID )
			//GET localhost:8080/mydevices/2
			//vrátí 2.záznam, tj. řádek pro ID=2
		}
	case "OPTIONS":
		OptionsApi(w, r) //možná někdo zavolá
	case "POST":
		CreateApi(w,r)
		//POST localhost:8080/mydevices/
		//vytvoří nový záznam
		//např. POST {"deviceName":"2249tiny Garage Controller","deviceLocation":"1733House","deviceIP":"192.168.0.45","deviceType":"Arduino","deviceBoard":"44 RobotDyn Wifi D1R2","deviceSwVersion":"2017-11-29","targetServer":"192.168.0.18","httpPort":"9091","note":"44super device"}
	case "PUT":
		UpdateApi(w,r, ulrDeviceID)
		//PUT localhost:8080/mydevices/2
		//aktualizuje záznam číslo 2.
		//{"deviceName":"-2306- Garage Controller","deviceLocation":"2017House","deviceIP":"192.168.0.45","deviceType":"Arduino","deviceBoard":"40 RobotDyn Wifi D1R2","deviceSwVersion":"2017-11-29","targetServer":"192.168.0.18","httpPort":"9091","note":"40super device"}
	case "DELETE":
		DeleteApi(w,r, ulrDeviceID)
		//DELETE localhost:8080/mydevices/2
		//smaže 2.záznam
	}
}


func RootTest(w http.ResponseWriter, req *http.Request) {

	config.TPL.ExecuteTemplate(w, "index.html", nil)


	////config.TPL.ExecuteTemplate(w, "update.gohtml", bk)

	//
	//w.Header().Set("Content-Type", "application/text") //nebo text/json
	//w.Header().Set("Connection", "keep-alive")
	////rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/device")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type");
	//w.Header().Set("Access-Control-Max-Age", "86400")
	//
	//fmt.Println("This is my text in TEST")
	//w.Write([]byte("This is my text in TEST "))
}