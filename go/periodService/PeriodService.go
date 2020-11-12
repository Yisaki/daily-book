package periodService

import (
	"daily/go/model"
	"daily/go/periodDao"
)

func ListRecord(page int) []model.Period {
	return periodDao.ListRecord(page, 10)
}

func SaveRecord(periodType int) {
	periodDao.SaveRecord(periodType)
}

func GetLastRecord() model.Period {
	return periodDao.GetLastRecord()
}
