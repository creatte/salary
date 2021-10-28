package modelimplementation

import (
	model "Salary/Model"
	"math"
	"time"
)


type HourlyClassification struct {
	itsRate      float64
	itsTimeCards map[time.Time]TimeCard
}

func (ho HourlyClassification) newHourlyClassification(hourlyRate float64){
	ho.itsRate = hourlyRate
}

func (ho HourlyClassification) GetRate() float64{ 
	return ho.itsRate
}

func (ho HourlyClassification) GetTimeCard(date time.Time) TimeCard{
	return ho.itsTimeCards[date]
}

func (ho HourlyClassification) AddTimeCard(tc TimeCard){
	ho.itsTimeCards[tc.GetDate()] = tc
}

func (ho HourlyClassification) CalculatePayForTimeCard(tc TimeCard) float64{
	hours := tc.itsHours
	overtime := math.Max(0.0,hours-8.0)
	straightTime := hours - overtime
	return straightTime * ho.itsRate + overtime * ho.itsRate * 1.5
}

func (ho HourlyClassification) CalculatePay(pc model.Paycheck) float64{
	totalPay := 0.0
//	payPeriodEndDate := pc.GetPayPeriodEndDate()
	for key,value := range ho.itsTimeCards{
		tc := value

		dt1 := pc.GetPayPeriodStartDate()
		dt2 := pc.GetPayPeriodEndDate()

		if dt1.Before(key) && key.Before(dt2){
			totalPay += ho.CalculatePayForTimeCard(tc)
		}
	}
	return totalPay
}