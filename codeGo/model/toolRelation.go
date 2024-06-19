package model


type ToolRequestByEmployee struct {
	Tool          Tool
	RequestedData string
	AccessType string
	Lvl1Approval string
	Lvl2Approval string
	Employee      Employee
}

type ToolAccessToEmployee struct{
	Tool Tool
	AccessType string
	Employee Employee
}

type ToolManagedByEmployee struct{
	Tool Tool
	Employee Employee
}

type ToolComesUnderDepartment struct{
	Tool Tool
	Employee Employee
}
