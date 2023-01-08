package applications

import (
	"github.com/steve-care-software/interpreter/domain/instructions"
	"github.com/steve-care-software/interpreter/domain/programs"
	"github.com/steve-care-software/interpreter/domain/programs/modules"
)

// NameBytesToString converts a name []byte to a string
type NameBytesToString func(name []byte) string

// NewApplication creates a new application
func NewApplication(
	nameBytesToStringFn NameBytesToString,
) Application {
	builder := programs.NewBuilder()
	instructionsBuilder := programs.NewInstructionsBuilder()
	instructionBuilder := programs.NewInstructionBuilder()
	applicationBuilder := programs.NewApplicationBuilder()
	attachmentsBuilder := programs.NewAttachmentsBuilder()
	attachmentBuilder := programs.NewAttachmentBuilder()
	valueBuilder := programs.NewValueBuilder()
	return createApplication(
		builder,
		instructionsBuilder,
		instructionBuilder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		valueBuilder,
		nameBytesToStringFn,
	)
}

// Application represents a program application
type Application interface {
	Compile(modules modules.Modules, instructions instructions.Instructions) (programs.Program, error)
	Execute(input []interface{}, program programs.Program) ([]interface{}, error)
}
