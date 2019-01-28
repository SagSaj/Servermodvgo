package moneyprocessing
import "main/modules/PersonStruct"
import "main/modules/tokenPLC"
func Registrtion(p PersonStruct) (err error) {
	return tockenplc.RegistrNewAccount(tockenplc.RegistrationInformation(){})

}
func Transfer(pFrom PersonStruct,pTo PersonStruct,value float32) (err error){
	return tockenplc.Transfer(tockenplc.TransferInformation(){})
}