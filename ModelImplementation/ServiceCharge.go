package modelimplementation

import "time"

type ServiceCharge struct {
	itsDate 	time.Time
	itsAmount 	float64
}

func (se ServiceCharge) newServiceCharge(date time.Time,amount float64){
	se.itsDate = date
	se.itsAmount =amount
}

func (se ServiceCharge) ServiceChargeGetAmount() float64{
	return se.itsAmount
}