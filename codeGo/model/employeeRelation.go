package model

type EmployeeCollaboratedWithEmployee struct{
	EmployeeKnows Employee
	TaskDescription string
	Employee Employee
}

type EmployeeWorksInTools struct{
	Employee Employee
	Role string
	Period string
	Tool Tool
}

type EmployeeSkilledInSkills struct{
	Employee Employee
	Experience string 
	Skills Skills
}

type employeeReportToEmployee struct{
	Employee Employee
	Manager Employee
}
