package card

import "math/rand"

func card() int {
	num := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	ren := make([]string, 6)
	for i := 0; i < 5; i++ {
		renn := num[rand.Intn(len(num))]
		ren = append(ren, renn)
	}
}