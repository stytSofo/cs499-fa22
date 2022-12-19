package rate

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bradfitz/gomemcache/memcache"
	rate "github.com/ucy-coast/hotel-app/internal/rate/proto"
	"github.com/ucy-coast/hotel-app/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseSession struct {
	MongoSession *mgo.Session
	CacheClient  *memcache.Client
}

func NewDatabaseSession(db_addr string, cache_addr string) *DatabaseSession {
	session, err := mgo.Dial(db_addr)
	if err != nil {
		log.Fatal("Could not establish connection with MongoDB at: " + db_addr)
	}
	log.Info("New Database session successfull...")

	memc_client := memcache.New(cache_addr)
	memc_client.Timeout = time.Second * 2
	memc_client.MaxIdleConns = 512

	return &DatabaseSession{
		MongoSession: session,
		CacheClient:  memc_client,
	}
}

func (db *DatabaseSession) LoadDataFromJsonFile(rateJsonPath string) {
	util.LoadDataFromJsonFile(db.MongoSession, "rate-db", "inventory", rateJsonPath)
}

// GetRates gets rates for hotels for specific date range.
func (db *DatabaseSession) GetRates(hotelIds []string) (RatePlans, error) {
	// TODO: Implement me

	ratePlans := make(RatePlans, 0)
	for _, id := range hotelIds {
		log.Info("Hotel ID: " + id)

		item, CacheErr := db.CacheClient.Get(id)
		//Found data in the cache
		if CacheErr == nil {
			hotel_rates := new(rate.RatePlan)

			log.Infof("Hit from cache hotelId = %v\n", id)

			if CacheErr = json.Unmarshal(item.Value, hotel_rates); CacheErr != nil {
				log.Warn("Error while unmarshalling : ", CacheErr)
			}
			ratePlans = append(ratePlans, hotel_rates)

		} else if CacheErr == memcache.ErrCacheMiss {
			hotel_rates := new(rate.RatePlan)

			log.Infof("Miss from cache at hotelId = %v\n", id)

			session := db.MongoSession.Copy()
			defer session.Close()
			c := session.DB("rate-db").C("inventory")

			err := c.Find(bson.M{"hotelId": id}).One(&hotel_rates)
			if err != nil {
				log.Info("Failed to get hotel rates: ", err)
			}
			if err == nil {
				ratePlans = append(ratePlans, hotel_rates)
			}

			//Fill the cache with the missing data
			log.Info("Filling cache with the missing data at hotel = ", id)
			rate_json, err := json.Marshal(hotel_rates)
			if err != nil {
				log.Errorf("Failed to marshal hotel [id: %v] with err:", hotel_rates.HotelId, err)
			}

			memc_str := string(rate_json)
			err = db.CacheClient.Set(&memcache.Item{Key: id, Value: []byte(memc_str)})
			if err != nil {
				log.Warn("MMC error: ", err)
			}

		} else {
			log.Errorf("Memcached error = %s\n", CacheErr)
			panic(CacheErr)
		}

	}
	return ratePlans, nil
}
