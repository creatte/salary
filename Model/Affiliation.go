package model
//协会会员
type Affiliation interface{
	CalculateDeductions(pc Paycheck) float64;
}