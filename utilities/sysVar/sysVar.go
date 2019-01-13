package sysVar

import (
	"log"
	"time"
)

func init() {

	loc = initLocation()
}

var loc *time.Location

const (
	//SysLocation - SysLocation
	SysLocation = "Asia/Bangkok"
)

//Location - To Load App Time Location
func Location() *time.Location {
	return loc
}

func initLocation() *time.Location {

	cacheLoc, err := time.LoadLocation(SysLocation)
	if err != nil {
		log.Fatalln("Fail to Load System Location : ", err, SysLocation)
	}

	return cacheLoc
}
