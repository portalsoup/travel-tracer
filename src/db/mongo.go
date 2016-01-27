package db
import (
	"gopkg.in/mgo.v2"
)

var (
	Mgo *DataStore = new(DataStore)
)

// DataStore stores a database session
type DataStore struct {
	session *mgo.Session
}

// OpenSession opens a new session and stores it in this DataStore.  This should be the first method
// called before using any other methods.
func (ds *DataStore) OpenSession(dbHost string) error {
	ds = new(DataStore)
	ses, err := mgo.Dial(dbHost)
	if err != nil {
		return err
	}

	//ses.SetSafe(&mgo.Safe{})
	ses.SetMode(mgo.Monotonic, true)

	Mgo.session = ses

	return nil
}

// CloseSession closes the session maintained by this DataStore.
func (ds *DataStore) CloseSession() {
	ds.session.Close()
}

// CopySession copies this session for use by other packages.
func (ds *DataStore) CopySession() *mgo.Session {
	return ds.session.Copy()
}

// CoordinateCol returns a pointer to the coordinates Collection.
func (ds *DataStore) CoordinateCol() *mgo.Collection {
	return ds.CopySession().DB("coordinates").C("coordinates")
}

// RouteCol returns a pointer to the routes Collection.
func (ds *DataStore) RouteCol() *mgo.Collection {
	return ds.CopySession().DB("routes").C("routes")
}