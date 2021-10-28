package abstracttransactions

import model "Salary/Model"

type ChangeClassificationTransaction struct {
	ChangeClassificationTransactioner
	ChangeEmployeeTransaction	
}

type ChangeClassificationTransactioner interface {
	GetClassification() model.PaymentClassfication
	GetSchedule() model.PaymentSchedule
}

func (ch ChangeClassificationTransaction) ChangeClassificationTransaction(empid int) ChangeClassificationTransaction{
	ch.itsEmpid = empid
	return ch
}