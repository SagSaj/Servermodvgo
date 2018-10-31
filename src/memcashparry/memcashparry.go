package memcashparry

import . "reststruct"
import mem "memcash"
import "PersonStruct"

type Parry struct {
	Parres []StructForREST
	Types  []string
	From   []int
	To     []int
}

func (p *Parry) AddADV(r StructForREST, T string, FromID, ToID int) {
	p.Parres = append(p.Parres, r)
	p.Types = append(p.Types, T)
	p.From = append(p.From, FromID)
	p.To = append(p.To, ToID)
}
func (p *Parry) DeleteADV(r StructForREST) {
	for i := 0; i < len(p.Parres); i = +1 {
		if p.Parres[i] == r {
			p.Parres = append(p.Parres[:i], p.Parres[i+1:]...)
			p.Types = append(p.Types[:i], p.Types[i+1:]...)
			p.From = append(p.To[:i], p.To[i+1:]...)
			p.To = append(p.To[:i], p.To[i+1:]...)
			return
		}
	}
}
func (p *Parry) DeleteADVbyIndex(i int) {

	p.Parres = append(p.Parres[:i], p.Parres[i+1:]...)
	p.Types = append(p.Types[:i], p.Types[i+1:]...)
	p.From = append(p.To[:i], p.To[i+1:]...)
	p.To = append(p.To[:i], p.To[i+1:]...)

}
func (p *Parry) ReplaceADV(r StructForREST, typeP string) {
	for i := 0; i < len(p.Parres); i = +1 {
		if p.Parres[i] == r {
			p.Types[i] = typeP
			return
		}
	}
}

var ParryMems map[string]Parry

func init() {
	ParryMems = make(map[string]Parry, 10000)
	go DeleteLongTocken()
}
func AddParry(p StructForREST, arenaID string, parryType string, ToID int, FromID int) {
	r, ok := ParryMems[arenaID]
	if !ok {
		pz := Parry{}
		pz.AddADV(p, parryType, FromID, ToID)
		ParryMems[arenaID] = pz
	} else {
		r.AddADV(p, parryType, FromID, ToID)
	}
}

//“incoming”: [{“arenaID”: 4372947891, “accountID”: 21798, “parryType”: “teamvictory”, “betValue”: 2}, {“arenaID”: 4372947891, “accountID”: 371389, “parryType”: “teamvictory”, “betValue”: 4}],
//	 “active”: [{“arenaID”: 4372947891, “accountID”: 327189, “parryType”: “teamvictory”, “betValue”: 1}],
//	“pending”: [],
//	“rejected”: [],
//	“declined”:
func GetIncoming(ID string, IDFrom int) []StructForREST {
	a := []StructForREST{}
	p, ok := ParryMems[ID]
	if ok {
		for i := 0; i < len(p.Parres); i = +1 {
			if p.Types[i] == "incoming" {
				a = append(a, p.Parres[i])
			}
		}
	}
	return a
}
func GetActive(ID string, IDFrom int) []StructForREST {
	a := []StructForREST{}
	p, ok := ParryMems[ID]
	if ok {
		for i := 0; i < len(p.Parres); i = +1 {
			if p.Types[i] == "active" {
				a = append(a, p.Parres[i])
			}
		}
	}
	return a
}
func GetPending(ID string, IDFrom int) []StructForREST {
	a := []StructForREST{}
	p, ok := ParryMems[ID]
	if ok {
		for i := 0; i < len(p.Parres); i = +1 {
			if p.Types[i] == "pending" {
				a = append(a, p.Parres[i])
			}
		}
	}
	return a
}
func GetRejected(ID string, IDFrom int) []StructForREST {
	a := []StructForREST{}
	p, ok := ParryMems[ID]
	if ok {
		for i := 0; i < len(p.Parres); i = +1 {
			if p.Types[i] == "rejected" {
				a = append(a, p.Parres[i])
			}
		}
	}
	return a
}
func GetDeclined(ID string, IDFrom int) []StructForREST {
	a := []StructForREST{}
	p, ok := ParryMems[ID]
	if ok {
		for i := 0; i < len(p.Parres); i = +1 {
			if p.Types[i] == "declined" {
				a = append(a, p.Parres[i])
			}
		}
	}
	return a
}
func IsAddParry(ArenaID string, ToAccountID int, FromAccountID int) bool {
	p, ok := ParryMems[ArenaID]

	if !ok {
		return true
	} else {
		for index, value := range p.To {
			if value == ToAccountID && FromAccountID == p.From[index] {
				return true
			}
			if value == FromAccountID && ToAccountID == p.From[index] {
				return true
			}
		}
		return false
	}
}
func VerifyActive(Tocken string, ArenaID string, accountIDTo int, bet float32) {
	//ok2 := mem.Arena.FindArenaEnd()
	_, ok := ParryMems[ArenaID]
	if ok {
		a := mem.Arena.FindArena(Tocken)
		pers, ok := PersonStruct.FindPersonByToken(Tocken)
		if ok {
			a.AddNewParry(pers.AccountID, accountIDTo, bet)
		}
	}
}
func VerifyReject(Tocken string, ArenaID string, accountIDTo int, bet float32) {
	p, ok := ParryMems[ArenaID]
	if ok {
		for index, value := range p.Types {
			if value == "pending" {
				p.DeleteADVbyIndex(index)
				VerifyDecline(Tocken, ArenaID, accountIDTo, bet)
				return
			}
			if value == "incoming" {
				p.DeleteADVbyIndex(index)
				VerifyDecline(Tocken, ArenaID, accountIDTo, bet)
				return
			}
		}
	}
}
func VerifyDecline(Tocken string, ArenaID string, accountIDTo int, bet float32) {
	p, ok := ParryMems[ArenaID]
	if ok {
		for index, value := range p.Types {
			if value == "pending" {
				p.DeleteADVbyIndex(index)
				VerifyDecline(Tocken, ArenaID, accountIDTo, bet)
				return
			}
			if value == "incoming" {
				p.DeleteADVbyIndex(index)
				VerifyDecline(Tocken, ArenaID, accountIDTo, bet)
				return
			}
		}
	}
}
func DeleteLongTocken() {

}
