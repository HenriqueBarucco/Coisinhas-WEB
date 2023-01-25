package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConectaComBancoDeDados() *sql.DB {
	//conexao := "developer:123456789@/alura-loja"
	conexao := os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@" + "tcp(" + os.Getenv("HOST") + ")/coisinhas_db"
	fmt.Println(conexao)

	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
