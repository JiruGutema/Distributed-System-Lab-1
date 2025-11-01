package main
import "fmt"

func main() {
	// Declare variables with explicit types
	var a int = 100
	var b float64 = 50
	var c string = "Jiru Gutema"
	var d bool = true

	// Short variable declaration (type inferred)
	e:= "Short Declaration"

	// Print the values of the variables
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)
	fmt.Println("Short Declared String:", e)
}

