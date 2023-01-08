package programs

import (
	"errors"

	"github.com/steve-care-software/interpreter/domain/programs/modules"
)

type applicationBuilder struct {
	pIndex      *uint
	module      modules.Module
	attachments Attachments
}

func createApplicationBuilder() ApplicationBuilder {
	out := applicationBuilder{
		pIndex:      nil,
		module:      nil,
		attachments: nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder()
}

// WithIndex adds an index to the builder
func (app *applicationBuilder) WithIndex(index uint) ApplicationBuilder {
	app.pIndex = &index
	return app
}

// WithModule adds a module to the builder
func (app *applicationBuilder) WithModule(module modules.Module) ApplicationBuilder {
	app.module = module
	return app
}

// WithAttachments add attachments to the builder
func (app *applicationBuilder) WithAttachments(attachments Attachments) ApplicationBuilder {
	app.attachments = attachments
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Application instance")
	}

	if app.module == nil {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	if app.attachments != nil {
		return createApplicationWithAttachments(*app.pIndex, app.module, app.attachments), nil
	}

	return createApplication(*app.pIndex, app.module), nil
}
