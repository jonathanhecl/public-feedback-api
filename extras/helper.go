package extras

import (
	"math/rand"
	"time"
)

func RandomCode() string {

	rand.Seed(time.Now().UnixNano())
	const numBytes = "0123456789ABCDEF"
	b := make([]byte, 8)
	for i := range b {
		b[i] = numBytes[rand.Intn(len(numBytes))]
	}
	return string(b)

}

func GetWebDomain() string {

	return ex.webDomain

}

func GetAPIDomain() string {

	return ex.apiDomain

}
