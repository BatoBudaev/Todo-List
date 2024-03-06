package db

import (
	"database/sql"
	"fmt"
	"github.com/BatoBudaev/Todo-List/internal/models"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sql.DB
}

func InitDB(user, password, dbname string) (*DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable client_encoding=UTF8", user, password, dbname)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			description VARCHAR(255) NOT NULL,
			completed BOOLEAN NOT NULL
		);
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}, nil
}

func (db *DB) CreateTask(task models.Task) (models.Task, error) {
	var id int
	err := db.QueryRow(`INSERT INTO tasks(description, completed) VALUES($1, $2) RETURNING id`, task.Description, task.Completed).Scan(&id)

	if err != nil {
		return models.Task{}, err
	}

	task.ID = id
	return task, nil
}

func (db *DB) GetTasks() ([]models.Task, error) {
	rows, err := db.Query(`SELECT * FROM tasks`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task

		if err := rows.Scan(&t.ID, &t.Description, &t.Completed); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
