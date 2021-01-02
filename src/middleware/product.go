package middleware

import (
	"encoding/json"
	"fmt"
	"microservices/src/config"
	"microservices/src/dao"
	"microservices/src/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productDao := dao.ProductDao{
			Db: db,
		}
		products, err := productDao.FindAll()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(products)
			// fmt.Println("Product List")
			// for _, product := range products {
			// 	fmt.Println("Id:", product.Id)
			// 	fmt.Println("Name:", product.Name)
			// 	fmt.Println("Price:", product.Price)
			// 	fmt.Println("Quantity:", product.Quantity)
			// 	fmt.Println("Status:", product.Status)
			// }
		}
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productDao := dao.ProductDao{
			Db: db,
		}
		product, err := productDao.Find(params["id"])
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
			// fmt.Println(product)
			// fmt.Println("Product Info")
			// fmt.Println("Id:", product.Id)
			// fmt.Println("Name:", product.Name)
			// fmt.Println("Price:", product.Price)
			// fmt.Println("Quantity:", product.Quantity)
			// fmt.Println("Status:", product.Status)
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productDao := dao.ProductDao{
			Db: db,
		}
		err := productDao.Create(&product)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			// fmt.Println("Lastest Id:", product.Id)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product.Id)
		}
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productDao := dao.ProductDao{
			Db: db,
		}
		rows, err := productDao.Update(product)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			if rows > 0 {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode("Done")
			}
		}
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		productDao := dao.ProductDao{
			Db: db,
		}
		rows, err := productDao.Delete(params["id"])
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			if rows > 0 {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode("Done")
			}
		}
	}
}
