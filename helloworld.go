package main

import "fmt"

// data types
/*
1. **Numeric Types:**
   - `int`: Signed integers (platform dependent size).
   - `int8`, `int16`, `int32`, `int64`: Signed integers of specific sizes.
   - `uint`: Unsigned integers (platform dependent size).
   - `uint8` (alias `byte`), `uint16`, `uint32`, `uint64`, `uintptr`: Unsigned integers of specific sizes.
   - `float32`, `float64`: Floating-point numbers.
   - `complex64`, `complex128`: Complex numbers with float32 and float64 real and imaginary parts respectively.

2. **Boolean Type:**
   - `bool`: Represents boolean values `true` or `false`.

3. **String Types:**
   - `string`: Represents a sequence of bytes or characters.

4. **Composite Types:**
   - `array`: Fixed-size collection of items of the same type.
   - `slice`: Variable-size sequence (like a dynamic array).
   - `map`: Unordered collection of key-value pairs.
   - `struct`: Collection of fields with named values.
   - `channel`: Communication channel for goroutines.
   - `interface`: Defines a set of methods that a concrete type must implement to be considered as implementing the interface.
   - `function`: A callable function type.

5. **Pointer Types:**
   - `pointer`: Holds the memory address of a value.

6. **Derived Types:**
   - `slice`, `map`, `channel`, `function`, `interface`: These are considered derived types as they are built using the basic types.

7. **Special Types:**
   - `nil`: Represents a zero value or an uninitialized value for pointers, functions, interfaces, slices, channels, and maps.
*/

func main() {
	// Hello world
	fmt.Println("Hello world from go")

	// explicit declaration
	var number int32 = 3000
	fmt.Println(number, "explicit using var")

	// implicit declaration
	var number2 = 3000
	fmt.Println(number2, "implicit using var")

	// Variables cannot be declared without var keyword,
	// but it can be don3 with walrus operator to do implicit declaration
	number3 := 3000
	fmt.Println(number3, "implicit using walrus operator")

	// The below can be done:

	/* var number int64 = 3000
	   var number = 3000
	   var number = int64(2000)
	   number := 3000
	   number := int64(3000)
	*/

	// The below cannot be done:

	/* var number := 3000 [We cannot use var and the walrus operator in a same declaration],
	   number = 3000 [We cannot declare a variable without var keyword, in this case
	                 we have to use walrus operator [:=] to declare variables without var keyword] */
}
