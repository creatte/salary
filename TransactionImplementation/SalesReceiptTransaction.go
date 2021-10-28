package transactionimplementation

import (
	database "Salary/Database"
	"Salary/transactionapplication"
	"time"
)


type SalesReceiptTransaction struct {
	itsEmpid     int
	itsAmount    float64
	itsSalesDate time.Time
	transactionapplication.Transaction
}

func (sa SalesReceiptTransaction) NewSalesReceiptTransaction(saleDate time.Time,amount float64,empid int){
	sa.itsSalesDate = saleDate
	sa.itsAmount = amount
	sa.itsEmpid = empid
}

func (sa SalesReceiptTransaction) Execute(){
	e := database.PayrollDatabase{}.GetEmployee(sa.itsEmpid)
	if e != nil{
		pc := e.GetAddress() 
	}
}