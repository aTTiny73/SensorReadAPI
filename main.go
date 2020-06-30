package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var sensorMap = make(map[string]SensorValues)

// SensorValues structure that contains all the individual measured values
type SensorValues struct {
	ID          int     `json:"id"`
	Temperature int     `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Co2         int     `json:"co2"`
}

func getTime() string {
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	return time
}

func getReadings(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.MarshalIndent(sensorMap, "", "  ")
	fmt.Fprintf(w, string(bytes))
}

func getReading(w http.ResponseWriter, req *http.Request) {
	ID := req.URL.Query().Get("ID") //wanted id/reading
	fmt.Println(ID)
	for _, v := range sensorMap {
		if strconv.Itoa(v.ID) == ID {
			bytes, _ := json.MarshalIndent(v, "", " ")
			fmt.Fprintf(w, string(bytes))

		} else {
			fmt.Println("Id not found")
		}
	}
}

func postReading(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var sensVal SensorValues
	err = json.Unmarshal(body, &sensVal)
	if err != nil {
		println(err)
	}
	sensorMap[getTime()] = sensVal
}

/*func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
*/

func main() {

	http.HandleFunc("/getReadings", getReadings)
	http.HandleFunc("/postReadings", postReading)
	http.HandleFunc("/getReading", getReading)
	http.ListenAndServe(":8090", nil)
}
