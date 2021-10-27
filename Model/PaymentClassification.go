package model

type PaymentClassfication interface{
	CalculatePay(pc Paycheck) float64;
}