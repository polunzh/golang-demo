package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.12903
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map assigned at \"ten\" is: %d\n", mapLit["ten"])

	_, ok := mapLit["one"]
	fmt.Println(ok)
	delete(mapLit, "one")
	_, ok = mapLit["one"]
	fmt.Println(ok)

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}
