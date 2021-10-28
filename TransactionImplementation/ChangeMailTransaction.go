package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeMailTransaction struct{
	itsAddress string
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeMailTransaction) NewChangeMailTransaction(empid int,address string){
	ch.itsAddress = address
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeMailTransaction) GetMethod() model.PaymentMethod{
	return modelimplementation.MailMethod{}.NewMailMethod(ch.itsAddress)
}

