package main

import (
	"fmt"

	exported "github.com/Usagi1234/go_basic101/Title"
)

func main() {

	//* Exported names
	fmt.Println("basic101")
	exported.Exported() // เรื่องการทำ Exported names หรือก็คือการดึง package มาใช้งาน
	//*การสร้างตัวแปรขึ้นต้นด้วยตัวพิมใหญ่จะสามารถส่งออกมาใช้งานได้

	// exported.Exported2()
}
