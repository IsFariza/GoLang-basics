package main

import (
	"github.com/FarizaIsmagambetova/Assignment1/Bank"
	"github.com/FarizaIsmagambetova/Assignment1/Company"
	"github.com/FarizaIsmagambetova/Assignment1/Library"
	"github.com/FarizaIsmagambetova/Assignment1/Shapes"
)

func main() {
	//===LIBRARY===
	Library.LibraryMenu()

	//===SHAPES===
	defaultShapes := Shapes.CreateShapes()
	Shapes.IterateShapes(defaultShapes)
	customShapes := Shapes.CreateCustomShapes()
	Shapes.IterateShapes(customShapes)

	//===COMPANY===
	Company.CompanyMenu()

	//===BANK ACCOUNT===
	Bank.BankAccountMenu()

}
