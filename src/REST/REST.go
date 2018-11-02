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
	. "reststruct"
	"strconv"
	"subdmongo"
)

var serverString = "localhost:8000" //5050
func LogString(s string) {
	log.Println("Inf: " + s)
}

//Done
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
		LogString(string(res2B))
		if m.AuthMethod == "password" {
			//	var p PersonStruct.Person
			p, err := PersonStruct.FindPersonByLogin(m.Login, m.Password)
			if err != nil {
				if err.Error() == "not found" {
					log.Println("not found")
					p, err = PersonStruct.InsertPersonWithID(m.Login, m.Password, m.AccountID)
					mo := Messageout{
						Balance: p.Balance,
						Status:  "ok",
						Token:   p.Tocken,
					}
					b, err := json.Marshal(mo)
					if err != nil {
						http.Error(w, err.Error(), 400)
					} else {
						w.Write(b)
					}
					if err != nil {
						http.Error(w, err.Error(), 400)
					}
				} else {
					http.Error(w, err.Error(), 400)
				}
			} else {
				mo := Messageout{
					Balance: 0.0,
					Status:  "exist",
					Token:   "",
				}
				b, err := json.Marshal(mo)
				if err != nil {
					http.Error(w, err.Error(), 400)
				} else {
					w.Write(b)
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
		LogString(string(res2B))
		if m.AuthMethod == "token" {
			//	var p PersonStruct.Person
			p, ok := PersonStruct.FindPersonByToken(m.Token)

			if !ok {

				mo := Messageout{
					Balance: 0,
					Status:  "not found",
					Token:   "",
				}
				b, err := json.Marshal(mo)
				if err == nil {
					w.Write(b)
				} else {
					http.Error(w, err.Error(), 400)
				}
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
			//	var p PersonStruct.Person
			p, err := PersonStruct.FindPersonByLogin(m.Login, m.Password)
			if err != nil {
				if err.Error() == "not found" {
					mo := Messageout{
						Balance: 0,
						Status:  "not found",
						Token:   "",
					}
					b, err := json.Marshal(mo)
					if err == nil {
						w.Write(b)
					} else {
						http.Error(w, err.Error(), 400)
					}
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
					LogString(string(b))
					w.Write(b)
				} else {
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
		LogString(string(res2B))
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			http.Error(w, err.Error(), 400)
		}
		mo := Messageout{
			Balance: p.Balance,
			Status:  "ok",
		}
		b, err := json.Marshal(mo)
		if err == nil {
			LogString(string(b))
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
	data, err := ios.ReadFile("Историяразм.psd")
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
		LogString(string(res2B))
		//
		a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		a.AddNewTockenWithoutTeam(m.Token)

		//
		mo := Messageout{
			Status: "ok",
		}
		b, err := json.Marshal(mo)
		if err == nil {
			LogString(string(b))
			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

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

	type Messageout struct {
		Status     string          `json:"status"`
		AccountIDs []int           `json:"accountIDs"`
		Pending    []StructForREST `json:"pending"`
		Active     []StructForREST `json:"active"`
		Incoming   []StructForREST `json:"incoming"`
		Rejected   []StructForREST `json:"rejected"`
		Declined   []StructForREST `json:"declined"`
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
		LogString(string(res2B))
		//Verify
		for _, value := range m.Active {
			memp.VerifyActive(m.Token, strconv.Itoa(m.ArenaID), value.AccountID, value.BetValue)
		}
		for _, value := range m.Rejected {
			memp.VerifyReject(m.Token, strconv.Itoa(m.ArenaID), value.AccountID, value.BetValue)
		}
		for _, value := range m.Declined {
			memp.VerifyDecline(m.Token, strconv.Itoa(m.ArenaID), value.AccountID, value.BetValue)
		}
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			http.Error(w, err.Error(), 400)
		}
		//
		massP := memp.GetPending(strconv.Itoa(m.ArenaID), p.AccountID)
		massA := memp.GetActive(strconv.Itoa(m.ArenaID), p.AccountID)
		massI := memp.GetIncoming(strconv.Itoa(m.ArenaID), p.AccountID)
		massR := memp.GetRejected(strconv.Itoa(m.ArenaID), p.AccountID)
		massD := memp.GetDeclined(strconv.Itoa(m.ArenaID), p.AccountID)
		//
		a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		strs := a.GetEnemies(m.Token)
		var massIds []int
		for _, e := range strs {
			igf, _ := strconv.Atoi(e)
			massIds = append(massIds, igf)
		}
		mo := Messageout{
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
			LogString(string(b))
			w.Write(b)
		} else {
			http.Error(w, err.Error(), 400)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

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
		LogString(string(res2B))
		//Verify
		res := StructForREST{
			ArenaID:   m.ArenaID,
			AccountID: m.ToAccountID,
			ParryType: m.ParryType,
			BetValue:  m.BetValue,
		}

		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			http.Error(w, "Token not find", 400)
		}
		res2 := StructForREST{
			ArenaID:   m.ArenaID,
			AccountID: p.AccountID,
			ParryType: m.ParryType,
			BetValue:  m.BetValue,
		}
		stat := "ok"
		if memp.IsAddParry(strconv.Itoa(m.ArenaID), m.ToAccountID, p.AccountID) {
			stat = "already exist"
		} else {
			memp.AddParry(res, strconv.Itoa(m.ArenaID), "incoming", m.ToAccountID, p.AccountID)
			memp.AddParry(res2, strconv.Itoa(m.ArenaID), "pending", m.ToAccountID, p.AccountID)
		}
		/* a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		strs := a.GetEnemies(m.Token)
		var massIds []int
		for _, e := range strs {
			igf, _ := strconv.Atoi(e)
			massIds = append(massIds, igf)
		} */

		//
		mo := Messageout{
			Status: stat,
		}
		b, err := json.Marshal(mo)

		if err == nil {
			LogString(string(b))
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
		LogString(string(res2B))
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
		LogString(string(b))
		if err == nil {
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
		Status  string          `json:"status"`
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
		res2B, _ := json.Marshal(m)
		LogString(string(res2B))
		//
		a := mem.Arena.FindArena(strconv.Itoa(m.ArenaID))
		p, ok := PersonStruct.FindPersonByToken(m.Token)
		if !ok {
			http.Error(w, "invalid token", 400)
		}
		if m.Data["victory"] {
			a.TokenWinWithoutTeam(strconv.Itoa(p.AccountID))
		} else {
			a.TokenLoseWithoutTeam(strconv.Itoa(p.AccountID))
		}

		//

		massVictory, ok := a.GetVictoriesWithoutTeam()
		mo := Messageout{
			Status:  "not ended",
			Victory: nil,
			Defeat:  nil,
			Balance: 0,
		}
		if ok {
			massLose, _ := a.GetLosesWithoutTeam()
			p, _ := PersonStruct.FindPersonByToken(m.Token)
			mo = Messageout{
				Status:  "ok",
				Victory: massVictory,
				Defeat:  massLose,
				Balance: p.Balance,
			}
		}

		c, err := json.Marshal(mo)
		LogString(string(c))
		if err != nil {
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

	http.HandleFunc("/currentVersion", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, par.CurrentVersion)
	}) //tested
	http.HandleFunc("/StatsAllPersons", HandleFunctionStatAllPerson)       //tested
	http.HandleFunc("/StatsActivePersons", HandleFunctionStatActivePerson) //tested
	http.HandleFunc("/StatAllBets", HandleFunctionStatAllBets)             //tested
	http.HandleFunc("/wotmod", HandleFunctionGetMod)                       //tested
	http.HandleFunc("/login", HandleFunctionLogin)                         //tested
	http.HandleFunc("/registration", HandleFunctionRegistration)           //tested
	///
	http.HandleFunc("/balance", HandleFunctionBalance) //tested
	//
	http.HandleFunc("/arena/enter", HandleFunctionArenaEnter)
	http.HandleFunc("/arena/situation", HandleFunctionArenaSituation)
	http.HandleFunc("/parry", HandleFunctionParry)
	http.HandleFunc("/arena/quit", HandleFunctionArenaQuit)
	http.HandleFunc("/arena/result", HandleFunctionArenaResult)
	if err := http.ListenAndServe(serverString, nil); err != nil {
		log.Fatal(err)
	}

}
