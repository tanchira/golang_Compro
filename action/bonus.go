package main

import "fmt"

func main() {
	want := 50000
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	fmt.Println("			Report  Bonus			")
	fmt.Println("***********************************")
	fmt.Println("  No (ลำดับพนักงาน)	= 			 " ,:)
	fmt.Println("  Name (ชื่อพนักงาน)	 =	 Bam	  ")
	fmt.Println("  Summit (ยอดขาย)     =      ", want)
	fmt.Println("***********************************")
	if want >= 50000 {
		fmt.Println(" Bonus (โบนัส)         =      ", Hugesales)
	} else {
		fmt.Println(" Bonus (โบนัส)         =      ", sales)
	}
}
