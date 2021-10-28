package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)


type AddHourlyEmployee struct {
	abstracttransactions.AddEmployeeTransaction
	itsHoursRate float64
}

func (ad AddHourlyEmployee) NewAddHourlyEmployee(empid int,name string,address string,hourlyRate float64){
	ad.itsHoursRate = hourlyRate
	abstracttransactions.AddEmployeeTransaction{}.NewAddEmployeeTransaction(empid,name,address)
}

func (ad AddHourlyEmployee) GetClassification() model.PaymentClassfication{
	return modelimplementation.HourlyClassification{}.NewHourlyClassification(ad.itsHoursRate)
}

func (ad AddHourlyEmployee) GetSchedule() model.PaymentSchedule{
	return modelimplementation.WeeklySchedule{}
}