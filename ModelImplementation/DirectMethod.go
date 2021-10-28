package modelimplementation

import model "Salary/Model"

//import  "Salary/Model"

type DirectMethod struct {
	itsBank    string
	itsAccount string
}

func (di DirectMethod) newDirctMethod(bank string, account string) {
	di.itsBank = bank
	di.itsAccount = account
}

func (di DirectMethod) GetBank() string {
	return di.itsBank
}

func (di DirectMethod) GetAccount() string {
	return di.itsAccount
}

//在CommissionedClassfication中已经定义，此处无需定义
//var Members model.PaymentMethod

func (di DirectMethod) Pay(pc model.Paycheck){
	pc.SetField("Disposition", "Direct")
}