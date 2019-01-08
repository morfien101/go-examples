package main

import "fmt"

func main() {
	peeps := map[string]map[string]string{
		"person1": {
			"age": "30",
			"sex": "male",
		},
		"Person2": {
			"age": "34",
			"sex": "female",
		},
		"Person3": {
			"age": "7",
			"sex": "male",
		},
		"Person4": {
			"age": "11",
			"sex": "male",
		},
	}

	fmt.Println(peeps["person1"])

	slic := make([]string, 0)

	for key, value := range peeps {
		fmt.Println("The key is", key)
		fmt.Println("Age:", value["age"], "Sex:", value["sex"])
		fmt.Println()
		slic = append(slic, key)
	}

	fmt.Println()
	fmt.Println(slic)
}
