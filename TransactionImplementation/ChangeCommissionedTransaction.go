package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeCommissionedTransaction struct{
	itsSalary float64
	itsCommissionRate float64
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeCommissionedTransaction) NewChangeCommissionedTransaction(empid int,salary float64,commissionRate float64){
	ch.itsSalary = salary
	ch.itsCommissionRate = commissionRate
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeCommissionedTransaction) GetClassification() model.PaymentClassfication{
	return modelimplementation.CommissionedClassification{}.NewCommissionedClassification(ch.itsSalary,ch.itsCommissionRate)
}

func (ch ChangeCommissionedTransaction) GetSchedule() model.PaymentSchedule{
	return modelimplementation.MonthlySchedule{}
}