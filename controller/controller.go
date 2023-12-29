package controller

import "market/storage/postgres"

type Controller struct {
	Store postgres.Store
}

func New(store postgres.Store) Controller {
	return Controller{
		store,
	}
}
