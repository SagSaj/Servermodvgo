package wotauth

import "net/http"
import jsoniter "github.com/json-iterator/go"
import "strconv"

func VerifyWotID(ID int) (bool, string) {
	r, err := http.Get("https://api.worldoftanks.ru/wot/account/info/?application_id=1416d3f7652ca060ae19ad1032f97c6a&account_id=" + strconv.Itoa(ID))
	if err != nil {
		return false, ""
	}
	var f jsoniter.Any
	var buf []byte
	_, err = r.Body.Read(buf)
	if err != nil {
		return false, ""
	}
	jsoniter.ConfigFastest.Unmarshal(buf, &f)
	if f.Get("status").ToString() == "error" {
		return false, ""
	}
	return true, f.Get("nickname").ToString()

}
