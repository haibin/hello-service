package tests

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/haibin/hello-service/business/data/schema"
	"github.com/haibin/hello-service/business/database"
	"github.com/jmoiron/sqlx"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

// Configuration for running tests.
var (
	dbImage = "postgres:13-alpine"
	dbPort  = "5432"
	dbArgs  = []string{"-e", "POSTGRES_PASSWORD=postgres"}
	AdminID = "5cf37266-3473-4006-984f-9325122678b7"
	UserID  = "45b5fbd3-755f-4379-8f07-a58d4a30fa2f"
)

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T) (*log.Logger, *sqlx.DB, func()) {
	c := startContainer(t, dbImage, dbPort, dbArgs...)

	db, err := database.Open(database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       c.Host,
		Name:       "postgres",
		DisableTLS: true,
	})
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Waiting for database to be ready ...")

	// Wait for the database to be ready. Wait 100ms longer between each attempt.
	// Do not try more than 20 times.
	var pingError error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
	}

	if pingError != nil {
		dumpContainerLogs(t, c.ID)
		stopContainer(t, c.ID)
		t.Fatalf("Database never ready: %v", pingError)
	}

	if err := schema.Migrate(db); err != nil {
		stopContainer(t, c.ID)
		t.Fatalf("Migrating error: %s", err)
	}

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()
		stopContainer(t, c.ID)
	}

	log := log.New(os.Stdout, "TEST : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	return log, db, teardown
}

// Test owns state for running and shutting down tests.
type Test struct {
	TraceID  string
	DB       *sqlx.DB
	Log      *log.Logger
	KID      string
	Teardown func()

	t *testing.T
}

// NewIntegration creates a database, seeds it, constructs an authenticator.
func NewIntegration(t *testing.T) *Test {
	log, db, teardown := NewUnit(t)

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	//// Build an authenticator using this private key and id for the key store.
	//auth, err := auth.New("RS256", keystore.NewMap(map[string]*rsa.PrivateKey{keyID: privateKey}))
	//if err != nil {
	//	t.Fatal(err)
	//}

	test := Test{
		TraceID:  "00000000-0000-0000-0000-000000000000",
		DB:       db,
		Log:      log,
		t:        t,
		Teardown: teardown,
	}

	return &test
}
