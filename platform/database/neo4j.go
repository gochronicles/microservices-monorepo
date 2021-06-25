package database

import (
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var Driver neo4j.Driver

// OpenDBConnection func for opening database connection.
func Neo4jConnection() (neo4j.Driver, error) {
	// Define a new PostgreSQL connection.
	uri := os.Getenv("NEO4J_URI")
	username := os.Getenv("NEO4J_USERNAME")
	password := os.Getenv("NEO4J_PASSWORD")
	fmt.Println(uri, username, password)

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	sess := driver.NewSession(neo4j.SessionConfig{})
	defer sess.Close()

	return driver, nil

}

func Write(function func(neo4j.Transaction) (interface{}, error)) (interface{}, error) {
	sess := Driver.NewSession(neo4j.SessionConfig{})
	defer sess.Close()
	result, err := sess.WriteTransaction(function)
	return result, err
}

func Read(function func(neo4j.Transaction) (interface{}, error)) (interface{}, error) {
	sess := Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer sess.Close()
	result, err := sess.ReadTransaction(function)
	return result, err

}

func init() {
	var err error
	Driver, err = Neo4jConnection()
	if err != nil {
		panic(err)
	}
}
