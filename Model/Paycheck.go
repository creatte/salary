package model

import "time"

type Paycheck struct {
	itsPayPeriodStartDate time.Time	//开始时间
	itsPayPeriodEndDate   time.Time	//结束时间
	itsGrossPay           float64	//总工资
	itsNetPay             float64	//净工资
	itsDeductions         float64	//扣除工资
	itsFields 			  map[string]string	
}

func (paycheck Paycheck) NewPaycheck(payPeriodStart time.Time,payPeriodEnd time.Time) Paycheck{
	paycheck.itsPayPeriodStartDate = payPeriodStart
	paycheck.itsPayPeriodEndDate = payPeriodEnd
	return paycheck
} 

func (paycheck Paycheck) SetGrosspay(GrossPay float64){
	paycheck.itsGrossPay = GrossPay
}

func (paycheck Paycheck) SetDeductions(deductions float64){
	paycheck.itsDeductions = deductions
}

func (paycheck Paycheck) SetNetPay(netPay float64){
	paycheck.itsNetPay = netPay
}

func (paycheck Paycheck) GetPayPeriodEndDate() time.Time{
	return paycheck.itsPayPeriodEndDate
}

func (paycheck Paycheck) GetGrossPay() float64{
	return paycheck.itsGrossPay
}

func (paycheck Paycheck) GetField(name string) string{
	return paycheck.itsFields[name]
}

func (paycheck Paycheck) GetDeductions() float64{
	return paycheck.itsDeductions
}

func (paycheck Paycheck) GetNetPay() float64{
	return paycheck.itsNetPay
}

func (paycheck Paycheck) SetField(name string,value string){
	paycheck.itsFields[name] = value
}

func (paycheck Paycheck) GetPayPeriodStartDate() time.Time{
	return paycheck.itsPayPeriodStartDate
}