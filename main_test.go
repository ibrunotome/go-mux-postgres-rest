package main_test

import (
	"os"
	"testing"

	"."
	"log"
	"github.com/joho/godotenv"
)

var a main.App

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a = main.App{}
	a.Initialize(
		os.Getenv("TEST_APP_DB_USERNAME"),
		os.Getenv("TEST_APP_DB_PASSWORD"),
		os.Getenv("TEST_APP_DB_NAME"),
		os.Getenv("TEST_AP_DB_PORT"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
id SERIAL,
name TEXT NOT NULL,
price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
CONSTRAINT products_pkey PRIMARY KEY (id)
)`
