package tockenplc

import (
	"bytes"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RegistrationInformation struct {
	Name              string `json:"name"`
	Owner             string `json:"owner"`
	Active            string `json:"active"`
	Registrar_account string `json:"registrar_account"`
	Referrer_account  string `json:"referrer_account"`
	Referrer_percent  string `json:"referrer_percent"`
	Broadcast         string `json:"broadcast"`
}

//struct rpc
const server = ":8000"

func RegistrNewAccount(r RegistrationInformation) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	http.Post(server, "rpc/json;", b)
	//io.Copy(os.Stdout, res.Body) check
	return nil
}

//transfer(string from, string to, string amount, string asset_symbol, string memo, bool broadcast = false)
type TransferInformation struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	Asset_symbol string `json:"asset_symbol"`
	Memo         string `json:"memo"`
	Broadcast    string `json:"broadcast"`
}

func Transfer(r TransferInformation) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	http.Post(server, "rpc/json;", b)
	//io.Copy(os.Stdout, res.Body) check
	return nil
}
