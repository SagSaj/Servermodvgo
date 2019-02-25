package wotauth

import "testing"
import "log"

func TestAuth(t *testing.T) {
	_, r := VerifyWotID(11100)
	log.Println("Nick " + r)
}
