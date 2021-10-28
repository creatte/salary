package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeHourlyTransaction struct{
	itsHoursRate float64
	abstracttransactions.ChangeClassificationTransaction
}

func (ch ChangeHourlyTransaction) NewChangeHourlyTransaction(empid int,hourlyRate float64){
	ch.itsHoursRate = hourlyRate
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeHourlyTransaction) GetClassification() model.PaymentClassfication{
	return modelimplementation.HourlyClassification{}.NewHourlyClassification(ch.itsHoursRate)
}

func (ch ChangeHourlyTransaction) GetSchedule() model.PaymentSchedule{
	return modelimplementation.WeeklySchedule{}
}