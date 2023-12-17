package main

import (
	"RS-Backend/api"
	"RS-Backend/config"
	"RS-Backend/dal/db"
)

var (
	dB db.IDB
)

func initDep() {
	dB = db.NewDB("host=localhost user=postgres password=postgres dbname=RS_DB port=5432 sslmode=disable TimeZone=Asia/Shanghai")

}

func main() {
	hook := config.InitFileLogger("logfile.log")
	defer hook.Flush()
	initDep()

	// Set up the router and start the server on port 8080.
	router := api.SetupRouter(dB)
	router.Run(":8080")
}
