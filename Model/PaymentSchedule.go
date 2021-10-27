package model

import "time"

type PaymentSchedule interface {
	IsPaydate(payDate time.Time) bool;
	GetPayPeriodStartDate(GetPayPeriodEndDate time.Time) time.Time;
}