package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
)
func dbConn() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "employees",
	}
	// Get a database handle.
	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

// albumsByArtist queries for albums that have the specified artist name.
// func albumsByArtist(name string) ([]Album, error) {
//     // An albums slice to hold data from returned rows.
//     var albums []Album
// 
//     rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
//     if err != nil {
//         return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//     }
//     defer rows.Close()
//     // Loop through rows, using Scan to assign column data to struct fields.
//     for rows.Next() {
//         var alb Album
//         if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
//             return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//         }
//         albums = append(albums, alb)
//     }
//     if err := rows.Err(); err != nil {
//         return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//     }
//     return albums, nil
// }
func startGraphql() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}
func main() {
	db:= dbConn()

	fmt.Print("Hello World")
}