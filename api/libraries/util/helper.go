package util

import (
	"github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

func StrToUint(word string) uint32 {
	num, _ := strconv.Atoi(word)
	return uint32(num)
}

func StrToInt(word string)	int  {
	num, _ := strconv.Atoi(word)
	return num
}

func StrToFloat(word string) float64 {
	word = strings.Replace(word, ",", ".", 1)
	num, _ := strconv.ParseFloat(word, 64)
	return num
}

func StrToBool(word string) bool {
	num, _ := strconv.ParseBool(word)
	return num
}

func StrToDate(word string) time.Time {
	date := strings.Split(word, "/")
	timeConvert := time.Date(StrToInt(date[2]), time.Month(StrToInt(date[1])), StrToInt(date[0]), 0, 0, 0, 0, time.UTC)
	return timeConvert
}

func StrToNullTime(word string, valid bool) mysql.NullTime {
	date := strings.Split(word, "/")
	timeConvert := time.Date(StrToInt(date[2]), time.Month(StrToInt(date[1])), StrToInt(date[0]), 0, 0, 0, 0, time.UTC)
	nullTime := mysql.NullTime{Valid:valid, Time:timeConvert}
	return nullTime
}