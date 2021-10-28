package transactionimplementation

import (
	database "Salary/Database"
	modelimplementation "Salary/ModelImplementation"
	"Salary/transactionapplication"
	"fmt"
	"log"
	"reflect"
	"strconv"
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
		pc := e.GetClassification() 
		if reflect.TypeOf(pc) == reflect.TypeOf(modelimplementation.CommissionedClassification{}){
			cc := pc.(*modelimplementation.CommissionedClassification)
			sr := modelimplementation.SalesReceipt{}.NewSalesReceipt(sa.itsSalesDate,sa.itsAmount)
			cc.AddSalesReceipt(sr)
		}else{
			log.Println(fmt.Errorf(database.PayrollExceptionMessage{}.NotCommissionedClassification +
				strconv.Itoa(e.GetEmpid())))
		}
	}
}