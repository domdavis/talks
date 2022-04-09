package main

import (
	"fmt"
	"math/rand"

	"bitbucket.org/idomdavis/tgac/db"
	"bitbucket.org/idomdavis/tgac/rng"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const size = 100

func main() {
	c := db.New()
	c.Clear()

	s := c.Session()

	// populate(s)
	// read(s)
	// interact(s)
	model(s)

}

func populate(s *db.Session) {
	//language=cypher
	const stmt = `CREATE (:Person {id:$id})`

	for i := 0; i < size; i++ {
		s.Run(stmt, map[string]interface{}{"id": i})
	}
}

func interact(s *db.Session) {
	//language=cypher
	const stmt = `
		MATCH (s:Person {id: $from}), (o:Person {id: $to})
		CREATE (s)-[:CONTACT {iteration: $i}]->(o)`

	const (
		days        = 28
		probability = 10
	)

	for i := 0; i < size; i++ {
		for j := 0; j < days; j++ {

			if !rng.Chance(probability) {
				continue
			}

			p := map[string]interface{}{"from": i, "to": rng.Select(size, i), "i": j}

			s.Run(stmt, p)
		}
	}
}

func read(s *db.Session) {
	//language=Cypher
	stmt := "MATCH (n) RETURN n.id AS id LIMIT 5"
	res := s.Run(stmt, nil)

	for _, r := range db.Collect(res) {
		fmt.Println(r.Get("id"))
	}
}

func model(s *db.Session) {
	const (
		population      = 100
		infectionChance = 10
		duration        = 28
	)

	infectedVisit := map[string]int{"shop": 100}
	healthyVisit := map[string]int{"shop": 25, "school": 25, "work": 25, "cafe": 25}
	patientZero := rand.Intn(population)

	places(s, infectedVisit, healthyVisit)

	for i := 0; i < population; i++ {
		//language=cypher
		stmt := `CREATE (:Person {id:$id, infected: $infected})`

		s.Run(stmt, map[string]interface{}{"id": i, "infected": i == patientZero})
	}

	for day := 1; day <= duration; day++ {
		infected := db.Collect(s.Run(`MATCH (p:Person {infected: true}) return p.id as id`, nil))
		healthy := db.Collect(s.Run(`MATCH (p:Person {infected: false}) return p.id as id`, nil))

		for _, i := range infected {
			id, _ := i.Get("id")

			visit(s, day, int(id.(int64)), infectedVisit)
		}

		for _, h := range healthy {
			id, _ := h.Get("id")

			visit(s, day, int(id.(int64)), healthyVisit)
		}

		for _, v := range visited(s, day) {
			if !rng.Chance(infectionChance) {
				continue
			}

			id, _ := v.Get("id")

			s.Run(`MATCH (p:Person {id: $id}) set p.infected = true`,
				map[string]interface{}{"id": int(id.(int64))})
		}
	}
}

func visited(s *db.Session, day int) []*neo4j.Record {
	//language=Cypher
	stmt := `
		MATCH (:Person {infected: true})-[:VISITED {day: $day-1}]->(l:Location)
			<-[:VISITED {day:$day}]-(p:Person {infected: false})
		RETURN DISTINCT(p.id) AS id
`
	return db.Collect(s.Run(stmt, map[string]interface{}{"day": day}))
}

func visit(s *db.Session, day, person int, places map[string]int) {
	for place, chance := range places {
		if rng.Chance(chance) {
			//language=cypher
			stmt := `
					MATCH (p:Person {id: $id})
					WITH p 
					MATCH (l:Location {name: $location})
					WITH p, l
					MERGE (p)-[:VISITED {day: $day}]->(l)`
			params := map[string]interface{}{
				"id": person, "location": place, "day": day,
			}

			s.Run(stmt, params)
		}
	}
}

func places(s *db.Session, maps ...map[string]int) {
	locations := map[string]struct{}{}

	for _, m := range maps {
		for k := range m {
			if _, ok := locations[k]; ok {
				continue
			}

			s.Run(`CREATE (:Location {name: $shop})`, map[string]interface{}{"shop": k})

			locations[k] = struct{}{}
		}
	}
}
