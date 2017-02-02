package main

import (
	"fmt"

	"gitalent.com/backend/db"
	uuid "github.com/satori/go.uuid"
)

type HomeHeader []struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Group string `json:"group"`
	UID   string `json:"id"`
}

func Homepage() {
	neo := new(db.Neo)
	neo.Init()

	section := fmt.Sprintf("CREATE (s: Section {label: 'Header Homepage', key: 'home.header', group: 'home', id: {uid} })")

	stmt, err := neo.Db.PrepareNeo(section)
	if err != nil {
		panic(err)
	}

	// Executing a statement just returns summary information
	result, err := stmt.ExecNeo(map[string]interface{}{"uid": uuid.NewV4().String()})
	if err != nil {
		panic(err)
	}
	numResult, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	// Closing the statment will also close the rows
	stmt.Close()

	pipeline, err := neo.Db.PreparePipeline(
		"MATCH (s: Section) WHERE s.key= 'home.header' CREATE (c: Content {label: 'Banner title', text: 'All most good', key: 'banner.title', id: {uid} })-[:BELONGS_TO]->(s)",
		"MATCH (s: Section) WHERE s.key= 'home.header' CREATE (c: Content {label: 'Banner subtile', text: 'All most good', key: 'banner.subtitle', id: {uid} })-[:BELONGS_TO]->(s)",
		"MATCH (s: Section) WHERE s.key= 'home.header' CREATE (c: Content {label: 'Banner text', text: 'All most good', key: 'banner.text', id: {uid} })-[:BELONGS_TO]->(s)",
		"MATCH (s: Section) WHERE s.key= 'home.header' CREATE (c: Content {label: 'Banner images', text: 'image.png:image.png:image.png', key: 'banner.images', id: {uid} })-[:BELONGS_TO]->(s)",
	)
	if err != nil {
		panic(err)
	}

	pipelineResults, err := pipeline.ExecPipeline(map[string]interface{}{"uid": uuid.NewV4().String()}, map[string]interface{}{"uid": uuid.NewV4().String()}, map[string]interface{}{"uid": uuid.NewV4().String()}, map[string]interface{}{"uid": uuid.NewV4().String()})
	if err != nil {
		panic(err)
	}

	for _, result := range pipelineResults {
		numResult, _ := result.RowsAffected()
		fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 2 (per each iteration)
	}

	err = pipeline.Close()
	if err != nil {
		panic(err)
	}

	neo.Db.Close()
}

func main() {
	Homepage()
}
