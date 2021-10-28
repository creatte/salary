package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type AddSalariedEmployee struct{
	itsSalary float64
	abstracttransactions.AddEmployeeTransaction
}

func (ad AddSalariedEmployee) NewAddSalariedEmployee(empid int,name string,address string,salary float64){
	ad.itsSalary = salary
	abstracttransactions.AddEmployeeTransaction{}.NewAddEmployeeTransaction(empid,name,address)
}

func (ad AddSalariedEmployee) GetClassification() model.PaymentClassfication{
	return modelimplementation.SalariedClassification{}.NewSalariedClassification(ad.itsSalary)
}

func (ad AddSalariedEmployee) GetSchedule() model.PaymentSchedule{
	return modelimplementation.MonthlySchedule{}
}