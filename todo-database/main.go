package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID            int
	Description   string
	CreatedDate   time.Time
	CompletedDate *time.Time
}

func main() {
	var (
		user, password, host, dbName, sslMode, todo, action string
	)
	var port int
	flag.StringVar(&user, "user", "", "Postgresql username")
	flag.StringVar(&password, "password", "", "Postgresql password")
	flag.StringVar(&host, "host", "", "Postgresql host")
	flag.IntVar(&port, "port", 0, "Postgresql port")
	flag.StringVar(&dbName, "dbname", "", "Postgresql host")
	flag.StringVar(&todo, "todo", "", "Todo item")
	flag.StringVar(&action, "action", "create", "Action for todo item")
	flag.StringVar(&sslMode, "sslmode", "disable", "Postgresql host")
	flag.Parse()
	db, err := connect(user, password, host, dbName, sslMode, port)
	if err != nil {
		panic(err)
	}
	switch action {
	case "create":
		err := create(db, todo)
		if err != nil {
			panic(err)
		}
	case "list":
		todos, err := list(db)
		if err != nil {
			panic(err)
		}
		for _, t := range todos {
			fmt.Printf("- %q\t%s\n", t.Description, t.CreatedDate.String())
		}
	default:
		panic("Supported actions are: list, create")
	}
}
func connect(user, password, host, dbName, sslMode string, port int) (*sql.DB, error) {
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + fmt.Sprintf("%d", port) + "/" + dbName + "?sslmode=" + sslMode
	db, err := sql.Open("postgres", connStr)
	fmt.Println("Successfully connected")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func create(db *sql.DB, todo string) error {
	res, err := db.Query(`INSERT INTO
todo
(id, description, created_date)
VALUES
(DEFAULT, $1, now());`,
		todo)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}
func list(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query(`SELECT
id,
description,
created_date,
completed_date
FROM todo;`)
	if err != nil {
		return []Todo{},
			fmt.Errorf("unable to get from todo table: %w", err)
	}
	todos := []Todo{}
	defer rows.Close()
	for rows.Next() {
		result := Todo{}
		err := rows.Scan(&result.ID,
			&result.Description,
			&result.CreatedDate,
			&result.CompletedDate)
		if err != nil {
			return []Todo{}, fmt.Errorf("row scan error: %w", err)
		}
		todos = append(todos, result)
	}
	return todos, nil
}
