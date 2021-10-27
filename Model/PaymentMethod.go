package model

type PaymentMethod interface{
	Pay(pc Paycheck);
}