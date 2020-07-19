package main

import "fmt"

func main() {
	// 1st way to initliaze map
	colors := map[string]string{
		"red":   "Laal",
		"green": "Hirwa",
		"blue":  "Nila",
		"Pink":  "Gulabi",
	}
	colors["Purple"] = "Jambhla"
	fmt.Println(colors)
	// 2nd way to initialize map
	var mob map[string]int64
	//mob["Shreyas"] = 406268
	fmt.Println(mob)
	// 3rd way to
	colu := make(map[string]string)
	colu["kala"] = "black"
	fmt.Println(colu)

	// deleting the map key value pair
	delete(colors, "green")
	fmt.Println(colors)
	// Iterate over the map

	iterate(colors)
}

func iterate(m map[string]string) {
	for i, j := range m {
		fmt.Println(i, "->", j)
	}
}
