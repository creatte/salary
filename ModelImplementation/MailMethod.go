package modelimplementation

import model "Salary/Model"

type MailMethod struct {
	itsAddress string
	model.PaymentMethod
}

func (ma MailMethod) newMailMethod(address string) MailMethod{
	ma.itsAddress = address
	return ma
}

func (ma MailMethod) GetAddress() string {
	return ma.itsAddress
}

func (ma MailMethod) Pay(pc model.Paycheck){
	pc.SetField("Disposition", "Mail")
}