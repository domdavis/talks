package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Connection wraps neo4j.Driver and provides some convenience functions to
// reduce boilerplate, and the amount I need to remember.
type Connection struct {
	neo4j.Driver
}

const (
	user     = "neo4j"
	password = "password"
	realm    = ""
	uri      = "bolt://localhost:7687"
)

// New database connection using the default connection parameters and
// Basic Authorisation. New will panic if there is an error.
func New() *Connection {
	auth := neo4j.BasicAuth(user, password, realm)
	driver, err := neo4j.NewDriver(uri, auth)

	if err != nil {
		panic(err)
	}

	return &Connection{Driver: driver}
}

// Session is a logical connection to the database.
func (c *Connection) Session() *Session {
	return &Session{Session: c.NewSession(neo4j.SessionConfig{})}
}

// Close the connection to the database. Close will panic if there is an error.
func (c *Connection) Close() {
	if err := c.Driver.Close(); err != nil {
		panic(err)
	}
}

// Clear the underlying database. Clear will panic if there is an error.
func (c *Connection) Clear() {
	const stmt = "MATCH (n) detach delete n"

	c.Session().Run(stmt, nil)
}
