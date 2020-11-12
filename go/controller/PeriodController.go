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

func ListRecord(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("ListRecord:", string(body))

	page := model.Page{}
	json.Unmarshal(body, &page)

	if page.Page < 1 {
		page.Page = 1
	}

	if page.PageSize < 1 {
		page.PageSize = 5
	}

	periods := periodService.ListRecord(page.Page, page.PageSize)

	result := model.Result{Success: true, Obj: periods}
	resultB, _ := json.Marshal(result)
	w.Write(resultB)
}
