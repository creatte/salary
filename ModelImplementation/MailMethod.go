package modelimplementation

import model "Salary/Model"

type MailMethod struct {
	itsAddress string
}

func (ma MailMethod) newMailMethod(address string) {
	ma.itsAddress = address
}

func (ma MailMethod) GetAddress() string {
	return ma.itsAddress
}

func (ma MailMethod) Pay(pc model.Paycheck){
	pc.SetField("Disposition", "Mail")
}