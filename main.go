package main

import "fmt"

func main() {
	transactions := []float64{}
	var result float64
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}

	for _, value := range transactions {
		result += value
	}
	fmt.Println(transactions, result)
}

func scanTransaction() float64 {
	fmt.Print("Введите транзакцию (n для выхода): ")
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}
