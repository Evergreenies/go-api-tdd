package common

import (
	"math/rand"
	"strings"
	"time"
)

const alphas = "abcdefghijklmnopqrstuvwxyz"

func RandomString(length int) string {
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	var sb strings.Builder
	k := len(alphas)
	for index := 0; index < length; index++ {
		c := alphas[random.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
