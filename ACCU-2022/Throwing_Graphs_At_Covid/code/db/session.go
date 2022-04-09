package db

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

// Session provides a convenience run function for nieve error handling.
type Session struct {
	neo4j.Session
}

// Run a statement, panicking on any errors.
func (s *Session) Run(stmt string, p map[string]interface{}) neo4j.Result {
	r, err := s.Session.Run(stmt, p)

	if err != nil {
		panic(err)
	}

	return r
}
