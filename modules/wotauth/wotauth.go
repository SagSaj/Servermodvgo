package wotauth

import (
	"io/ioutil"
	"log"
	. "main/modules/reststruct"
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

func VerifyWotID(ID int) (bool, string) {
	r, err := http.Get("https://api.worldoftanks.ru/wot/account/info/?application_id=1416d3f7652ca060ae19ad1032f97c6a&account_id=" + strconv.Itoa(ID))
	if err != nil {
		return false, ""
	}
	var f jsoniter.Any
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, ""
	}
	log.Println("authbyte " + string(body) + r.Status)
	jsoniter.ConfigFastest.Unmarshal([]byte(body), &f)
	if f.Get("status").ToString() == "error" {
		return false, ""
	}
	return true, f.Get("data").Get(strconv.Itoa(ID)).Get("nickname").ToString()

}
func RegisterID(ID int) (bool, LoginInformation) {
	r, err := http.Get("https://api.worldoftanks.ru/wot/account/info/?application_id=1416d3f7652ca060ae19ad1032f97c6a&account_id=" + strconv.Itoa(ID))
	l := LoginInformation{}
	if err != nil {
		return false, l
	}
	var f jsoniter.Any
	var buf []byte
	_, err = r.Body.Read(buf)
	if err != nil {
		return false, l
	}

	jsoniter.ConfigFastest.Unmarshal(buf, &f)
	log.Println("auth " + f.ToString())
	if f.Get("status").ToString() == "error" {
		return false, l
	}
	l = LoginInformation{NameInWot: f.Get("data").Get(strconv.Itoa(ID)).Get("nickname").ToString()}
	return true, l
}
