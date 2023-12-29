package controller

import (
	"fmt"
	"github.com/google/uuid"
	"market/models"
)

func (c Controller) InsertProduct() {
	product := GetProductInfo()

	if err := c.Store.ProductRepo.Insert(product); err != nil {
		fmt.Println("error is while inserting", err.Error())
		return
	}

	fmt.Println("product added!")
}

func (c Controller) GetProductByID() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error is while parsing id", err.Error())
		return
	}
	product, err := c.Store.ProductRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err)
		return
	}
	fmt.Println("product: ", product)

}

func (c Controller) GetProductList() {
	products, err := c.Store.ProductRepo.GetList()
	if err != nil {
		fmt.Println("error is while get list", err.Error())
		return
	}
	fmt.Println("products: ", products)
}

func GetProductInfo() models.Product {
	var (
		name        string
		category_id string
		price       int
	)
	fmt.Print("input name: ")
	fmt.Scan(&name)

	fmt.Print("input price: ")
	fmt.Scan(&price)

	fmt.Print("input category id: ")
	fmt.Scan(&category_id)

	return models.Product{
		Name:        name,
		Price:       price,
		Category_id: uuid.MustParse(category_id),
	}
}
