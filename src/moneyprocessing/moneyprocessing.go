package moneyprocessing
import "PersonStruct"
import "tokenPLC"
func Registrtion(p PersonStruct) (err error) {
	return tockenplc.RegistrNewAccount(tockenplc.RegistrationInformation(){})

}
func Transfer(pFrom PersonStruct,pTo PersonStruct,value float32) (err error){
	return tockenplc.Transfer(tockenplc.TransferInformation(){})
}