package dao

import (
	"daily/go/logger"
	"daily/go/model"
	"daily/go/utils"
)

/**
声明一个类
*/
type PetDao struct {
	dbUtils *utils.DBUtils
}

/**
构造函数
*/
func NewPetDao() *PetDao {
	var petDaoP *PetDao = &PetDao{}
	petDaoP.dbUtils = utils.GetInstance()
	return petDaoP
}

func (petDao *PetDao) GetLastRecord() model.Pet {

	db := petDao.dbUtils.GetDB()
	rows, err := db.Query("SELECT * FROM pet_daily ORDER BY id DESC LIMIT 1")
	defer rows.Close()

	if err != nil {
		logger.Error.Println(err)
	}
	var pet model.Pet
	for rows.Next() {
		pet = model.Pet{}

		err := rows.Scan(&pet.Id, &pet.HappenTime, &pet.Type, &pet.CreateTime, &pet.UpdateTime)
		if err != nil {
			panic(err)
		}

	}

	return pet
}

func (petDao *PetDao) SaveRecord(pet model.Pet) int64 {

	db := petDao.dbUtils.GetDB()
	nowStr := utils.GetNowStr()
	res, err := db.Exec("INSERT INTO pet_daily(happen_time,type) VALUES (?,?)", nowStr, pet.Type)

	if err != nil {
		logger.Error.Println(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		logger.Error.Println(err)
	}
	return num

}
func (petDao *PetDao) ListRecord(page model.Page) []model.Pet {
	db := petDao.dbUtils.GetDB()

	startIndex := (page.Page - 1) * page.PageSize
	endIndex := startIndex + page.PageSize

	rows, err := db.Query("SELECT * FROM pet_daily ORDER BY id DESC LIMIT ?,?", startIndex, endIndex)
	defer rows.Close()

	if err != nil {
		logger.Error.Println(err)
	}
	var pets = make([]model.Pet, 0)

	for rows.Next() {
		pet := model.Pet{}

		err := rows.Scan(&pet.Id, &pet.HappenTime, &pet.Type, &pet.CreateTime, &pet.UpdateTime)
		if err != nil {
			logger.Error.Println(err)
		}

		pets = append(pets, pet)
	}

	return pets
}
