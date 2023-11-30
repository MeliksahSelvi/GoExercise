package service

import "Go_Projects/api/entity"

type CityService struct {
	Repo CityRepo
}

func NewCityService(repo CityRepo) *CityService {
	return &CityService{
		Repo: repo,
	}
}

func (c CityService) SaveCity(city entity.City) {
	c.Repo.Insert(city)
}

func (c CityService) GetAllCities() []entity.City {
	return c.Repo.List()
}

func (c CityService) GetCityByName(name string) *entity.City {
	return c.Repo.GetByName(name)
}

func (c CityService) GetCityById(id int) *entity.City {
	return c.Repo.GetById(id)
}

type CityRepo interface {
	Insert(entity.City)
	List() []entity.City
	GetByName(name string) *entity.City
	GetById(id int) *entity.City
}
