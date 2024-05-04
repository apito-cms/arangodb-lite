package arangodb

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"net/http"
	"os"
	"sync"
	"time"
)

const defaultTimeOut = time.Second * 10

func init() {
	// Set up APEX log handler
	log.SetHandler(text.New(os.Stderr))
}

// Config for the database session.
type Config struct {
	Timeout         time.Duration
	KeepAlivePeriod time.Duration
	// By default use JWT to authenticate.
	// TODO create basic auth function
	UseHttpBasicAuth bool
	// Log all http requests to db.
	DebugMode bool
	// Automatically create edge/collection on insert if non existing
	AutoCreateColOnInsert bool
	TLS                   bool
}

type Connection struct {
	client *http.Client
	header http.Header

	mu sync.Mutex
	// Connection options
	config *Config
	// Host address
	host string
	// Database
	db string
	// Authentication token
	token string
	// Collection cache
	colCache map[string]map[string]bool
}

func NewConnection(host, username, password string, config *Config) (*Connection, error) {
	var err error
	c := new(Connection)
	c.config = config
	c.host = buildHostAddress(host, config.TLS)
	c.header = http.Header{}
	c.colCache = make(map[string]map[string]bool)
	// Set default headers
	c.header.Set("Content-Type", "application/json")

	// Set custom timeout.
	// See https://goo.gl/NLk64L
	timeOut := defaultTimeOut
	if c.config.Timeout > 0 {
		timeOut = c.config.Timeout
	}

	// Connect to server
	c.client = &http.Client{
		Timeout: timeOut,
	}

	// Authenticate to the database
	err = c.authenticate(username, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}
