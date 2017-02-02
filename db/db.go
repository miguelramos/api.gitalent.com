package db

import (
	"fmt"
	"log"
	"os"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/joho/godotenv"
)

type Neo struct {
	Db bolt.Conn
}

func (n *Neo) Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	url := fmt.Sprintf("bolt://%s:%s", host, port)

	driver := bolt.NewDriver()
	n.Db, _ = driver.OpenNeo(url)
	//defer n.Db.Close()
}
