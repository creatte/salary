package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	database "Salary/Database"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeMemberTransaction struct{
	itsDues 		float64
	itsMemberId 	int
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeMemberTransaction) NewChangeMemberTransaction(empid int,memberid int,dues float64){
	ch.itsMemberId = memberid
	ch.itsDues = dues
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeMemberTransaction) RecordMembership(e model.Employee){
	database.PayrollDatabase{}.AddUnionMember(ch.itsMemberId,&e)
}

func (ch ChangeMemberTransaction) GetAffiliation() model.Affiliation{
	return modelimplementation.UnionAffiliation{}.NewUnionAffiliation(ch.itsMemberId,ch.itsDues)
}