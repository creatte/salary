package modelimplementation

import "time"

type WeeklySchedule struct {
}

func (we WeeklySchedule) IsPaydate(paydate time.Time) bool{
	return paydate.Weekday() == time.Friday
}

func (we WeeklySchedule) GetPayPeriodStartDate(payPeriodEndDate time.Time) time.Time{
	return payPeriodEndDate.AddDate(0,0,-6)
}