package command

// Command - absctract container holding command values
type Command struct {
	Command string `json:"command"`

	Modifiers string `json:"modifiers"`
}

// Modifiers - abstract containers holding specific parameters for a command
type Modifiers struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
