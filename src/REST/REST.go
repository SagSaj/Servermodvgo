package REST

import (
	par "Params"
	"PersonStruct"
	"encoding/json"
	"fmt"
	ios "io/ioutil"
	"log"
	mem "memcash"
	memp "memcashparry"
	"net/http"
	"os"
	. "reststruct"
	"strconv"
	"subdmongo"
	"time"
)

var serverString = "8000" //5050
func LogString(s string, funct string) {
	log.Println("Inf " + funct + ":" + s)
}

type MessageError struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
} //Errors
//Done
func DropBase() {
	PersonStruct.DropBase()
}
func HandleFunctionRegistration(w http.ResponseWriter, r *http.Request) {
	//Done
	/*POST /login
	Параметры от клиента:
	{“accountID”: 20388892, //Wargaming account ID
	“login”: “client_mail@mail.com”, //логин в нашей системе (может отличаться от того, на который зарегистрирован аккаунт в WoT)
	“auth_method”: “token” // “password” или “token” (в каком поле искать пароль)
	“token”: “r47r3y789h2378d932y6r98”, // токен, замена пароля
	“password”: “” // в любом запросе будет ЛИБО токен ЛИБО пароль
	}

	Ответ от сервера:
	{“token”: “”, // обязательный параметр, в дальнейших запросах играет роль подтверждения валидности сессии
	“balance”: 10.8, // число
	“status”: “ok” // “WRONG_ACCOUNT_ID”, “INVALID_TOKEN”
	}*/
	type Message struct {
		AccountID  int    `json:"accountID"`   //Wargaming account ID
		Login      string `json:"login"`       //логин в нашей системе (может отличаться от того, на который зарегистрирован аккаунт в WoT)
		AuthMethod string `json:"auth_method"` // “password” или “token” (в каком поле искать пароль)
		Token      string `json:"token"`
		Password   string `json:"password"` // в любом запросе будет ЛИБО токен ЛИБО пароль
	}
	type Messageout struct {
		Token   string  `json:"token"`
		Balance float32 `json:"balance"`
		Status  string  `json:"status"`
	}
	var m Message
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "registration")
		if m.AuthMethod == "password" {
			//	var p PersonStruct.Person
			p, err := PersonStruct.FindPersonByLogin(m.Login, m.Password)
			if err != nil {
				if err.Error() == "not found" {

					p, err = PersonStruct.InsertPersonWithID(m.Login, m.Password, m.AccountID)
					if err != nil {
						mo := MessageError{Error: "LOGIN_EXIST"}
						b, err := json.Marshal(mo)
						if err != nil {
							http.Error(w, err.Error(), 401)
						} else {
							w.Write(b)
						}
						return
					}

					mo := Messageout{
						Balance: p.Balance,
						Status:  "ok",
						Token:   p.Tocken,
					}

					b, err := json.Marshal(mo)
					if err != nil {
						http.Error(w, err.Error(), 401)
					} else {
						w.Write(b)
					}
				} else {
					http.Error(w, err.Error(), 400)
				}
			} else {

				if err != nil {
					http.Error(w, err.Error(), 400)
				} else {
					mo := MessageError{Error: "LOGIN_EXISTS"}
					b, err := json.Marshal(mo)
					if err != nil {
						http.Error(w, err.Error(), 401)
					} else {
						w.Write(b)
					}
					return
				}
				if err != nil {
					http.Error(w, err.Error(), 400)
				}
			}

		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

//Done
func HandleFunctionLogin(w http.ResponseWriter, r *http.Request) {
	//Done
	/*POST /login
	Параметры от клиента:
	{“accountID”: 20388892, //Wargaming account ID
	“login”: “client_mail@mail.com”, //логин в нашей системе (может отличаться от того, на который зарегистрирован аккаунт в WoT)
	“auth_method”: “token” // “password” или “token” (в каком поле искать пароль)
	“token”: “r47r3y789h2378d932y6r98”, // токен, замена пароля
	“password”: “” // в любом запросе будет ЛИБО токен ЛИБО пароль
	}

	Ответ от сервера:
	{“token”: “”, // обязательный параметр, в дальнейших запросах играет роль подтверждения валидности сессии
	“balance”: 10.8, // число
	“status”: “ok” // “WRONG_ACCOUNT_ID”, “INVALID_TOKEN”
	}*/
	type Message struct {
		AccountID  int    `json:"accountID"`   //Wargaming account ID
		Login      string `json:"login"`       //логин в нашей системе (может отличаться от того, на который зарегистрирован аккаунт в WoT)
		AuthMethod string `json:"auth_method"` // “password” или “token” (в каком поле искать пароль)
		Token      string `json:"token"`
		Password   string `json:"password"` // в любом запросе будет ЛИБО токен ЛИБО пароль
	}
	type Messageout struct {
		Token   string  `json:"token"`
		Balance float32 `json:"balance"`
		Status  string  `json:"status"`
	}
	var m Message
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "Login")
		if m.AuthMethod == "token" {
			//	var p PersonStruct.Person
			p, ok := PersonStruct.FindPersonByToken(m.Token)

			if !ok {
				mo := MessageError{Error: "INVALID_TOKEN"}
				b, err := json.Marshal(mo)
				if err != nil {
					http.Error(w, err.Error(), 401)
				} else {
					w.Write(b)
				}
				return
			} else {
				if p.AccountID == m.AccountID {
					mo := Messageout{
						Balance: p.Balance,
						Status:  "ok",
						Token:   p.Tocken,
					}
					b, err := json.Marshal(mo)
					if err == nil {
						w.Write(b)
					} else {
						http.Error(w, err.Error(), 400)
					}
				} else {
					http.Error(w, "Not found this AccountID", 400)
				}
			}
		}
		if m.AuthMethod == "password" {
			//	var p PersonStruct.Person errors by found
			p, err := PersonStruct.FindPersonByLogin(m.Login, m.Password)
			if err != nil {
				log.Println(err.Error())
				if err.Error() == "not found" {
					log.Println("sad")
					mo := MessageError{Error: "WRONG_ACCOUNT_ID"}
					b, err := json.Marshal(mo)
					if err != nil {
						http.Error(w, err.Error(), 401)
					} else {
						w.Write(b)
					}
					return
				} else {

					http.Error(w, err.Error(), 400)
				}
			} else {
				//Add AccountID
				PersonStruct.AddAccountIDLogIN(p.Tocken, m.AccountID)
				mo := Messageout{
					Balance: p.Balance,
					Status:  "ok",
					Token:   p.Tocken,
				}
				b, err := json.Marshal(mo)
				if err == nil {
					LogString(string(b), "Login")
					w.Write(b)
				} else {
					LogString(string(b), "Login")
					http.Error(w, err.Error(), 400)
				}
			}
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

//Done
func HandleFunctionBalance(w http.ResponseWriter, r *http.Request) {
	//Done
	/*POST /balance

	Параметры от клиента:
	{“token”: “”}

	Ответ от сервера:
	{“balance”: 10.8, “status”: “ok” // INVALID_TOKEN}

	*/
	type Message struct {
		Token string `json:"token"`
	}
	type Messageout struct {
		Balance float32 `json:"balance"`
		Status  string  `json:"status"`
	}
	var m Message
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "Balance")
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {

			mo := MessageError{Error: "INVALID_TOKEN"}
			b, err := json.Marshal(mo)
			if err != nil {
				http.Error(w, err.Error(), 401)
			} else {
				w.Write(b)
			}
			return
		}
		mo := Messageout{
			Balance: p.Balance,
			Status:  "ok",
		}
		b, err := json.Marshal(mo)
		if err == nil {
			LogString(string(b), "Balance")
			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

//Done
func HandleFunctionGetMod(w http.ResponseWriter, r *http.Request) {
	/*GET /wotmod
	Параметры от клиента: нет
	Ответ сервера: файл модификации в бинарном формате.
	*/
	data, err := ios.ReadFile("Dueler.zip")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

//Done
func HandleFunctionArenaEnter(w http.ResponseWriter, r *http.Request) {
	// Done
	/*POST /arena/enter

	Параметры от клиента:
	{“token”: “”,
	“arenaID”: 4372947891}

	Ответ от сервера:
	{“status”: “ok” // либо “errorID” (e.g. “INVALID_TOKEN”)
	}*/
	type Message struct {
		ArenaID int    `json:"arenaID"`
		Token   string `json:"token"`
	}

	type Messageout struct {
		Status string `json:"status"`
	}
	var m Message
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "Enter")

		//
		a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			mo := MessageError{Error: "INVALID_TOKEN"}
			b, err := json.Marshal(mo)
			if err != nil {
				http.Error(w, err.Error(), 401)
			} else {
				w.Write(b)
			}
			return
		}
		(a).AddNewTockenWithoutTeam(p.AccountID)
		log.Println("Enter")
		a = mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		log.Println(a.GetEnemiesWithoutTeam())
		//
		mo := Messageout{
			Status: "ok",
		}
		b, err := json.Marshal(mo)
		if err == nil {
			LogString(string(b), "Enter")
			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

type MessageoutSit struct {
	Status     string          `json:"status"`
	AccountIDs []int           `json:"accountIDs"`
	Pending    []StructForREST `json:"pending"`
	Active     []StructForREST `json:"active"`
	Incoming   []StructForREST `json:"incoming"`
	Rejected   []StructForREST `json:"rejected"`
	Declined   []StructForREST `json:"declined"`
}

var mapSit map[string]MessageoutSit

//Done
func HandleFunctionArenaSituation(w http.ResponseWriter, r *http.Request) {
	//NOne
	/*POST /arena/situation

	Параметры от клиента:
	{“arenaID”: 4372947891,
	“token”: “”,
	“pending”: [],  // исходящие от игрока
	“active”: [{“arenaID”: 4372947891, “accountID”: 327189, “parryType”: “victory”, “betValue”: 1}],  // активные (принятые)
	“incoming”: [],  // входящие к игроку
	“rejected”: [],  // отклоненные
	“declined”: []  // отозванные
	}

	Ответ от сервера:
	{“status”: “ok”, // “INVALID_TOKEN”
	“accountIDs”: [21798, 371389, 327189] //массив accountID, попавших в этот же бой.
	“incoming”: [{“arenaID”: 4372947891, “accountID”: 21798, “parryType”: “teamvictory”, “betValue”: 2}, {“arenaID”: 4372947891, “accountID”: 371389, “parryType”: “teamvictory”, “betValue”: 4}],
	“active”: [{“arenaID”: 4372947891, “accountID”: 327189, “parryType”: “teamvictory”, “betValue”: 1}],
	“pending”: [],
	“rejected”: [],
	“declined”: []
	}*/
	type Message struct {
		ArenaID  int             `json:"arenaID"`
		Token    string          `json:"token"`
		Pending  []StructForREST `json:"pending"`
		Active   []StructForREST `json:"active"`
		Incoming []StructForREST `json:"incoming"`
		Rejected []StructForREST `json:"rejected"`
		Declined []StructForREST `json:"declined"`
	}

	if r.Method == "POST" {
		var m Message
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//res2B, _ := json.Marshal(m)
		//LogString(string(res2B), "Situation")
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			mo := MessageError{Error: "INVALID_TOKEN"}
			b, err := json.Marshal(mo)
			if err != nil {
				http.Error(w, err.Error(), 401)
			} else {
				w.Write(b)
			}
			return
		}
		//Verify
		//pers, ok := PersonStruct.FindPersonByToken(m.Token)
		// for _, value := range m.Active {

		// 	if memp.VerifyActive(strconv.Itoa(m.ArenaID), pers.AccountID, value.BetValue) {

		// 		if ok {
		// 			a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		// 			a.AddNewParry(value.FromAccountID, value.ToAccountID, value.BetValue)
		// 		}
		// 	}
		// }

		//
		massP := memp.GetPending(strconv.Itoa(m.ArenaID), p.AccountID)
		massA := memp.GetActive(strconv.Itoa(m.ArenaID), p.AccountID)
		massI := memp.GetIncoming(strconv.Itoa(m.ArenaID), p.AccountID)
		massR := memp.GetRejected(strconv.Itoa(m.ArenaID), p.AccountID)
		massD := memp.GetDeclined(strconv.Itoa(m.ArenaID), p.AccountID)
		//
		if len(massP) > 0 {
			if time.Now().Sub(massP[0].CreatedAt).Minutes() > float64(2) {
				massP = nil

			}

		}
		if len(massI) > 0 {
			if time.Now().Sub(massI[0].CreatedAt).Minutes() > float64(2) {
				massI = nil
			}
		}
		a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		if a == nil {
			http.Error(w, "Arena didn't find", 402)
			return
		}
		strs := a.GetEnemiesWithoutTeam()
		massIds := []int{}
		for _, e := range strs {
			massIds = append(massIds, e)
		}
		mo := MessageoutSit{
			Status:     "ok",
			AccountIDs: massIds,
			Pending:    massP,
			Active:     massA,
			Incoming:   massI,
			Rejected:   massR,
			Declined:   massD,
		}

		b, err := json.Marshal(mo)
		if err == nil {
			temp, ok := mapSit[m.Token]
			if ok {
				if len(temp.Pending) != len(mo.Pending) || len(temp.Active) != len(mo.Active) || len(temp.Incoming) != len(mo.Incoming) {
					LogString(string(b), "Situation")
				}
			}
			mapSit[m.Token] = mo

			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

var INI_ID int

//Done
func HandleFunctionParry(w http.ResponseWriter, r *http.Request) {
	/*POST /parry

	Параметры от клиента:
	{
	 “arenaID”: 4372947891,
	 “token”: “”,
	“toAccountID”: 3424234,
	 “parryType”: “teamvictory”,
	 “betValue”: 2
	}

	Ответ от сервера:
	{
	“status”: “ok” // либо “errorID” (e.g. “parryAlreadyRegistered”, “INVALID_TOKEN”)
	}

	предложение пари от оппонента принимается по сокету на порту 5050:
	{“status”: “ok”, // “INVALID_TOKEN”
	 “accountIDs”: [21798, 371389, 327189] //массив accountID, попавших в этот же бой.
	 “incoming”: [{“arenaID”: 4372947891, “accountID”: 21798, “parryType”: “teamvictory”, “betValue”: 2}, {“arenaID”: 4372947891, “accountID”: 371389, “parryType”: “teamvictory”, “betValue”: 4}],
	 “active”: [{“arenaID”: 4372947891, “accountID”: 327189, “parryType”: “teamvictory”, “betValue”: 1}],
	“pending”: [],
	“rejected”: [],
	“declined”: []
	}*/
	type Message struct {
		ArenaID     int     `json:"arenaID"`
		Token       string  `json:"token"`
		ToAccountID int     `json:"toAccountID"`
		ParryType   string  `json:"parryType"`
		BetValue    float32 `json:"betValue"`
	}
	/* type Messageout struct {
		Status     string          `json:"status"`
		AccountIDs []int           `json:"accountIDs"`
		Pending    []StructForREST `json:"pending"`
		Active     []StructForREST `json:"active"`
		Incoming   []StructForREST `json:"incoming"`
		Rejected   []StructForREST `json:"rejected"`
		Declined   []StructForREST `json:"declined"`
	} */
	type Messageout struct {
		Status string `json:"status"`
	}
	if r.Method == "POST" {
		var m Message
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "Parry")
		//Verify
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			mo := MessageError{Error: "INVALID_TOKEN"}
			b, err := json.Marshal(mo)
			if err != nil {
				http.Error(w, err.Error(), 401)
			} else {
				w.Write(b)
			}
			return
		}
		stat := "ok"
		memp.Clear(p.AccountID, strconv.Itoa(m.ArenaID))
		if m.ToAccountID == 0 {
			//////////
			//////////
			log.Println(r.RequestURI)
			a, are := mem.Arena.FindArenaIDByAccountID(p.AccountID)
			if r.RequestURI == "/parry/activate/" {
				massI := memp.GetIncoming(are, p.AccountID)
				if len(massI) == 0 {
					log.Println("Can't find parry")
					log.Println(are)
					log.Println(p.AccountID)
					http.Error(w, "Can't find parry", 400)
					return
				}
				value := massI[0]

				if memp.VerifyActive(are, value.ToAccountID, value.BetValue) {
					a.AddNewParry(value.FromAccountID, value.ToAccountID, value.BetValue)
				}
			}
			if r.RequestURI == "/parry/reject/" {
				massI := memp.GetIncoming(are, p.AccountID)
				if len(massI) == 0 {
					log.Println("Can't find parry")
					http.Error(w, "Can't find parry", 400)
					return
				}
				value := massI[0]
				temp := strconv.Itoa(value.ToAccountID)
				log.Println(m.Token + " " + are + " " + temp)
				//memp.VerifyDecline(m.Token, are, value.ToAccountID, value.BetValue)
				memp.VerifyReject(m.Token, are, value.ToAccountID, value.BetValue)
			}
			if r.RequestURI == "/parry/decline/" {
				log.Println("In decline")
				massI := memp.GetPending(are, p.AccountID)
				if len(massI) == 0 {
					log.Println("Can't find parry")
					http.Error(w, "Can't find parry", 400)
					return
				}
				value := massI[0]
				log.Println("Verifying")
				temp := strconv.Itoa(value.ToAccountID)
				log.Println(m.Token + " " + are + " " + temp)
				memp.VerifyDecline(m.Token, are, value.ToAccountID, value.BetValue)
				//memp.VerifyReject(m.Token, are, value.ToAccountID, value.BetValue)
			}

		} else {
			res := StructForREST{
				ID:            INI_ID,
				ArenaID:       m.ArenaID,
				FromAccountID: p.AccountID,
				ToAccountID:   m.ToAccountID,
				ParryType:     m.ParryType,
				BetValue:      m.BetValue,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Status:        "incoming",
			}
			res2 := StructForREST{
				ID:            INI_ID,
				ArenaID:       m.ArenaID,
				FromAccountID: p.AccountID,
				ToAccountID:   m.ToAccountID,
				ParryType:     m.ParryType,
				BetValue:      m.BetValue,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Status:        "pending",
			}
			INI_ID++

			if memp.IsAddParry(strconv.Itoa(m.ArenaID), m.ToAccountID, p.AccountID) {
				stat = "already exist"
			} else {
				memp.AddParry(res, strconv.Itoa(m.ArenaID), "incoming", m.ToAccountID, p.AccountID)
				memp.AddParry(res2, strconv.Itoa(m.ArenaID), "pending", p.AccountID, m.ToAccountID)
			}
			/* a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
			strs := a.GetEnemies(m.Token)
			var massIds []int
			for _, e := range strs {
				igf, _ := strconv.Atoi(e)
				massIds = append(massIds, igf)
			} */

			//
		}
		mo := Messageout{
			Status: stat,
		}
		b, err := json.Marshal(mo)

		if err == nil {
			LogString(string(b), "Parry")
			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//Done
func HandleFunctionArenaQuit(w http.ResponseWriter, r *http.Request) {
	/*POST /arena/quit
	// при выходе из боя все пари, входящие для этого игрока, считаются отклоненными им (активные остаются активными, исходящие считаются отозванными)
	{
	“token”: “”,
	“arenaID”: 4372947891
	}

	ответ:
	{“status”: “ok” // “INVALID_TOKEN”}*/

	type Message struct {
		ArenaID int
		Token   string
	}
	type Messageout struct {
		Status string
	}
	if r.Method == "POST" {
		var m Message
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "Quit")
		//
		a, ok := mem.Arena.FindArenaEnd(strconv.Itoa(m.ArenaID))
		stans := "ok"
		if !ok {
			stans = "INVALID_TOKEN"
		} else {
			a.TokenLoseWithoutTeam(m.Token)
		}
		//
		mo := Messageout{
			Status: stans,
		}
		b, err := json.Marshal(mo)

		if err == nil {
			LogString(string(b), "Quit")
			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//Done
func HandleFunctionArenaResult(w http.ResponseWriter, r *http.Request) {
	/*

		POST /arena/result

		// отправляется при получении игрой результата боя, если во время этого боя было предложено кем-либо хотя бы одно пари

		Параметры от клиента:
		{
		“token”: “”,
		“arenaID”: 4372947891,
		“data”: {“victory”: true}
		}

		Ответ от сервера:
		{“status”: “ok”, // “INVALID_TOKEN”
		“victory”: [{“arenaID”: 4372947891, “accountID”: 327189, “parryType”: “teamvictory”, “betValue”: 1}],
		“defeat”: [],
		“balance”: 20
		}*/

	type Message struct {
		ArenaID int             `json:"arenaID"`
		Token   string          `json:"token"`
		Data    map[string]bool `json:"data"`
	}
	type Messageout struct {
		Victory []StructForREST `json:"victory"`
		Defeat  []StructForREST `json:"defeat"`
		Balance float32         `json:"balance"`
	}
	var m Message
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			mo := MessageError{Error: "INVALID_TOKEN"}
			b, err := json.Marshal(mo)
			if err != nil {
				http.Error(w, err.Error(), 401)
			} else {
				w.Write(b)
			}
			return
		}
		res2B, _ := json.Marshal(m)
		LogString(string(res2B), "result")
		//
		a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		tempArray := memp.GetActive(a.IDArena, p.AccountID)
		if len(tempArray) > 0 {
			if m.Data["victory"] {
				a.TokenWinWithoutTeam(m.Token)
			} else {
				a.TokenLoseWithoutTeam(m.Token)
			}
		} else {
			mo := MessageError{Error: "ARENA_NOT_FOUND"}
			b, err := json.Marshal(mo)
			if err != nil {
				http.Error(w, err.Error(), 401)
			} else {
				w.Write(b)
			}
			return
		}
		massVictory, ok := a.GetVictoriesWithoutTeam(m.Token)
		massLose, _ := a.GetLosesWithoutTeam(m.Token)
		mo := Messageout{
			Victory: nil,
			Defeat:  nil,
			Balance: 0,
		}
		if ok {

			mo = Messageout{
				Victory: massVictory,
				Defeat:  massLose,
				Balance: p.Balance,
			}
		}

		//

		type Messageout2 struct {
			Arena  Messageout `json:"arena"`
			Status string     `json:"status"`
		}
		mtemp := Messageout2{Arena: mo, Status: "ok"}
		c, err := json.Marshal(mtemp)
		//c, err := json.Marshal(mo)
		LogString(string(c), "result")
		if err == nil {
			w.Write(c)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//Done
func HandleFunctionStatAllPerson(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		fmt.Fprintf(w, strconv.Itoa(subdmongo.GetAllPersons()))

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//Done
func HandleFunctionStatActivePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fmt.Fprintf(w, strconv.FormatUint(subdmongo.GetStats(), 10))

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//Done
func HandleFunctionStatAllBets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fmt.Fprintf(w, strconv.FormatUint(subdmongo.GetStats(), 10))

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//Done
func GoServerListen() {
	/*GET /currentVersion
	Параметры от клиента: нет
	Ответ сервера: строка вида v.1.0.0 */
	mapSit = make(map[string]MessageoutSit, 2)
	INI_ID = 0
	port := os.Getenv("PORT")
	if port == "" {
		port = serverString
	}
	http.HandleFunc("/currentVersion/", func(w http.ResponseWriter, r *http.Request) {

		log.Println("GetVersion")
		fmt.Fprintf(w, par.CurrentVersion)
	}) //tested
	http.HandleFunc("/StatsAllPersons/", HandleFunctionStatAllPerson)       //tested
	http.HandleFunc("/StatsActivePersons/", HandleFunctionStatActivePerson) //tested
	http.HandleFunc("/StatAllBets/", HandleFunctionStatAllBets)             //tested
	http.HandleFunc("/wotmod/", HandleFunctionGetMod)
	http.HandleFunc("/account/login/", HandleFunctionLogin)
	http.HandleFunc("/account/register/", HandleFunctionRegistration)
	////account/register/
	http.HandleFunc("/balance/", HandleFunctionBalance)
	//
	http.HandleFunc("/arena/enter/", HandleFunctionArenaEnter)
	http.HandleFunc("/arena/situation/", HandleFunctionArenaSituation)
	http.HandleFunc("/parry/", HandleFunctionParry)
	http.HandleFunc("/arena/quit/", HandleFunctionArenaQuit)
	http.HandleFunc("/arena/result/", HandleFunctionArenaResult)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Println("Strange", r.RequestURI)
		fmt.Fprintf(w, r.RequestURI)
	})

	log.Println("Started")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
