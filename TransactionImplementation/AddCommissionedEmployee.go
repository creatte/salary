package transactionimplementation

import (
	"Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type AddCommissionedEmployee struct{
	itsSalary float64
	itsCommissionRate float64
	abstracttransactions.AddEmployeeTransaction
}

func (ad AddCommissionedEmployee) newAddCommissionedEmployee(empid int,name string,address string,salary float64,commissionRate float64){
	ad.itsSalary = salary
	ad.itsCommissionRate = commissionRate
	abstracttransactions.AddEmployeeTransaction{}.NewAddEmployeeTransaction(empid,name,address)
}

func (ad AddCommissionedEmployee) GetClassification() model.PaymentClassfication{
	return modelimplementation.CommissionedClassification{}.NewCommissionedClassification(ad.itsSalary,ad.itsCommissionRate)
}

func (ad AddCommissionedEmployee) GetSchedule() model.PaymentSchedule{
	return modelimplementation.BiweeklySchedule{}
}