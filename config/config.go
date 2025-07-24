package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=arthur password=1234 dbname=user_db sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao verificar conexão:", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
