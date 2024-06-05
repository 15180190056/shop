package dbserver

import (
	"errors"
	"mystopproject/conf"
	"strings"
)

func CreateProduct(producttable *conf.Product) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	DB.Create(&producttable)
	return err
}

func GetProductByName(productname string) (*conf.Product, error) {
	var Product conf.Product
	DB, err := InitDB()
	if err != nil {
		return nil, err
	}
	DB.Find(&Product, "product_name=?", productname)
	return &Product, nil // 用户存在，返回用户指针和nil错误
}

func UpdateProduct(product *conf.Product) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	DB.Table("products").Where("product_name = ?", product.ProductName).Updates(product)
	return nil
}

func DeleteProductByName(productname string) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	result := DB.Where("product_name = ?", productname).Delete(&conf.Product{})
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "record not found") {
			return errors.New("商品不存在") // 用户不存在，返回自定义错误
		}
		return result.Error
	}
	return nil
}
