package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeSalariedTransaction struct{
	itsSalary float64
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeSalariedTransaction) NewChangeSalariedTransaction(empid int,salary float64){
	ch.itsSalary = salary
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeSalariedTransaction) GetClassification() model.PaymentClassfication{
	return modelimplementation.SalariedClassification{}.NewSalariedClassification(ch.itsSalary)
}

func (ch ChangeSalariedTransaction) GetSchedule() model.PaymentSchedule{
	return modelimplementation.MonthlySchedule{}
}