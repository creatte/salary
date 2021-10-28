package transactionimplementation

import (
	database "Salary/Database"
	model "Salary/Model"
	"Salary/transactionapplication"
	"time"
)



type PaydayTransaction struct{
	transactionapplication.Transaction
	itsPayDate time.Time
	itsPaychecks map[int]model.Paycheck
}

func (pa PaydayTransaction) NewPaydayTransaction(payDate time.Time){
	pa.itsPayDate = payDate
}

func (pa PaydayTransaction) Execute(){
	empIds := database.PayrollDatabase{}.GetAllEmployeeIds()
	for empId := range empIds{
		e := database.PayrollDatabase{}.GetEmployee(empId)
		if (e != nil){
			if(e.IsPayDate(pa.itsPayDate)){
				pc := model.Paycheck{}.NewPaycheck(e.GetPayPeriodStartDate(pa.itsPayDate),pa.itsPayDate)
				pa.itsPaychecks[empId] = pc
				e.Payday(pc)
			}
		}
	}
}