package wotauth

import (
	"io/ioutil"
	"log"
	"strconv"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func _TestAuth(t *testing.T) {
	_, r := VerifyWotID(11100)
	log.Println("Nick " + r)
}
func TestJsonfil(t *testing.T) {
	i := 11100
	body, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Println(err.Error())
	}
	var f jsoniter.Any
	err = jsoniter.ConfigFastest.Unmarshal([]byte(body), &f)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(f.Get("data").Get(strconv.Itoa(i)).Get("nickname").ToString())
}
