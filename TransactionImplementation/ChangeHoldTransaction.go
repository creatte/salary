package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeHoldTransaction struct{
	abstracttransactions.ChangeMethodTransaction
}

func (ch ChangeHoldTransaction) NewChangeHoldTransaction(empid int){
	abstracttransactions.ChangeMethodTransaction{}.NewChangeEmployeeTransaction(empid)	
} 

func (ch ChangeHoldTransaction) GetMethod() model.PaymentMethod{
	return modelimplementation.HoldMethod{}
}