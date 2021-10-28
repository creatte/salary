package modelimplementation

import "time"

type ServiceCharge struct {
	itsDate 	time.Time
	itsAmount 	float64
}

func (se ServiceCharge) newServiceCharge(date time.Time,amount float64) ServiceCharge{
	se.itsDate = date
	se.itsAmount =amount
	return se
}

func (se ServiceCharge) ServiceChargeGetAmount() float64{
	return se.itsAmount
}