package modelimplementation

import model "Salary/Model"

type SalariedClassification struct {
	itsSalary float64
	model.PaymentClassfication
}

func (sa SalariedClassification) NewSalariedClassification(salary float64) SalariedClassification{
	sa.itsSalary = salary
	return sa
}

func (sa SalariedClassification) GetSalary() float64 {
	return sa.itsSalary
}

func (sa SalariedClassification) CalculatePay(pc model.Paycheck) float64{
	return sa.itsSalary
}