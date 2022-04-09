package db

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

// Collect all the records from the result, panicking on error.
func Collect(r neo4j.Result) []*neo4j.Record {
	records, err := r.Collect()

	if err != nil {
		panic(err)
	}

	return records
}
