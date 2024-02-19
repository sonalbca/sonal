package main

//Without defer

func ReadWrite() bool {
	file.Open("file")

	if failureX {
		file.Close() //And here...
		return false
	}
	if failureY {
		file.Close() //And here...
		return false
	}
	file.Close() //And here...
	return true
}

//A lot of code is repeated here. To overcome this Go has the defer statement.
//The code above could be rewritten as follows. This makes the function more readable,
//shorter and puts the Close right next to the Open.
