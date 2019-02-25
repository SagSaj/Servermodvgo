package subdmongo

import (
	// "fmt"
	"log"
	"testing"
)

func TestDBOtest(t *testing.T) {
	InsertIntoDatabase(SessionGame{"Login", "Session", 0})
	sesGame := FindBySession("Session")
	for _, s := range sesGame {
		log.Println(s.Login)
	}
}
