package config

import (
	"log"
	"os"
)

type Neo4jConfig struct {
	URI      string
	Username string
	Password string
}

func NewNeo4jConfig() *Neo4jConfig {
	uri := os.Getenv("NEO4J_URI")
	if uri == "" {
		uri = "neo4j://localhost:7687"
	}

	username := os.Getenv("NEO4J_USERNAME")
	if username == "" {
		username = "neo4j"
	}

	password := os.Getenv("NEO4J_PASSWORD")
	if password == "" {
		password = "password"
	}

	log.Printf("Neo4jConfig: %+v", Neo4jConfig{
		URI:      uri,
		Username: username,
		Password: password,
	})

	return &Neo4jConfig{
		URI:      uri,
		Username: username,
		Password: password,
	}
}
