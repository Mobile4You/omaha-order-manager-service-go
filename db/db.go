package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	host     = "mongodb://localhost:28017"
	database = "order"
)

var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
)

// MgoDb represent connection mongo and session
type MgoDb struct {
	Session *mgo.Session
	Db      *mgo.Database
	//Col     *mgo.Collection
}

func start() {

	if mainSession == nil {
		var err error
		mainSession, err = mgo.Dial(host)
		if err != nil {
			log.Printf("Error starting connection to database: %v", err)
			return
		}
		mainSession.SetMode(mgo.Primary, true)
		mainDb = mainSession.DB(database)
	}
}

// Open connection
func (mg *MgoDb) Open() {
	mg.Session = mainSession.Copy()
	mg.Db = mg.Session.DB(database)
}

// Close connection
func (mg *MgoDb) Close() bool {
	defer mg.Session.Close()
	return true
}
