package periodDao

import (
	"daily/go/model"
	"daily/go/utils"
)

var dbUtils *utils.DBUtils

func init() {
	dbUtils = utils.GetInstance()
}

func GetLastRecord() model.Period {

	db := dbUtils.GetDB()
	rows, err := db.Query("SELECT * FROM period ORDER BY id DESC LIMIT 1")
	defer rows.Close()

	if err != nil {
		panic(err)
	}
	var period model.Period
	for rows.Next() {
		period = model.Period{}

		err := rows.Scan(&period.Id, &period.HappenTime, &period.Type, &period.CreateTime, &period.UpdateTime)
		if err != nil {
			panic(err)
		}

	}

	return period
}

func SaveRecord(periodType int) int64 {
	db := dbUtils.GetDB()
	//当前时间字符串
	nowStr := utils.GetNowStr()
	res, err := db.Exec(`INSERT INTO period(happen_time,type) VALUES (?,?)`, nowStr, periodType)
	if err != nil {
		panic(err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return num
}

func ListRecord(page int, pageSize int) []model.Period {
	db := dbUtils.GetDB()

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	rows, err := db.Query("SELECT * FROM period ORDER BY id DESC LIMIT ?,?", startIndex, endIndex)
	defer rows.Close()

	if err != nil {
		panic(err)
	}
	var periods = make([]model.Period, 0)

	for rows.Next() {
		period := model.Period{}

		err := rows.Scan(&period.Id, &period.HappenTime, &period.Type, &period.CreateTime, &period.UpdateTime)
		if err != nil {
			panic(err)
		}

		periods = append(periods, period)
	}

	return periods
}
