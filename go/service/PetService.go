package service

import (
	"daily/go/dao"
	"daily/go/model"
)

/**
声明一个类
*/
type PetService struct {
	petDao *dao.PetDao
}

/**
构造函数
*/
func NewPetService() {
	petService := &PetService{}
	petService.petDao = dao.NewPetDao()
}

func (petService *PetService) ListRecord(page model.Page) []model.Pet {
	return petService.petDao.ListRecord(page)
}

func (petService *PetService) SaveRecord(pet model.Pet) {
	petService.petDao.SaveRecord(pet)
}

func (petService *PetService) GetLastRecord() model.Pet {
	return petService.petDao.GetLastRecord()
}
