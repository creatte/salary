package model
//非协会会员
type NoAffiliation struct{
	Members Affiliation
}

func (noaffiliation NoAffiliation) CalculateDeductions(pc paycheck) float64{
	return 0
}

