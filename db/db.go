package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var DB *sql.DB

func main() {
	initializeDBConnection()
	http.HandleFunc("/", getCustomerWithGetEndpoint)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getCustomerWithGetEndpoint(w http.ResponseWriter, r *http.Request) {
	customers, err := getCustomers()
	if err != nil {
		fmt.Println(err)
	}
	marshal, err := json.Marshal(customers)
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(string(marshal))
}

func initializeDBConnection() {
	config := NewEnvDbConfig()
	_, err := ConnectToDb(config)
	if err != nil {
		fmt.Println(err)
	}
}

type EnvDBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewEnvDbConfig() *EnvDBConfig {
	return &EnvDBConfig{
		host:     "localhost",
		port:     "3306",
		username: "root",
		password: "Meliksahse28",
		database: "securesales",
	}
}

func (c *EnvDBConfig) GetHost() string {
	return c.host
}

func (c *EnvDBConfig) GetPort() string {
	return c.port
}

func (c *EnvDBConfig) GetUsername() string {
	return c.username
}

func (c *EnvDBConfig) GetPassword() string {
	return c.password
}

func (c *EnvDBConfig) GetDatabase() string {
	return c.database
}

func ConnectToDb(config *EnvDBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.GetUsername(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	db, err := sql.Open("mysql", connectionString)
	DB = db
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

func getCustomers() ([]Customer, error) {
	rows, err := DB.Query("Select * from customer")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Address, &customer.CreatedAt, &customer.Email, &customer.ImageUrl, &customer.Name, &customer.Phone, &customer.Status, &customer.Type)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

type Customer struct {
	Id        int    `json:"id"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	Email     string `json:"email"`
	ImageUrl  string `json:"imageUrl"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	Type      string `json:"type"`
}
