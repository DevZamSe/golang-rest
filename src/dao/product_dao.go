package dao

import (
	"database/sql"
	"microservices/src/entities"
)

type ProductDao struct {
	Db *sql.DB
}

func (productDao ProductDao) Create(product *entities.Product) error {
	result, err := productDao.Db.Exec("insert into product(name, price, quantity, status) values(?,?,?,?)",
		product.Name, product.Price, product.Quantity, product.Status)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

func (productDao ProductDao) Update(product entities.Product) (int64, error) {
	result, err := productDao.Db.Exec("update product set name = ?, price = ?, quantity = ?, status = ? where id = ?",
		product.Name, product.Price, product.Quantity, product.Status, product.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productDao ProductDao) Find(id string) (entities.Product, error) {
	rows, err := productDao.Db.Query("select * from product where id = ?", id)
	if err != nil {
		return entities.Product{}, err
	} else {
		product := entities.Product{}
		for rows.Next() {
			var id int64
			var name string
			var price float32
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return entities.Product{}, err2
			} else {
				product = entities.Product{id, name, price, quantity, status}
			}
		}
		return product, nil
	}
}

func (productDao ProductDao) FindAll() ([]entities.Product, error) {
	rows, err := productDao.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{}
		for rows.Next() {
			var id int64
			var name string
			var price float32
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{id, name, price, quantity, status}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productDao ProductDao) Delete(id string) (int64, error) {
	result, err := productDao.Db.Exec("delete from product where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
