package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeDirectTransaction struct{
	itsBank string
	itsAccount string
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeDirectTransaction) NewChangeDirectTransaction(empid int,bank string,account string){
	ch.itsBank = bank
	ch.itsAccount = account
	abstracttransactions.ChangeMethodTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeDirectTransaction) GetMethod() model.PaymentMethod{
	return modelimplementation.DirectMethod{}.NewDirctMethod(ch.itsBank,ch.itsAccount)
}
