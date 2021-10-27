package database

type PayrollExceptionMessage struct{
	EmployeeNotExist string
	NotHoulyClassification string
	NotCommissionedClassification string
}

//为三种异常给予常量
const (
	EmployeeNotExist = "不存在该雇员"
	NotHoulyClassification = "该雇员不是时薪雇员"
	NotCommissionedClassification = "该雇员不是佣金雇员"
)