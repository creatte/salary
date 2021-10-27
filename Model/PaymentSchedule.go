package model

import "time"

type PaymentMethod interface {
	IsPaydate(payDate time.Time) bool;
	GetPayPeriodStartDate(GetPayPeriodEndDate time.Time);
}