package repository

import (
	"Go_Projects/api/entity"
	"database/sql"
	"fmt"
)

type CityRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (r CityRepo) Insert(city entity.City) {
	statement, err := r.db.Prepare("INSERT INTO cities(name,code) values($1,$2)")

	result, err := statement.Exec(city.Name, city.Code)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result.RowsAffected())
	}
}

func (r CityRepo) List() []entity.City {
	var cityList []entity.City

	rows, err := r.db.Query("Select id,name,code From cities")
	if err != nil {
		fmt.Println(err)
		return cityList
	} else {
		for rows.Next() {
			var city entity.City
			err := rows.Scan(&city.Id, &city.Name, &city.Code)
			if err != nil {
				fmt.Println(err)
			} else {
				cityList = append(cityList, city)
			}
		}
		rows.Close()

		return cityList
	}
}

func (r CityRepo) GetByName(name string) *entity.City {
	statement, err := r.db.Prepare("Select id,name,code from cities where name=$1")
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		var city entity.City
		statement.QueryRow(name).Scan(&city.Id, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err)
		}

		return &city
	}
}

func (r CityRepo) GetById(id int) *entity.City {
	statement, err := r.db.Prepare("Select id,name,code from cities where id=$1")
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		var city entity.City
		err = statement.QueryRow(id).Scan(&city.Id, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		return &city
	}
}
