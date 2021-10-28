package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
)

type ChangeNameTransaction struct{
	itsName string
	abstracttransactions.ChangeEmployeeTransaction
}

func (ch ChangeNameTransaction) newChangeNameTransaction(empid int,name string){
	ch.itsName = name
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeNameTransaction) Change(e model.Employee){
	e.SetName(ch.itsName)
}

