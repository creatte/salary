package model
//工资支付方式
type PaymentMethod interface{
	Pay(pc Paycheck);
}