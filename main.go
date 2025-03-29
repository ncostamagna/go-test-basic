package main

import (
	"fmt"
	"github.com/ncostamagna/go-test-basic/internal/logic"
)

func main() {
	total := logic.Sum(1, 2, 3, 4, 5)
	println(total)

	fmt.Println(logic.CheckEmail("test@test.com"))
	fmt.Println(logic.CheckEmail("nahuel@gmail.com.ar"))
	fmt.Println(logic.CheckEmail("test123@test.ai"))
	fmt.Println(logic.CheckEmail("   test@test.com"))
	fmt.Println(logic.CheckEmail("test"))
	fmt.Println(logic.CheckEmail("test@.  test.com"))
	fmt.Println(logic.CheckEmail("test1test.com"))
}
