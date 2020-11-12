package test

import (
	"daily/go/model"
	"daily/go/periodDao"
	"encoding/json"
	"fmt"
	"testing"
)

func TestDao(b *testing.T) {
	num := periodDao.SaveRecord(1)
	fmt.Println(num)

	period := periodDao.GetLastRecord()

	marshal, _ := json.Marshal(&period)
	fmt.Println(string(marshal))

	periods := periodDao.ListRecord(1, 5)
	fmt.Println(len(periods))
	marshals, _ := json.Marshal(&periods)
	fmt.Println(string(marshals))

	result1 := model.Result{true, periods}
	marshals2, _ := json.Marshal(&result1)
	fmt.Println(string(marshals2))
}
