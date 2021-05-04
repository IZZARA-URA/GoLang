package main

import (
	"fmt"
)

func main() {
	fnShow()
}

func fnShow() {
	var product1 Product
	product1.name = "Arduino"
	product1.price = 100
	product1.stock = 20
	show(product1)
	//update(&product1)
	show(product1)
	//product1 = product1.clear()
	//product1 = product1.setDiscount(1)
	show(product1)
}

type Product struct {
	name  string
	price int
	stock int
}

func show(p Product) {
	fmt.Println(p)
}

// Struc and Pionter

func update(p *Product) {
	//fmt.Println(*p)

	p.price = p.price + 100
	p.stock = p.stock + 10
}

// Add Function in Struct

func (p Product) clear() Product {
	p.price = 0
	p.stock = 0
	return p
}

func (p Product) setDiscount(d int) Product {
	p.price = p.price - d
	return p
}
