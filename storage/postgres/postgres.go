package postgres

import (
	"database/sql"
	"fmt"
	"market/config"
)

type Store struct {
	DB           *sql.DB
	CategoryRepo categoryRepo
	ProductRepo  productRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf("host = %s port = %s user = %s password = %s database = %s sslmode = disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	category := NewCategory(db)
	product := NewProduct(db)

	return Store{
		DB:           db,
		CategoryRepo: category,
		ProductRepo:  product,
	}, nil
}
