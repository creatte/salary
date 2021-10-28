package modelimplementation

import "Salary/Model"

type HoldMethod struct {
	model.PaymentMethod
}

func (ho HoldMethod) Pay(pc model.Paycheck){
	pc.SetField("Disposition", "Hold")
}