package controller

import (
	"daily/go/logger"
	"daily/go/model"
	"daily/go/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var petService *service.PetService

func init() {
	petService=service.NewPetService()
}

func SavePetRecord(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	logger.Info.Println("SavePetRecord:", string(body))
	pet := model.Pet{}
	json.Unmarshal(body, &pet)

	petService.SaveRecord(pet)

	result := model.Result{Success: true}
	resultB, _ := json.Marshal(result)
	w.Write(resultB)
}

func GetPetLastRecord(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("GetPetLastRecord:")
	pet := petService.GetLastRecord()

	result := model.Result{Success: true, Obj: pet}
	resultB, _ := json.Marshal(result)
	w.Write(resultB)
}

func ListPetRecord(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	logger.Info.Println("ListRecord:", string(body))

	page := model.Page{}
	json.Unmarshal(body, &page)

	if page.Page < 1 {
		page.Page = 1
	}

	if page.PageSize < 1 {
		page.PageSize = 5
	}

	pets := petService.ListRecord(page)

	result := model.Result{Success: true, Obj: pets}
	resultB, _ := json.Marshal(result)
	w.Write(resultB)
}
