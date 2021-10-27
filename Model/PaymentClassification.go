package model
//工资分类
type PaymentClassfication interface{
	CalculatePay(pc Paycheck) float64;
}