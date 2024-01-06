package main

import "fmt"

//Anonymous function accessing variable on each iteration of loop inside function body.

func main() {
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}
