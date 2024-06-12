package main

import "fmt"

func main() {
	salaries := map[string]int{"Eric": 1000, "Ana": 2000, "João": 3000}
	//fmt.Println(salaries["Eric"])
	delete(salaries, "João")
	salaries["Name"] = 5000
	//fmt.Println(salaries["Name"])

	//sal := map[string]int{} // Cria um map vazio
	sal := make(map[string]int) // Cria um map vazio
	fmt.Println(sal)

	// O _ ignora o valor, se chama blank identifier
	for _, salary := range salaries {
		fmt.Printf("Salary is %d\n", salary)
	}
}
