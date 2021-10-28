package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
)

type ChangeAddressTransaction struct{
	itsAddress string
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeAddressTransaction) NewChangeAddressTransaction(empid int,address string){
	ch.itsAddress = address
	abstracttransactions.ChangeEmployeeTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeAddressTransaction) Change(e model.Employee){
	e.SetAddress(ch.itsAddress)
}
