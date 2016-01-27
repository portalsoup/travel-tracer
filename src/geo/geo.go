package geo
import (
	"gopkg.in/mgo.v2/bson"
	"jcleary/traveltracer/src/db"
)


// Coordinate represents a single point on earth using latitude and longitude.
type Coordinate struct {
	Id bson.ObjectId 	`bson:"_id"`
	Latitude float64 	`bson:"latitude"`
	Longitude float64 	`bson:"longitude"`
}

// Route is a series of coordinates that represent a traveled path
type Route struct {
	Id bson.ObjectId 			`bson:"_id"`
	Coordinates []Coordinate 	`bson:"coordinates"`
}

// StoreCoordinate stores a coordinate into the coordinate collection
func StoreCoordinate(coordinate Coordinate) (coordinateId bson.ObjectId, err error) {

	col := db.Mgo.CoordinateCol()

	_, err = col.UpsertId(coordinate.Id, &coordinate)

	if err != nil {
		return "", err
	}

	return coordinate.Id, nil
}

// FindCoordinate finds a coordinate with a matching ID in the coordinate collection
func FindCoordinate(coordinateId bson.ObjectId) (result Coordinate, err error) {
	col := db.Mgo.CoordinateCol()

	result = Coordinate{}
	err = col.FindId(coordinateId).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

// FindAllCoordinates finds all coordinates contained in the coordinate collection
func FindAllCoordinates() (result []Coordinate, err error) {
	col := db.Mgo.CoordinateCol()
	result = []Coordinate{}
	err = col.Find(nil).All(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}


