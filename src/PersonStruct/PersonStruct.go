package PersonStruct

import (
	"fmt"
	"log"
	mongo "subdmongo"
	"time"
)

type PersonsBalance struct {
	Balance float32
}

var ServicePerson map[string]PersonService

type PersonService struct {
	PersonInf    Person
	TimeActivity time.Time
}

func init() {
	ServicePerson = make(map[string]PersonService, 10000)
	go DeleteLongTocken()
}
func DeleteLongTocken() {
	for true {
		time.Sleep(60 * time.Second)
		for key, e := range ServicePerson {
			if time.Now().Sub(e.TimeActivity).Minutes() > float64(60) {
				delete(ServicePerson, key)
			}
		}
	}
}

type Person struct {
	Login     string
	Password  string
	Tocken    string
	Balance   float32
	WinCount  int
	LoseCount int
	AccountID int
}

func generateTocken(login string) string {

	for key, e := range ServicePerson {
		if e.PersonInf.Login == login {
			e.TimeActivity = time.Now()
			return key
		}
	}
	return login + time.Now().String()
}
func InsertPerson(login, password string) (Person, error) {
	//	p := Person{Login: login, Password: password}
	//	Login := mongo.LoginInformation{Login: login, Password: password}
	l, err := mongo.RegistrNewPerson(login, password)
	var p Person
	if err != nil {
		return p, err
	}
	p.Login = login
	p.Password = password
	p.Balance = l.Balance
	p.WinCount = l.WinCount
	p.LoseCount = l.LoseCount
	p.AccountID = l.IDAccount
	p.Tocken = generateTocken(login)
	ServicePerson[p.Tocken] = PersonService{PersonInf: p, TimeActivity: time.Now()}
	return p, nil
}
func InsertPersonWithID(login, password string, ID int) (Person, error) {
	//	p := Person{Login: login, Password: password}
	//	Login := mongo.LoginInformation{Login: login, Password: password}
	l, err := mongo.RegistrNewPersonWithID(login, password, ID)
	var p Person
	if err != nil {
		return p, err
	}
	p.Login = login
	p.Password = password
	p.Balance = l.Balance
	p.WinCount = l.WinCount
	p.AccountID = l.IDAccount
	p.LoseCount = l.LoseCount
	p.Tocken = generateTocken(login)
	ServicePerson[p.Tocken] = PersonService{PersonInf: p, TimeActivity: time.Now()}
	return p, nil
}
func FindPersonByLogin(login, password string) (Person, error) {
	var p Person

	l, err := mongo.FindPerson(login, password)
	if err != nil {
		log.Print(err.Error())
		return p, err
	}
	log.Println("Finded search")
	p.Login = login
	p.Password = password
	p.Balance = l.Balance
	log.Println(fmt.Sprintf("%f - ", p.Balance))
	p.WinCount = l.WinCount
	p.AccountID = l.IDAccount
	p.LoseCount = l.LoseCount
	p.Tocken = generateTocken(login)
	ServicePerson[p.Tocken] = PersonService{PersonInf: p, TimeActivity: time.Now()}
	return p, nil
}
func FindPersonByToken(Toc string) (Person, bool) {
	p, ok := ServicePerson[Toc]
	if !ok {
		return Person{}, false
	}
	ServicePerson[Toc] = PersonService{PersonInf: p.PersonInf, TimeActivity: time.Now()}
	return p.PersonInf, true
}
func WinMatch(Token string, bet float32) {
	p, ok := FindPersonByToken(Token)
	if ok {
		if bet == 0 {
			bet = 1
		}
		p.Balance = p.Balance + bet
		p.WinCount += 1
		mongo.SetBalanceAndWinCount(p.Login, bet, 1, 0)
	} else {
		log.Println("Tocken " + Token + " didn't find in MatchResult")
	}
}
func GetAllActivePersons() int {
	return len(ServicePerson)
}
func LoseMatch(Token string, bet float32) {
	p, ok := FindPersonByToken(Token)
	if ok {

		p.Balance = p.Balance - bet
		p.LoseCount += 1
		mongo.SetBalanceAndWinCount(p.Login, -bet, 0, 1)
	} else {
		log.Println("Tocken " + Token + " didn't find in MatchResult")
	}

}
func AddAccountIDLogIN(Token string, AccountID int) {
	p, ok := ServicePerson[Token]
	if !ok {
		log.Println("Bad  tocken " + Token + " in registration")
	}
	p.PersonInf.AccountID = AccountID
	ServicePerson[Token] = PersonService{PersonInf: p.PersonInf, TimeActivity: time.Now()}
}
