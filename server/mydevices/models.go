package mydevices

import (
	"errors"
	"net/http"
	"github.com/mkrcek/sql-server/server/config"
	"encoding/json"
	"fmt"
)

type Device struct {
	deviceID	int
	deviceName   string
	deviceLocation  string
	deviceIP string
	deviceType string
	deviceBoard string
	deviceSwVersion string
	targetServer string
	httpPort string
	note string
}

type DeviceSetup struct {				//for JSON
	DeviceID	int `json:"deviceID"`
	DeviceName   string `json:"deviceName"`
	DeviceLocation  string `json:"deviceLocation"`
	DeviceIP string `json:"deviceIP"`
	DeviceType string `json:"deviceType"`
	DeviceBoard string `json:"deviceBoard"`
	DeviceSwVersion string `json:"deviceSwVersion"`
	TargetServer string `json:"targetServer"`
	HttpPort string `json:"httpPort"`
	Note string `json:"note"`
}


func AllDevices() ([]Device, error) {		//vrátí celou tabulku, všechyn záznamy
	rows, err := config.DB.Query("SELECT * FROM mydevices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mds := make([]Device, 0)
	for rows.Next() {
		dv := Device{}
		err := rows.Scan(&dv.deviceID, &dv.deviceName, &dv.deviceLocation, &dv.deviceIP,
			&dv.deviceType, &dv.deviceBoard, &dv.deviceSwVersion, &dv.targetServer,
			&dv.httpPort, &dv.note) // order matters
		if err != nil {
			return nil, err
		}
		mds = append(mds, dv)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return mds, nil
}

func OneDevice(r *http.Request, rowNumber int) (Device, error) {		//vrtí jen jeden záznam
	dv := Device{}

	ulrDeviceID:=rowNumber

	row := config.DB.QueryRow("SELECT * FROM mydevices WHERE deviceID = $1", ulrDeviceID)

	err := row.Scan(&dv.deviceID,&dv.deviceName, &dv.deviceLocation, &dv.deviceIP,
		&dv.deviceType, &dv.deviceBoard, &dv.deviceSwVersion, &dv.targetServer,
		&dv.httpPort, &dv.note)
	if err != nil {
		return dv, err
	}
	return dv, nil
}


func PutDevicePost(r *http.Request) (Device, error) {

	dv := Device{}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post DeviceSetup //vytěžení z JSON parametrů
	json.Unmarshal(body, &post)



	dv.deviceID = post.DeviceID
	dv.deviceName = post.DeviceName
	dv.deviceLocation = post.DeviceLocation
	dv.deviceIP = post.DeviceIP
	dv.deviceType = post.DeviceType
	dv.deviceBoard = post.DeviceBoard
	dv.deviceSwVersion = post.DeviceSwVersion
	dv.targetServer = post.TargetServer
	dv.httpPort = post.HttpPort
	dv.note = post.Note


	// insert values
	_, err := config.DB.Exec("INSERT INTO mydevices  (deviceName, deviceLocation, deviceIP, deviceType, deviceBoard, deviceSwVersion, targetServer, httpPort, note) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		dv.deviceName, &dv.deviceLocation, &dv.deviceIP, &dv.deviceType, &dv.deviceBoard, &dv.deviceSwVersion, &dv.targetServer, &dv.httpPort, &dv.note)
	if err != nil {
		return dv, errors.New("500. Internal Server Error." + err.Error())
	}
	return dv, nil
}

func UpdateBookApi(r *http.Request, rowNumber int) (Device, error) {

	dv := Device{}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post DeviceSetup //vytěžení z JSON parametrů

	err := json.Unmarshal(body, &post)

	ulrDeviceID := rowNumber


	if ulrDeviceID == 0 {
		return dv, errors.New("400. Bad Request.")
	}


	dv.deviceID = post.DeviceID
	dv.deviceName = post.DeviceName
	dv.deviceLocation = post.DeviceLocation
	dv.deviceIP = post.DeviceIP
	dv.deviceType = post.DeviceType
	dv.deviceBoard = post.DeviceBoard
	dv.deviceSwVersion = post.DeviceSwVersion
	dv.targetServer = post.TargetServer
	dv.httpPort = post.HttpPort
	dv.note = post.Note

	// update values

	_, err = config.DB.Exec("UPDATE mydevices SET deviceName = $2, deviceLocation = $3, deviceIP=$4, deviceType=$5, deviceBoard=$6, deviceSwVersion=$7, targetServer=$8, httpPort=$9, note=$10  WHERE deviceID = $1;",
		ulrDeviceID, dv.deviceName, dv.deviceLocation, &dv.deviceIP, &dv.deviceType, &dv.deviceBoard, &dv.deviceSwVersion, &dv.targetServer, &dv.httpPort, &dv.note)
		//pozor na PRVNÍ argument - je to číslo záznamu z URL
	//občas hloupá chyba: syntax error at or near "note" = chybí čárka před note :-)

	if err != nil {
		fmt.Print("CHYBA je tady ")
		fmt.Println(err)
		return dv, err
	}
	return dv, nil
}



func DeleteBookApi(r *http.Request, rowNumber int) error {

	ulrDeviceID:= rowNumber

	if ulrDeviceID == 0 {
		fmt.Println("chyba 400 Bad Request")
		return errors.New("400. Bad Request.")
	}

	Result, err := config.DB.Exec("DELETE FROM mydevices WHERE deviceId=$1;", ulrDeviceID)
	if err != nil {
		fmt.Println("chyba - 500 internal server error")
		return errors.New("500. Internal Server Error")
	}
	fmt.Print("Pocet smazaných záznamů: ")
	fmt.Println(Result.RowsAffected())
	return nil
}


//není potřeba zavřít DB?