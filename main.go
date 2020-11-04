package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Loc struct {
	Location  string    `json:"Location"`
	LocalTime time.Time `json:"LocalTime"`
}

type server struct{}

func getLocalTime(loc string) time.Time {
	t := time.Now()
	location, err := time.LoadLocation(loc)
	if err != nil {
		fmt.Println(err)
	}
	return t.In(location)
}

func LocationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	switch ps.ByName("city") {
	case "Lublin":
		Location := Loc{"Lublin", getLocalTime("Europe/Warsaw")}
		json.NewEncoder(w).Encode(Location)
	case "New_York":
		Location := Loc{"New_York", getLocalTime("America/New_York")}
		json.NewEncoder(w).Encode(Location)
	case "Sydney":
		Location := Loc{"Sydney", getLocalTime("Australia/Sydney")}
		json.NewEncoder(w).Encode(Location)
	default:
		return
	}

}

func main() {
	router := httprouter.New()
	router.GET("/:city", LocationHandler)

	http.ListenAndServe(":7500", router)
}
