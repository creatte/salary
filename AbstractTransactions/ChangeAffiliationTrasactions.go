package abstracttransactions

import (
	model"Salary/Model"	
	"fmt"
)


type ChangeAffiliationTransaction struct {
	ChangeEmployeeTransaction
	model.Affiliation
	ChangeAffiliationTransactioner
}

type ChangeAffiliationTransactioner interface {
	RecordMembership(e model.Employee)
	GetAffiliation() model.Affiliation
}

func (ch ChangeAffiliationTransaction) newChangeAffiliationTransaction(empid int) ChangeAffiliationTransaction{
	ch.itsEmpid = empid
	return ch
}

func (ch ChangeAffiliationTransaction) Change(e model.Employee){
	ch.RecordMembership(e)
	af := ch.GetAffiliation()
	e.SetAffiliation(af)
}
//全局定义会员成员
var Members model.Affiliation

func (ch ChangeAffiliationTransaction) CalculateDeductions(pc model.Paycheck){
	err := "The method or operation is not implemented."
	fmt.Println(err)
}