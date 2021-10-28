package modelimplementation

import model "Salary/Model"

type SalariedClassification struct {
	itsSalary float64
	model.PaymentClassfication
}

func (sa SalariedClassification) SalariedClassification(salary float64) {
	sa.itsSalary = salary
}

func (sa SalariedClassification) GetSalary() float64 {
	return sa.itsSalary
}

func (sa SalariedClassification) CalculatePay(pc model.Paycheck) float64{
	return sa.itsSalary
}