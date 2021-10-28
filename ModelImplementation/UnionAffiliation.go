package modelimplementation

import (
	model "Salary/Model"
	"time")


type UnionAffiliation struct {
	itsMemberId       int
	itsDues           float64
	itsServiceCharges map[time.Time]ServiceCharge
}

func (un UnionAffiliation) newUnionAffiliation(memberId int,dues float64) UnionAffiliation{
	un.itsMemberId = memberId
	un.itsDues = dues
	return un
}

func (un UnionAffiliation) GetServiceCharge(date time.Time) ServiceCharge{
	return un.itsServiceCharges[date]
}

func (un UnionAffiliation) AddServiceCharge(date time.Time,amount float64){
	sc := ServiceCharge{date,amount}
	un.itsServiceCharges[date] = sc
}

func (un UnionAffiliation) GetDues() float64{
	return un.itsDues
}

func (un UnionAffiliation) GetMemberId() int{
	return un.itsMemberId
}

func (un UnionAffiliation) CalculateDeductions(pc model.Paycheck) float64{
	totalServiceCharge := 0.0;
    totalDues := 0.0;
	for key,valuse := range un.itsServiceCharges{
		sc := valuse
		dt1 := pc.GetPayPeriodStartDate()
		dt2 := pc.GetPayPeriodEndDate()
		if dt1.Before(key) && key.Before(dt2){
			totalServiceCharge += sc.ServiceChargeGetAmount()
		}
	}
	fridays := un.NumberOfFridaysInPayPeriod(pc.GetPayPeriodStartDate(),pc.GetPayPeriodEndDate())
	totalDues = un.itsDues * float64(fridays)
	return totalDues + totalServiceCharge
}

func (un UnionAffiliation) NumberOfFridaysInPayPeriod(payPeriodStart time.Time,payPeriodEnd time.Time) int{
	fridays := 0
	for date := payPeriodStart; date.Before(payPeriodEnd); date = date.AddDate(0,0,1){
		if date.Weekday() == time.Friday{
			fridays++
		}
	}
	return fridays
}