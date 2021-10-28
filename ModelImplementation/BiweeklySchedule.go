package modelimplementation

import "time"

type BiweeklySchedule struct {
}

func (bi BiweeklySchedule) IsPaydate(payDate time.Time) bool {
	firstPayableFriday := time.Date(2001,11,9,0,0,0,0,time.Local)
	now := time.Now()
	ts := now.Sub(firstPayableFriday)//用当前时间-firstPayableFriday
	daysSinceFirstPayableFriday := int(ts.Hours()/24) //算出相差天数
	return daysSinceFirstPayableFriday%14 == 0
}

func (bi BiweeklySchedule) GetPayPeriodStartDate(payPeriodEndDate time.Time) time.Time{
	return payPeriodEndDate.AddDate(0,0,-13) //获取13天得的时间
}