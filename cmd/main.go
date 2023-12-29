package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"market/config"
	"market/controller"
	"market/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		fmt.Println("error is while connecting to db", err.Error())
		return
	}
	defer store.DB.Close()

	con := controller.New(store)

	//con.CreateCategory()
	//con.GetCategoryByID()
	//con.GetCategoryList()

	//con.InsertProduct()
	//con.GetProductByID()
	//con.GetProductList()

	con.Sum()
}
