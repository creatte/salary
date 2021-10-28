package modelimplementation

import (
	model "Salary/Model"
	"time")


type MonthlySchedule struct {
	model.PaymentSchedule
}

func (mo MonthlySchedule) IsLastDayOfMonth(date time.Time) bool {
	m1 := date.Month()
	m2 := date.AddDate(0,0,1).Month()
	return (m1 != m2)
}

func (mo MonthlySchedule) IsPaydate(payDate time.Time) bool{
	return mo.IsLastDayOfMonth(payDate)
}

func (mo MonthlySchedule) GetPayPeriodStartDate(payPeriodEndDate time.Time) time.Time{
	lastDayOfMonth := payPeriodEndDate.Day()
	firstDayOfMonth := payPeriodEndDate.AddDate(0,0,-(lastDayOfMonth - 1))
	return firstDayOfMonth
}