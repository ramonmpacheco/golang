package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4 %s %s\n", p1, p2)
}

func f5() (string, string) {
	return "R1", "R2"
}

func main() {
	f1()
	f2("param1", "param2")

	r3, r4 := f3(), f4("param1", "param2")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println(r51, r52)

}
