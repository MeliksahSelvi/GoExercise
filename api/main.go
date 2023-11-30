package main

import (
	"Go_Projects/api/brokers"
	"Go_Projects/api/cache"
	"Go_Projects/api/entity"
	"Go_Projects/api/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"strconv"
)

var (
	db    *sql.DB
	dbErr error
)

func main() {
	connectDB()

	cityRepo := repository.NewRepo(db)
	rabbitMQ := brokers.NewRabbitMQ()
	redis := cache.NewRedis()
	http.HandleFunc("/city/findAll", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			http.Error(writer, "unsupported method", http.StatusMethodNotAllowed)
			return
		}

		cityCache := redis.Get()
		if cityCache != nil {
			fmt.Println("Cities cache de var")
			writer.Write(cityCache)
			return
		}

		fmt.Println("Cities cache de yok")
		cityList := cityRepo.List()
		cityListBytes, _ := json.Marshal(cityList)

		go func(data []byte) {
			fmt.Println("Cities cache put edildi")
			redis.Put(data)
		}(cityListBytes)

		writer.Write(cityListBytes)
	})

	http.HandleFunc("/city/findById", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			http.Error(writer, "unsupported method", http.StatusMethodNotAllowed)
			return
		}

		cityIdStr := request.URL.Query().Get("id")
		cityId, _ := strconv.Atoi(cityIdStr)
		city := cityRepo.GetById(cityId)
		if city == nil {
			http.Error(writer, "not found", http.StatusNotFound)
			return
		} else {
			fmt.Println("girdi :(")
			json.NewEncoder(writer).Encode(city)
		}
	})

	http.HandleFunc("/city/save", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(writer, "unsupported method", http.StatusMethodNotAllowed)
			return
		}

		var city entity.City
		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			fmt.Println("hatabadrequest")
			return
		}

		if err = json.Unmarshal(bodyBytes, &city); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		cityRepo.Insert(city)
		rabbitMQ.Publish([]byte("city created with name: " + city.Name))
		writer.WriteHeader(http.StatusCreated)
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println(err)
		}
	}()

	<-make(chan struct{})
}

func connectDB() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "1234"
	dbName := "godb"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, dbErr = sql.Open("postgres", psqlInfo)

	if dbErr != nil {
		panic(dbErr)
	}
}
