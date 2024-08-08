package controller

import (
	"fmt"

	"github.com/VijayVPatil/Helping-Hand/model"
	"github.com/VijayVPatil/Helping-Hand/store"
	"github.com/VijayVPatil/Helping-Hand/util"
)

type DbServer struct {
	DBConnection store.ConnectionDB
}

func (dbs *DbServer) NewServer(DbStore store.DbConnect) {

	util.SetLogger()
	util.Logger.Infof("Logger setup done.....\n")
	//Calling ConnectDB() method via interface from store package
	dbs.DBConnection = &DbStore
	err := dbs.DBConnection.ConnectDB()

	if err != nil {
		util.Logger.Errorf("error while creating connection to store ")
	} else {
		util.Logger.Infof("Connected with store\n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Connected to Store", nil)
	}
	fmt.Printf("The Server : %v\n", dbs)
}

type ServerOperations interface {
	NewServer(DbStore store.DbConnect)
}
