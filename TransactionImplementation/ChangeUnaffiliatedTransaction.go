package transactionimplementation

import (
	abstracttransactions "Salary/AbstractTransactions"
	model "Salary/Model"
	modelimplementation "Salary/ModelImplementation"
)

type ChangeUnaffiliatedTransaction struct{
	abstracttransactions.ChangeAffiliationTransaction
}

func (ch ChangeUnaffiliatedTransaction) NewChangeUnaffiliatedTransaction(empid int){
	abstracttransactions.ChangeAffiliationTransaction{}.NewChangeEmployeeTransaction(empid)
}

func (ch ChangeUnaffiliatedTransaction) RecordMembership(e model.Employee){
	af := e.GetAffiliation()
	// if (af is ModelImplementation.UnionAffiliation)
    //         {
    //             ModelImplementation.UnionAffiliation uaf = (ModelImplementation.UnionAffiliation)af;
    //             int memberId = uaf.GetMemberId();
    //             DataBase.PayrollDatabase.Default.RemoveUnionMember(memberId);
    //         }
}

func (ch ChangeUnaffiliatedTransaction) GetAffiliation() model.Affiliation{
	return model.NoAffiliation{}
}