package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	database "Salary/Database"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
	"reflect"
)

type ChangeUnaffiliatedTransaction struct{
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeUnaffiliatedTransaction) NewChangeUnaffiliatedTransaction(empid int){
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeUnaffiliatedTransaction) RecordMembership(e model.Employee){
	af := e.GetAffiliation()
	 if reflect.TypeOf(af) == reflect.TypeOf(modelimplementation.UnionAffiliation{}){
     	uaf := af.(modelimplementation.UnionAffiliation)
    	memberId := uaf.GetMemberId()
    	database.PayrollDatabase{}.RemoveUnionMember(memberId)
}
}

func (ch ChangeUnaffiliatedTransaction) GetAffiliation() model.Affiliation{
	return model.NoAffiliation{}
}