package programs

type instruction struct {
	value     Value
	execution Application
}

func createInstructionWithValue(
	value Value,
) Instruction {
	return createInstructionInternally(value, nil)
}

func createInstructionWithExecution(
	execution Application,
) Instruction {
	return createInstructionInternally(nil, execution)
}

func createInstructionInternally(
	value Value,
	execution Application,
) Instruction {
	out := instruction{
		value:     value,
		execution: execution,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *instruction) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *instruction) Value() Value {
	return obj.value
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *instruction) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *instruction) Execution() Application {
	return obj.execution
}
