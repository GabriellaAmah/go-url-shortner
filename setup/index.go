package setup

import (

	"go.mongodb.org/mongo-driver/mongo"
)


type AppConnections struct {
	Database *mongo.Database
}

var connections = &AppConnections{}

func (appCon *AppConnections) GetConnections() *mongo.Database {
	dbSetup := DB{}
	dbSetup.ConnectDb()

	appCon.Database = dbSetup.database

	return dbSetup.database
}

func AppConnectionsSetUp() *mongo.Database{

	if connections.Database != nil {
		return connections.Database
	}

	database := connections.GetConnections()
	connections.Database = database

	return database


	
}
