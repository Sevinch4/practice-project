package controller

import (
	"fmt"
	"github.com/google/uuid"
	"sort"
)

func (c Controller) Sum() {
	products, err := c.Store.ProductRepo.GetList()
	if err != nil {
		fmt.Println("error is while get product list", err)
	}

	categories, err := c.Store.CategoryRepo.GetList()
	if err != nil {
		fmt.Println("error is while get category list", err)
	}

	catMap := make(map[uuid.UUID]string)
	sumMap := make(map[uuid.UUID]int)

	for _, c := range categories {
		catMap[c.ID] = c.Name
	}
	for _, p := range products {
		sumMap[p.Category_id] += p.Price
	}

	type Count struct {
		Key   uuid.UUID
		Value int
	}

	var counts = []Count{}

	for i, v := range sumMap {
		counts = append(counts, Count{i, v})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})

	counts = counts[:1]

	for _, v := range counts {
		fmt.Println("Category Name: ", catMap[v.Key], " -> ", v.Value)
	}

}
