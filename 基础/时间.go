package main

import (
	"fmt"
	"time"
)

func nowTime() {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("类型:\n%T\n值:\n%v\n", nowTime, nowTime)
	
	timestamp := time.Now().Unix()
	fmt.Printf("类型:\n%T\n值:\n%v\n", timestamp, timestamp)

}

func stringToUnixTime(inputStringTime string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, inputStringTime, loc)
	timeUnix := theTime.Unix()
	return strconv.FormatInt(timeUnix, 10)
}
