package environment

// Environment - abstract container holding variables, configurations of a certain node
type Environment struct {
	EnvironmentVariables []*Variable
}

// Variable - container holding a variable's data (pointer), and identification properties (id, type)
type Variable struct {
	VariableType       string       `json:"type"`       // VariableType - type of variable (e.g. string, block, etc...)
	VariableIdentifier string       `json:"identifier"` // VariableIdentifier - id of variable (used for querying)
	VariableData       *interface{} `json:"data"`       // VariableData - pretty self-explanatory (usually a pointer to a struct)
}
