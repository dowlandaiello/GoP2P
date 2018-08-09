package command

// Command - absctract container holding command values
type Command struct {
	Command string `json:"command"`

	Modifiers *Modifiers `json:"modifiers"`
}

// Modifiers - abstract containers holding specific parameters for a command
type Modifiers struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewCommand - attempt to initialize new instance of command struct with specified command, modifiers.
func NewCommand(command string, modifiers *Modifiers) (*Command, error) {
	return &Command{}, nil
}
