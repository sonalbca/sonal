package main

import (
	"fmt"
	"sort"
)

/*
In this tutorial you will learn what is a map data type and when to use it in Golang.

A map is a data structure that provides you with an unordered collection of
key/value pairs (maps are also sometimes called associative arrays in Php, hash tables in Java, or dictionaries in Python).
Maps are used to look up a value by its associated key. You store values into the map based on a key.

The strength of a map is its ability to retrieve data quickly based on the key.
A key works like an index, pointing to the value you associate with that key.

A map is implemented using a hash table, which is providing faster lookups on the data element and you can easily
retrieve a value by providing the key.Maps are unordered collections, and there's no way to predict the order in which the
key/value pairs will be returned. Every iteration over a map could return a different order.
*/
func main() {
	//Map1()  //Map initialization
	//Map2()  //Empty Map declaration
	//Map3()  //Map declaration using make function
	//Map4()  //Map Length
	//Map5() //Accessing Items
	//Map6() //Adding Items
	//Map7() //Update Values
	//Map8() //Delete Items
	//Map9() //Iterate over a Map
	//Map10() // Truncate Map
	Map11() //Sort Map Keys
	//Map12() //Sort Map Values
	//Map13() //Merge Maps

}

//---------------------------------- 1 -------------------------------//
//Map initialization
//In Golang maps are written with curly brackets, and they have keys and values. Creating an instance of a map data type.

var employee = map[string]int{"Mark": 10, "Sandy": 20}

func Map1() {
	fmt.Println(employee)
}

//---------------------------------- 2 -------------------------------//
//Empty Map declaration
//Map employee created having string as key-type and int as value-type

func Map2() {
	var employee = map[string]int{}
	fmt.Println(employee)        // map[]
	fmt.Printf("%T\n", employee) // map[string]int
}

//---------------------------------- 3 -------------------------------//
//Map declaration using make function
//The make function takes as argument the type of the map and it returns an initialized map.

func Map3() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)

	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)
}

//---------------------------------- 4 -------------------------------//
//Map Length
//To determine how many items (key-value pairs) a map has, use built-in len() function.

func Map4() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20

	// Empty Map
	employeeList := make(map[string]int)

	fmt.Println(len(employee))     // 2
	fmt.Println(len(employeeList)) // 0
}

//The len() function will return zero for an uninitialized map.

//---------------------------------- 5 -------------------------------//
//Accessing Items
//You can access the items of a map by referring to its key name, inside square brackets.

func Map5() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}

	fmt.Println(employee["Mark"])
}

//---------------------------------- 6 -------------------------------//
//Adding Items
//Adding an item to the map is done by using a new index key and assigning a value to it.

func Map6() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map

	employee["Rocky"] = 30 // Add element
	employee["Josef"] = 40

	fmt.Println(employee)
}

//---------------------------------- 7 -------------------------------//
//Update Values
//You can update the value of a specific item by referring to its key name.

func Map7() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map

	employee["Mark"] = 50 // Edit item
	fmt.Println(employee)
}

//---------------------------------- 8 -------------------------------//
//Delete Items
//The built-in delete function deletes an item from a given map associated with the provided key.

func Map8() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	employee["Rocky"] = 30
	employee["Josef"] = 40

	fmt.Println(employee)

	delete(employee, "Mark")
	fmt.Println(employee)
}

//---------------------------------- 9 -------------------------------//
//Iterate over a Map
//The for…range loop statement can be used to fetch the index and element of a map.

//Each iteration returns a key and its correlated element content.

func Map9() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}

//---------------------------------- 10 -------------------------------//
// Truncate Map
// There are two methods to clear all items from a Map.

func Map10() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	// Method - I
	for k := range employee {
		delete(employee, k)
		fmt.Println(employee)
	}

	// Method - II
	employee = make(map[string]int)
}

//---------------------------------- 11 -------------------------------//
//Sort Map Keys
//A Keys slice created to store keys value of map and then sort the slice. The sorted slice used to print values of map in key order.

func Map11() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	keys := make([]string, 0, len(unSortedMap))

	for k, _ := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, unSortedMap[v])
	}
}

//---------------------------------- 12 -------------------------------//
//Sort Map Values
//To sort the key values of a map, you need to store them in Slice and then sort the slice.

func Map12() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	// Int slice to store values of map.
	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Sort slice values.
	sort.Ints(values)

	// Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}
}

//---------------------------------- 13 -------------------------------//
//Merge Maps
//The keys and values of second map getting added in first map.

func Map13() {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)
}
