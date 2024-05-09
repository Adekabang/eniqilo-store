package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Adekabang/eniqilo-store/model"
	"github.com/google/uuid"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepository{Db: db}
}

func (m *ProductRepository) AddProduct(payload model.AddUpdateProduct) bool {

	uuidProduct := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO product (id, name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidProduct, payload.Name, payload.Sku, payload.Category, payload.Category, payload.Notes, payload.Price, payload.Stock, payload.Location, payload.IsAvailable)
	if err2 != nil {
		log.Println(err2)
		return false
	}

	return true
}

func (m *ProductRepository) GetProduct(payload model.ParamsGetProduct) bool {

	rows, err := m.Db.Query("SELECT * FROM product")
	if err != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()

	log.Println("rows", rows)

	var products []model.ResponseGetProduct
	if rows != nil {
		for rows.Next() {
			var (
				id        string
				name      string
				createdAt string
			)

			err := rows.Scan(&id, &name, &createdAt)
			if err != nil {
				log.Println(err)
				continue
			}

			// Parse the custom format into a slice of strings
			createdAtFormated, err := time.Parse(time.RFC3339, createdAt)
			if err != nil {
				fmt.Println(err)
			}

			cat := model.ResponseGetProduct{Id: id, Name: name, CreatedAt: createdAtFormated.String()}
			products = append(products, cat)
		}
	}

	log.Println("products", products)

	return true
}

func (m *ProductRepository) UpdateProduct(payload model.AddUpdateProduct) bool {

	log.Println("payload", payload)

	return true
}

func (m *ProductRepository) DeleteProduct(id model.ProductUri) bool {

	log.Println("id", id)

	return true
}
