package modelimplementation

import (
	"time"
	"Salary/Model"
)

type CommissionedClassification struct {
	itsCommissionRate float64
	itsSalary         float64
	itsReceipts       map[time.Time]SalesReceipt
}

func (co CommissionedClassification) newCommissionedClassification(salary float64,commissionRate float64){
	co.itsSalary = salary
	co.itsCommissionRate = commissionRate
}

func (co CommissionedClassification) GetSalary() float64{
	return co.itsSalary
}

func (co CommissionedClassification) AddSalesReceipt(sr SalesReceipt){
	co.itsReceipts[sr.GetSaleDate()] = sr
}

func (co CommissionedClassification) GetReceipt(date time.Time) SalesReceipt{
	return co.itsReceipts[date]
}

func (co CommissionedClassification)GetRate() float64{
	return co.itsCommissionRate
}

var Members model.PaymentClassfication

func (co CommissionedClassification) CaculatePay(pc model.Paycheck) float64{
	commission := 0.0
	for key,value := range co.itsReceipts{
		recepit := value
		dt1 := pc.GetPayPeriodStartDate()
		dt2 := pc.GetPayPeriodEndDate()
		//time1.Before(time2)为time2-time1,返回bool
		if dt1.Before(key) && key.Before(dt2){
			commission += recepit.GetAmount() * co.itsCommissionRate
		}
	}
	return co.itsSalary + commission
}
