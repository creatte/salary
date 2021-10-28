package abstracttransactions

import (
	database "Salary/Database"
	model "Salary/Model")


type ChangeEmployeeTransaction struct {
	ChangeEmployeeTransactioner
	itsEmpid int
}

type ChangeEmployeeTransactioner interface {
	Change(e model.Employee)
}

func (ch ChangeEmployeeTransaction) newChangeEmployeeTransaction(empid int) ChangeEmployeeTransaction{
	ch.itsEmpid = empid
	return ch
}

func (ch ChangeEmployeeTransaction) Execute(){
	e := database.PayrollDatabase{}.GetEmployee(ch.itsEmpid)
	if e != nil{
		ch.Change(*e)
	}
}