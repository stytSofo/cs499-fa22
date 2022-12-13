//go:build mongodb

package rate

import (
	log "github.com/sirupsen/logrus"

	"github.com/ucy-coast/hotel-app/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseSession struct {
	MongoSession *mgo.Session
}

func NewDatabaseSession(db_addr string) *DatabaseSession {
	session, err := mgo.Dial(db_addr)
	if err != nil {
		log.Fatal("Could not establish connection with MongoDB")
		log.Fatal(db_addr)
		log.Fatal(err)
	}
	log.Info("New session successfull...")

	return &DatabaseSession{
		MongoSession: session,
	}
}

func (db *DatabaseSession) LoadDataFromJsonFile(rateJsonPath string) {
	util.LoadDataFromJsonFile(db.MongoSession, "rate-db", "inventory", rateJsonPath)
}

// GetRates gets rates for hotels for specific date range.
func (db *DatabaseSession) GetRates(hotelIds []string) (RatePlans, error) {
	// TODO: Implement me
	session := db.MongoSession.Copy()
	defer session.Close()
	c := session.DB("rate-db").C("inventory")

	ratePlans := make(RatePlans, 0)
	for _, id := range hotelIds {
		hotel_rates := new(RatePlans)
		log.Info("Hotel ID ")
		err := c.Find(bson.M{"id": id}).One(&hotel_rates)
		if err != nil {
			log.Fatalf("Failed to get hotel rates: ", err)
		}
		ratePlans = append(ratePlans, *hotel_rates...)
	}
	return ratePlans, nil
}
