package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var order string
	var num int
	var sales = []int{}
	var counter int = 0
	for true {
		fmt.Println("-------------XYZ CAFE-------------------")
		fmt.Println("---------------MENU---------------------")
		fmt.Println("========================================")
		fmt.Println("C:  coffee           40rs/-")
		fmt.Println("D : dosa             80rs/-")
		fmt.Println("T : tomato soup      20rs/-")
		fmt.Println("I : idli             20rs/-")
		fmt.Println("V : vada             25rs/-")
		fmt.Println("B : bhature &chhole  30rs/-")
		fmt.Println("P : paneer pakoda    30rs/-")
		fmt.Println("M : manchurian       90rs/-")
		fmt.Println("H : hakka noodle     70rs/-")
		fmt.Println("F : French fries     30rs/-")
		fmt.Println("J : jalebi           30rs/-")
		fmt.Println("L;  lemonade         15rs/-")
		fmt.Println("S:  spring roll      20rs/-")
		fmt.Println("=======================================")
		fmt.Println("END for closing")
		fmt.Println("Please place the order: ")
		fmt.Scan(&order)
		order = strings.ToUpper(order)

		if order == "END" {
			total_earning(sales)
		} else {
			fmt.Println("quantity of food is:")
			fmt.Scan(&num)
		}
		bill := bill(num, order)
		fmt.Println("--------------------")
		fmt.Println("Total bill: ", bill)
		sales = append(sales, bill)
		counter++

	}
}

func total_earning(sale []int) {
	income := 0

	for i := 0; i < len(sale); i++ {
		income = sale[i] + income
	}
	fmt.Println("Total earnings for the day is : ", income)
	os.Exit(0)
}

func bill(quantity int, order string) int {
	m := map[string]int{
		"C": 40,
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
