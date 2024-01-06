package main

//With defer

func ReadWrite() bool {
	file.Open("file")
	defer file.Close() //file.Close() is added to defer list
	// Do your thing
	if failureX {
		return false // Close() is now done automatically
	}

	if failureY {
		return false // And here too
	}

	return true // And here
}
