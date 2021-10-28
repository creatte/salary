package modelimplementation

import "time"

type TimeCard struct {
	itsDate time.Time
	itsHours float64
}

func (ti TimeCard) newTimeCard(date time.Time, hours float64){
	ti.itsDate = date
	ti.itsHours = hours
}

func (ti TimeCard) GetDate() time.Time{
	return ti.itsDate
}

func (ti TimeCard) GetHours() float64{
	return ti.itsHours
}