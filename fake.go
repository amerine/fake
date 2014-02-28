// Package fake provides functions for generating fake data.
package fake

import (
	"math/rand"
	"time"
)

func randomString(list []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	value := list[rand.Intn(len(list))]
	return value
}
