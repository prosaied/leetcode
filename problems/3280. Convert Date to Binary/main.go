package main

import (
	"fmt"
	"strconv"
)

func convertDateToBinary(date string) string {
	// var result string
	year, _ := strconv.Atoi(date[0:4])
	yearBin := strconv.FormatInt(int64(year), 2)
	month, _ := strconv.Atoi(date[5:7])
	monthBin := strconv.FormatInt(int64(month), 2)
	day, _ := strconv.Atoi(date[8:10])
	dayBin := strconv.FormatInt(int64(day), 2)
	return yearBin + "-" + monthBin + "-" + dayBin
}

func main() {
	fmt.Println(convertDateToBinary("2080-02-29"))
}
