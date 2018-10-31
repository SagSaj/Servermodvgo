package subdmongo

import (

	// "fmt"

	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SessionGame Struct of Session
type SessionGame struct {
	Login     string
	Session   string
	TimeStart int
}

//var session *mgo.Session
var err error
var dBName = "DB"

func init() {
	/* session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true) */

}

var mgoSession *mgo.Session

// Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("127.0.0.1")
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}

//InsertIntoDatabase Session
func InsertIntoDatabase(p SessionGame) {
	//initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	c := session.DB(dBName).C("Sessions")

	err = c.Insert(&p)
	if err != nil {
		log.Println(err)
	}
}

//FindBySession ss
func FindBySession(Ses string) []SessionGame {
	//initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	result := []SessionGame{}
	c := session.DB(dBName).C("Sessions")
	err = c.Find(bson.M{"Session": Ses}).All(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

//DeletebyTimeOut dt
func DeletebyTimeOut(ti float64) {
	//	initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	c := session.DB(dBName).C("Sessions")
	err = c.Remove(bson.M{"TimeStart": bson.M{"$lte": ti}})
	if err != nil {
		log.Println(err)
	}
}

//LoginInformation ii
type LoginInformation struct {
	Login     string
	Password  string
	Balance   float32
	WinCount  int
	LoseCount int
	IDAccount int
}

//RegistrNewPerson rnp
func RegistrNewPerson(login, password string) (LoginInformation, error) {
	//	initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	def := LoginInformation{}
	_, b, err := findPerson(login)
	if err != nil {
		return def, err
	}
	if b {
		return def, errors.New("Exist")
	}
	l := LoginInformation{Login: login, Password: password, Balance: 100, WinCount: 0, LoseCount: 0}
	c := session.DB(dBName).C("Persons")
	err = c.Insert(&l)
	if err != nil {
		log.Println("Registr" + err.Error())
		return l, err
	}
	return l, nil
}
func RegistrNewPersonWithID(login, password string, ID int) (LoginInformation, error) {
	//	initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	def := LoginInformation{}
	_, b, err := findPerson(login)
	if err != nil {
		return def, err
	}
	if b {
		return def, errors.New("Exist")
	}
	l := LoginInformation{Login: login, Password: password, Balance: 100, WinCount: 0, LoseCount: 0, IDAccount: ID}
	c := session.DB(dBName).C("Persons")
	err = c.Insert(&l)
	if err != nil {
		log.Println("Registr" + err.Error())
		return l, err
	}
	return l, nil
}

//FindPerson fp
func FindPerson(login, password string) (LoginInformation, error) {
	//initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	result0 := LoginInformation{}
	c := session.DB(dBName).C("Persons")
	err = c.Find(bson.M{"login": login, "password": password}).One(&result0)
	if err != nil {
		log.Println("FindPerson" + err.Error())
		return result0, err
	}
	return result0, nil
}

//find pers
func findPerson(login string) (LoginInformation, bool, error) {
	//initiateSession()
	//defer session.Close()
	result0 := LoginInformation{}
	session := GetMongoSession()
	defer session.Close()
	c := session.DB(dBName).C("Persons")
	err = c.Find(bson.M{"login": login}).One(result0)

	if err != nil {
		if err.Error() == "not found" {
			return result0, false, nil
		}
		log.Println(err)
		return result0, false, err
	}
	return result0, true, nil
}
func GetBalance(Login string) (bool, float32) {
	session := GetMongoSession()
	defer session.Close()
	result0 := LoginInformation{}
	c := session.DB(dBName).C("Persons")
	err = c.Find(bson.M{"login": Login}).One(&result0)
	if err != nil {
		log.Println(err)
		return false, 0
	}
	return true, result0.Balance
}
func GetAllPersons() int {
	session := GetMongoSession()
	defer session.Close()

	c := session.DB(dBName).C("Persons")
	n, err := c.Find(nil).Count()
	if err != nil {
		log.Println(err)
		return 0
	}
	return n
}
func SetBalanceAndWinCount(login string, balanceChange float32, winCount int, loseCount int) error {
	result, b, err := findPerson(login)
	session := GetMongoSession()
	defer session.Close()
	if err != nil {
		return err
	}
	if !b {
		return errors.New("Not exist")
	}
	result.LoseCount = result.LoseCount + loseCount
	result.WinCount = result.WinCount + winCount
	result.Balance = result.Balance + balanceChange
	c := session.DB(dBName).C("Persons")
	err = c.Insert(&result)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type Stats struct {
	prices uint64
}

func RegistrStats(pric uint64) {
	//	initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	p := Stats{prices: pric}
	c := session.DB(dBName).C("Stats")
	err = c.Insert(&p)
	if err != nil {
		log.Println("Stats" + err.Error())
	}
}
func IncrementStats() {
	i := GetStats()
	i += 1
	RegistrStats(i)
}
func GetStats() uint64 {
	//	initiateSession()
	//defer session.Close()
	session := GetMongoSession()
	defer session.Close()
	p := Stats{}
	c := session.DB(dBName).C("Stats")
	err = c.Find(nil).One(&p)
	if err != nil {
		log.Println("Stats" + err.Error())
		return 0
	}
	return p.prices
}
