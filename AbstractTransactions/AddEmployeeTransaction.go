package abstracttransactions

import "Salary/Model"

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

func (addEmployeeTransaction AddEmployeeTransaction) newAddEmployeeTransaction(empid int,name string,address string){
	addEmployeeTransaction.itsEmpid = empid
	addEmployeeTransaction.itsName = name
	addEmployeeTransaction.itsAddress = address
}

func (addEmployeeTransaction AddEmployeeTransaction) Execute(){
	// pc := addEmployeeTransaction.GetClassification()
	// ps := addEmployeeTransaction.GetSchedule()
}