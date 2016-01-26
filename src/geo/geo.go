package geo
import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"jcleary/traveltracer/src/db"
)


// func getRoute(routeId bson.ObjectId) (coords []Coordinate, err error) {

// }

// Coordinate represents a single point on earth using latitude and longitude.
type Coordinate struct {
	Id bson.ObjectId 	`bson:"_id"`
	Latitude float64 	`bson:"latitude"`
	Longitude float64 	`bson:"longitude"`
}

type Route struct {
	Id bson.ObjectId 			`bson:"_id"`
	Coordinates []Coordinate 	`bson:"coordinates"`
}


func StoreCoordinate(coordinate Coordinate) error {
	col := db.Mgo.CoordinateCol()

	upsertData := bson.M{"$set": coordinate}
	info, err := col.Upsert(coordinate.Id, upsertData)

	if err != nil {
		return err
	}

	log.Println("UpsertId -> ", info, err)

	return nil
}

func FindCoordinate(coordinateId bson.ObjectId) error {
	col := db.Mgo.CoordinateCol()

	result := Coordinate{}
	err := col.FindId(coordinateId).One(&result)

	if err != nil {
		log.Println("FindId errorL ", err)
		return err
	}

	log.Println("Found the coordinate: ", result)
	return nil
}

