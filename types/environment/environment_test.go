package environment

import (
	"testing"

	"github.com/dowlandaiello/GoP2P/common"
)

// TestNewEnvironment - test functionality of NewEnvironment() function
func TestNewEnvironment(t *testing.T) {
	env, err := NewEnvironment() // Initialize new environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	serializedEnv, err := common.SerializeToBytes(env) // Serialize environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	envID := common.Sha3(serializedEnv) // Get hash value of environment

	t.Logf("created environment with ID %s", envID) // Log success
}

// TestQueryType - test functionality of QueryType() function
func TestQueryType(t *testing.T) {
	env, err := NewEnvironment() // Initialize new environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	variable, err := NewVariable("string", "test") // Create new string variable

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = env.AddVariable(variable, false) // Add variable to environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	foundVariable, err := env.QueryType(variable.VariableType) // Query variable type

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found variable %s", foundVariable.VariableIdentifier) // Log success
}

// TestQueryValue - test functionality of QueryValue() function
func TestQueryValue(t *testing.T) {
	env, err := NewEnvironment() // Initialize new environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	variable, err := NewVariable("string", "test") // Create new string variable

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = env.AddVariable(variable, false) // Add variable to environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	foundVariable, err := env.QueryValue("test") // Query variable value

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found variable %s", foundVariable.VariableIdentifier) // Log success
}

// TestNewVariable - test functionality of variable initialization function
func TestNewVariable(t *testing.T) {
	variable, err := NewVariable("string", "test") // Create new string variable

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("created variable %s", variable) // Log success
}
