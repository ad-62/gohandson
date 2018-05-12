package main

import (
	"time"
	"math/rand"
)

func main() {
	omikuji()
}

func omikuji(){
	rand.Seed(time.Now().Unix())
	s := random(1,7)
	result := map[int]string{6:"大吉",5:"中吉",4:"中吉",3:"吉",2:"吉",1:"凶"}
	print(result[s])
}

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

