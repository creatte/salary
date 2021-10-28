package transactionimplementation

import (
	transactionapplication "Salary/TransactionApplication"
	"time")


type ServiceChargeTransaction struct {
	itsDate time.Time
	itsCharge float64
	itsMemberId int
	transactionapplication.Transaction
}

func (se ServiceChargeTransaction) NewServiceChargeTransaction(memberId int,date time.Time,charge float64){
	se.itsMemberId = memberId
	se.itsDate = date
	se.itsCharge = charge
}

func (se ServiceChargeTransaction) Execute(){
	
}