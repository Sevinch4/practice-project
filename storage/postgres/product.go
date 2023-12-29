package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"market/models"
)

type productRepo struct {
	DB *sql.DB
}

func NewProduct(db *sql.DB) productRepo {
	return productRepo{DB: db}
}

func (p productRepo) Insert(prod models.Product) error {
	id := uuid.New()
	if _, err := p.DB.Exec(`insert into product (id,name,price,category_id) values($1,$2,$3,$4)`,
		&id, &prod.Name, &prod.Price, &prod.Category_id); err != nil {
		return err
	}

	return nil
}

func (p productRepo) GetByID(id uuid.UUID) (models.Product, error) {
	pr := models.Product{}

	if err := p.DB.QueryRow(`select * from product where id = $1`, id).Scan(&pr.ID, &pr.Name, &pr.Price, &pr.Category_id, &pr.Created_at, &pr.Updated_at); err != nil {
		return models.Product{}, err
	}
	return pr, nil
}

func (p productRepo) GetList() ([]models.Product, error) {
	products := []models.Product{}

	rows, err := p.DB.Query(`select * from product`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		pr := models.Product{}
		if err = rows.Scan(&pr.ID, &pr.Name, &pr.Price, &pr.Category_id, &pr.Created_at, &pr.Updated_at); err != nil {
			return nil, err
		}
		products = append(products, pr)
	}
	return products, nil
}
