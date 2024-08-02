package main

import (
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	s := `
		type Student {
			id: Int!
			name: String!
			totalCourses: Int!
		}
		type School {
			name: String!
			address: String!
			students: [Student!]!
		}
		type Query {
			name: String!
			school: School
		}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
