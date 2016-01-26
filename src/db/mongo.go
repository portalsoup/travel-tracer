package db
import (
	"gopkg.in/mgo.v2"
)

var (
	Mgo *DataStore = new(DataStore)
)

type DataStore struct {
	session *mgo.Session
	database string
}

func (ds *DataStore) OpenSession(dbHost string) error {
	ds = new(DataStore)
	ses, err := mgo.Dial(dbHost)
	if err != nil {
		return err
	}
	ses.SetSafe(&mgo.Safe{})

	Mgo.session = ses

	return nil
}

func (ds *DataStore) CloseSession() {
	ds.session.Close()
}

func (ds *DataStore) CopySession() *mgo.Session {
	return ds.session.Copy()
}

func (ds *DataStore) CoordinateCol() *mgo.Collection {
	return ds.CopySession().DB(ds.database).C("coordinates")
}


func (ds *DataStore) RouteCol() *mgo.Collection {
	return ds.CopySession().DB(ds.database).C("routes")
}