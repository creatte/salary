package abstracttransactions

import (
	database "Salary/Database"
	model "Salary/Model"
	transactionapplication "Salary/TransactionApplication"
)


type ChangeEmployeeTransaction struct {
	ChangeEmployeeTransactioner
	transactionapplication.Transaction
	itsEmpid int
}

type ChangeEmployeeTransactioner interface {
	Change(e model.Employee)
}

func (ch ChangeEmployeeTransaction) NewChangeEmployeeTransaction(empid int) ChangeEmployeeTransaction{
	ch.itsEmpid = empid
	return ch
}

func (ch ChangeEmployeeTransaction) Execute(){
	e := database.PayrollDatabase{}.GetEmployee(ch.itsEmpid)
	if e != nil{
		ch.Change(*e)
	}
}