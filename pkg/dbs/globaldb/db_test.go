package globaldb

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/influenzanet/study-service/pkg/types"
)

var testDBService *GlobalDBService

const (
	testDBNamePrefix = "TEST_"
)

// Pre-Test Setup
func TestMain(m *testing.M) {
	setupTestDBService()
	result := m.Run()
	dropTestDB()
	os.Exit(result)
}

func setupTestDBService() {
	connStr := os.Getenv("GLOBAL_DB_CONNECTION_STR")
	username := os.Getenv("GLOBAL_DB_USERNAME")
	password := os.Getenv("GLOBAL_DB_PASSWORD")
	prefix := os.Getenv("GLOBAL_DB_CONNECTION_PREFIX")
	if connStr == "" || username == "" || password == "" {
		log.Fatal("Couldn't read DB credentials.")
	}
	URI := fmt.Sprintf(`mongodb%s://%s:%s@%s`, prefix, username, password, connStr)
	var err error
	Timeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		log.Fatal("DB_TIMEOUT: " + err.Error())
	}
	IdleConnTimeout, err := strconv.Atoi(os.Getenv("DB_IDLE_CONN_TIMEOUT"))
	if err != nil {
		log.Fatal("DB_IDLE_CONN_TIMEOUT" + err.Error())
	}
	mps, err := strconv.Atoi(os.Getenv("DB_MAX_POOL_SIZE"))
	MaxPoolSize := uint64(mps)
	if err != nil {
		log.Fatal("DB_MAX_POOL_SIZE: " + err.Error())
	}
	testDBService = NewGlobalDBService(
		types.DBConfig{
			URI:             URI,
			Timeout:         Timeout,
			IdleConnTimeout: IdleConnTimeout,
			MaxPoolSize:     MaxPoolSize,
			DBNamePrefix:    testDBNamePrefix,
		},
	)
}

func dropTestDB() {
	log.Println("Drop test database: globaldb package")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := testDBService.DBClient.Database(testDBNamePrefix + "global-infos").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}