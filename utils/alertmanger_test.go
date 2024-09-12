package utils

import (
	"log"
	"testing"
)

func TestSelectAlertsByKey(t *testing.T) {
	key, err := SelectAlertsByKey("L1", "status", "active")
	if err != nil {
		log.Println(err)
	}
	println(key)
}
