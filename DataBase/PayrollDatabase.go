package database

import (
	"Salary/Model"
)


type PayrollDatabase struct{
	
	itsEmployees map[int]*model.Employee //因为需返回空，建议使用指针
	itsUnionMembers map[int]int
}
// private static PayrollDatabase theInstance = new PayrollDatabase();
//全局变量theInstance
var theInstance = &PayrollDatabase{
	itsEmployees: make(map[int]*model.Employee),
	itsUnionMembers: make(map[int]int),
}

func (payrollDatabase PayrollDatabase) clear(){
	// for i := range payrollDatabase.itsEmployees {
	// 	delete(payrollDatabase.itsEmployees, i)
	// }
	// for j := range payrollDatabase.itsUnionMembers{
	// 	delete(payrollDatabase.itsUnionMembers,j)
	// }
	payrollDatabase.itsEmployees = make(map[int]*model.Employee)
	payrollDatabase.itsUnionMembers = make(map[int]int)
}

func (payrollDatabase PayrollDatabase) GetEmployee(empid int) *model.Employee{
	if its,ok := payrollDatabase.itsEmployees[empid];ok{
		return its
	}
	return nil
}

func (payrollDatabase PayrollDatabase) AddEmployee(empid int,e *model.Employee){
	payrollDatabase.itsEmployees[empid] = e
}

func (PayrollDatabase PayrollDatabase) PayrollDatabase(){
}

func (payrollDatabase PayrollDatabase) DeleteEmployee(empid int){
	delete(payrollDatabase.itsEmployees,empid)
}

func (payrollDatabase PayrollDatabase) AddUnionMember(memberid int,e *model.Employee){
	empid := e.GetEmpid()
	payrollDatabase.itsUnionMembers[empid] = memberid
}

func (payrollDatabase PayrollDatabase) GetUnionMember(memberid int) *model.Employee{
	if its,ok := payrollDatabase.itsUnionMembers[memberid];ok{
		return payrollDatabase.itsEmployees[its]
	}else{
		return nil
	}
}

func (payrollDatabase PayrollDatabase) RemoveUnionMember(memberid int){
	delete(payrollDatabase.itsUnionMembers,memberid)
}

func (payrollDatabase PayrollDatabase) GetAllEmployeeIds() []int{
	empIds := make([]int,0)
	for key,_ := range payrollDatabase.itsEmployees{
		empIds = append(empIds,key)
	}
	return empIds
}



// public static PayrollDatabase Default
// {
// 	get
// 	{
// 		return theInstance;
// 	}
// }
