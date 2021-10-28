package abstracttransactions

import (
	"Salary/Model"
	"Salary/ModelImplementation"
	"Salary/Database"
)

type AddEmployeeTransaction struct{
	AddEmployeeTransactioner //内嵌，类似于java继承
	itsEmpid int	//员工编号
	itsName string	//员工姓名
	itsAddress string	//员工
}

type AddEmployeeTransactioner interface{
	GetClassification() model.PaymentClassfication
	GetSchedule() model.PaymentSchedule
}

func (add AddEmployeeTransaction) newAddEmployeeTransaction(empid int,name string,address string) AddEmployeeTransaction{
	add.itsEmpid = empid
	add.itsName = name
	add.itsAddress = address
	return add
}

func (add AddEmployeeTransaction) Execute(){
	pc := add.GetClassification()
	ps := add.GetSchedule()
	pm := modelimplementation.HoldMethod{}
	e := model.Employee{}.NewEmployee(add.itsEmpid,add.itsName,add.itsAddress)
	e.SetClassification(pc)
	e.SetSchedule(ps)
	e.SetMethod(pm)
	database.PayrollDatabase{}.AddEmployee(add.itsEmpid,&e)
}