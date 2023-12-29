package controller

import (
	"fmt"
	"github.com/google/uuid"
	"market/models"
)

func (c Controller) CreateCategory() {
	category := GetCategoryInfo()

	if err := c.Store.CategoryRepo.InsertCategory(category); err != nil {
		fmt.Println("error is while inserting", err.Error())
		return
	}
	fmt.Println("category added")
}

func (c Controller) GetCategoryByID() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error is while parsing id", err.Error())
		return
	}

	category, err := c.Store.CategoryRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err.Error())
		return
	}
	fmt.Println("category: ", category)
}

func (c Controller) GetCategoryList() {
	categories, err := c.Store.CategoryRepo.GetList()
	if err != nil {
		fmt.Println("error is while get list", err.Error())
		return
	}
	fmt.Println("categories: ", categories)
}

func GetCategoryInfo() models.Category {
	var (
		name string
	)
	fmt.Print("input name: ")
	fmt.Scan(&name)

	return models.Category{
		Name: name,
	}

}
