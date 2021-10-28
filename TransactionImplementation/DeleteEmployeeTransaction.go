package transactionimplementation

import (
	"Salary/transactionapplication"
	"Salary/Database"
)

type DeleteEmployeeTransaction struct{
	itsEmpid int
	transactionapplication.Transaction
}

func (de DeleteEmployeeTransaction) DeleteEmployeeTransaction(empid int){
	de.itsEmpid = empid
}

func (de DeleteEmployeeTransaction) Execute(){
	database.PayrollDatabase{}.DeleteEmployee(de.itsEmpid)
}