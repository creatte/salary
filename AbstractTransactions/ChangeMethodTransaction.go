package abstracttransactions

import (
	"Salary/Model"	
	"fmt"
)


type ChangeMethodTransaction struct {
	ChangeMethodTransactioner
	ChangeEmployeeTransaction
	model.PaymentMethod
}

type ChangeMethodTransactioner interface {
	GetMethod() model.PaymentMethod
}

func (ch ChangeMethodTransaction) newChangeMethodTransaction(empid int) ChangeMethodTransaction{
	ch.itsEmpid = empid
	return ch
}

func (ch ChangeMethodTransaction) change(e model.Employee){
	pm := ch.GetMethod()
	e.SetMethod(pm)
}

func (ch ChangeMethodTransaction) Pay(pc model.Paycheck){
	err := "The method or operation is not implemented."
	fmt.Println(err)
}