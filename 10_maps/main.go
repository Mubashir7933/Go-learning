package main

import (
	"fmt"
	"maps"
)

func main() {
	//maps guide

	NewMap := make(map[string]int)

	NewMap["mubashir"] = 23
	NewMap["ali"] = 25

	//fmt.Println(NewMap["mubashir"])

	// //withMap := map[string]int{
	// 	"mubashir": 23,
	// 	"Haseeb":   25,
	// }
	// delete(withMap, "Haseeb")
	// fmt.Println(withMap)

	//checking maps if equal or not using maps function
	// order doesn't matter in maps

	k1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	k2 := map[string]int{
		"two": 2,
		"one": 1,
	}

	fmt.Println(maps.Equal(k1, k2))
}
