package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var order string
	var quantity int
	var sales = []int{}
	var counter int = 0
	for true {
		fmt.Println("---------XYZ CAFE ----------")
		fmt.Println("DRINKS")
		fmt.Print("  C: coffee, 40rs\n  t:tea,20rs\n  L: lemonade,15rs\n  T:tomato soup,20rs\n")
		fmt.Println("SNACKS AND TIFFINS")
		fmt.Print(" D: dosa, 80 rs\n I : idli , 20rs\n V : vada, 25rs.\n B: bature(cholae), 30rs\n P: paneer pakoda, 30rs\n M: manchuria, 90rs\n H: hakka noodle, 70rs.\n F: French fries, 30rs\n J: jalebi, 30rs\n L:  S: spring roll, 20rs\n")
		fmt.Println("-----------------------")
		fmt.Println("Press END For closing the day.")
		fmt.Println("Please place the order: ")
		fmt.Scan(&order)
		order = strings.ToUpper(order)

		if order == "END" {
			total_income(sales)
		} else {
			fmt.Scan(&quantity)
		}
		bill := bill(quantity, order)
		fmt.Println("####################")
		fmt.Println("Total bill: ", bill)
		sales = append(sales, bill)
		counter++

	}
}

func total_income(sale []int) {
	sum := 0

	for i := 0; i < len(sale); i++ {
		sum = sale[i] + sum
	}
	fmt.Println("Total Income for the day is : ", sum)
	os.Exit(0)
}

func bill(quantity int, order string) int {
	m := map[string]int{
		"C": 40,
		"t": 20,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	totalbill := quantity * m[order]
	return totalbill
}
