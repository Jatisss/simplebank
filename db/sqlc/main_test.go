package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jati/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot Load Configuration ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to DB ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
