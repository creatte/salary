package model

import "time"

type Employee struct {
	itsEmpid   int
	itsName    string
	itsAddress string

	itsClassification PaymentClassfication
	itsSchedule       PaymentSchedule
	itsPaymentMethod  PaymentMethod
	itsAffiliation    Affiliation
}

func (employee Employee) SetEmployee(empid int, name string, address string) {
	employee.itsEmpid = empid
	employee.itsName = name
	employee.itsAddress = address
}

func (employee Employee) SetClassification(pc PaymentClassfication) {
	employee.itsClassification = pc
}

func (employee Employee) SetSchedule(ps PaymentSchedule) {
	employee.itsSchedule = ps
}

func (employee Employee) SetMethod(pm PaymentMethod) {
	employee.itsPaymentMethod = pm
}

func (employee Employee) GetName() string {
	return employee.itsName
}

func (employee Employee) GetClassification() PaymentClassfication {
	return employee.itsClassification
}

func (employee Employee) GetSchedule() PaymentSchedule {
	return employee.itsSchedule
}

func (employee Employee) GetMethod() PaymentMethod {
	return employee.itsPaymentMethod
}

func (employee Employee) SetAffiliation(af Affiliation) {
	employee.itsAffiliation = af
}

func (employee Employee) GetAffiliation() Affiliation {
	return employee.itsAffiliation
}

func (employee Employee) GetEmpid() int {
	return employee.itsEmpid
}
func (employee Employee) SetName(name string) {
	employee.itsName = name
}

func (employee Employee) SetAddress(address string) {
	employee.itsAddress = address
}

func (employee Employee) GetAddress() string {
	return employee.itsAddress
}

func (employee Employee) IsPayDate(payDate time.Time) bool {
	return employee.itsSchedule.IsPaydate(payDate)
}

func (employee Employee) GetPayPeriodStartDate(payPeriodEndDate time.Time) time.Time{
	return employee.itsSchedule.GetPayPeriodStartDate(payPeriodEndDate)
}

func (employee Employee) Payday(pc Paycheck) {
	grossPay := employee.itsClassification.CalculatePay(pc)
	deductions := employee.itsAffiliation.CalculateDeductions(pc)
	netPay := grossPay - deductions
	pc.SetGrosspay(grossPay)
	pc.SetDeductions(deductions)
	pc.SetNetPay(netPay)
	employee.itsPaymentMethod.Pay(pc)

}