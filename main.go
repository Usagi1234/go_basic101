package main

import (
	"fmt"

	sqrt "github.com/Usagi1234/go_basic101/if"
)

// functions "github.com/Usagi1234/go_basic101/Functions"
// named "github.com/Usagi1234/go_basic101/Named"
// exported "github.com/Usagi1234/go_basic101/Title"
// "github.com/Usagi1234/go_basic101/multiple"

func main() {

	//! Package, variables และ function
	/*
		//* Exported names
		fmt.Println("basic101")
		exported.Exported() // เรื่องการทำ Exported names หรือก็คือการดึง package มาใช้งาน
		//*การสร้างตัวแปรขึ้นต้นด้วยตัวพิมใหญ่จะสามารถส่งออกมาใช้งานได้

		// exported.Exported2()

		//* Functions
		fmt.Println("Functions")
		r := functions.Sum(5, 5)
		fmt.Println(r)

		//* Multiple
		x, y := multiple.Swap("multi", "ple")
		fmt.Println(x + y)

		//* Named return values
		fmt.Println("Named return values")
		fmt.Println(named.Split(1))

		o, e := named.Split(1)
		fmt.Println("x", o, "y", e)

		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
	*/

	//! For
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//! if
	fmt.Println(sqrt.Sqrt(2), sqrt.Sqrt(-4))

}
