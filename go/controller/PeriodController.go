package controller

import (
	"daily/go/model"
	"daily/go/periodService"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {

}

func SaveRecord(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("SaveRecord:", string(body))
	period := model.Period{}
	json.Unmarshal(body, &period)

	periodService.SaveRecord(period.Type)

	result := model.Result{Success: true}
	resultB, _ := json.Marshal(result)
	w.Write(resultB)
}

func GetLastRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetLastRecord:")
	period := periodService.GetLastRecord()

	result := model.Result{Success: true, Obj: period}
	resultB, _ := json.Marshal(result)
	w.Write(resultB)
}
