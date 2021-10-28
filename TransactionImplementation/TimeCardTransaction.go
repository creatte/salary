package transactionimplementation

import (
	database "Salary/Database"
	"Salary/transactionapplication"
	"time"
)


type TimeCardTransaction struct {
	itsDate time.Time
	itsHours float64
	itsEmpid int
	transactionapplication.Transaction
}

func (ti TimeCardTransaction) NewTimeCardTransaction(date time.Time,hours float64,empid int){
	ti.itsDate = date
	ti.itsHours = hours
	ti.itsEmpid = empid
}

func (ti TimeCardTransaction) Execute(){
	e := database.PayrollDatabase{}.GetEmployee(ti.itsEmpid)
	if e != nil{
		pc := e.GetClassification()
		if pc
	}
}