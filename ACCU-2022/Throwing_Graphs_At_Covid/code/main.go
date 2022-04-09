package main

import "bitbucket.org/idomdavis/tgac/db"

func main() {
	const population = 100

	c := db.New()
	c.Clear()

	s := c.Session()

	for i := 0; i < population; i++ {
		//language=cypher
		stmt := `CREATE (:Person {id: $id})`
		
		s.Run(stmt, map[string]interface{}{"id": i})
	}

	c.Close()
}
