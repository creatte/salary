package database

import (
	"Salary/Model"
)


type PayrollDatabase struct{
	// private static PayrollDatabase theInstance = new PayrollDatabase();
	itsEmployees map[int]*model.Employee //因为需返回空，建议使用指针
	itsUnionMembers map[int]int
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
// public static PayrollDatabase Default
// {
// 	get
// 	{
// 		return theInstance;
// 	}
// }
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

// public List<int> GetAllEmployeeIds()
// {
// 	List<int> empIds = new List<int>(10);
// 	foreach (int empid in itsEmployees.Keys)
// 	{
// 		empIds.Add(empid);
// 	}

// 	return empIds;
// }
// theInstance := &PayrollDatabase{}