package grammars

import (
	"errors"

	"github.com/steve-care-software/ast/domain/grammars/values"
)

type composeElementBuilder struct {
	value      values.Value
	occurences uint
}

func createComposeElementBuilder() ComposeElementBuilder {
	out := composeElementBuilder{
		value:      nil,
		occurences: 0,
	}

	return &out
}

// Create initializes the builder
func (app *composeElementBuilder) Create() ComposeElementBuilder {
	return createComposeElementBuilder()
}

// WithValue adds a value to the builder
func (app *composeElementBuilder) WithValue(value values.Value) ComposeElementBuilder {
	app.value = value
	return app
}

// WithOccurences add occurences to the builder
func (app *composeElementBuilder) WithOccurences(occurences uint) ComposeElementBuilder {
	app.occurences = occurences
	return app
}

// Now builds a new ComposeElement instance
func (app *composeElementBuilder) Now() (ComposeElement, error) {
	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a ComposeElement instance")
	}

	if app.occurences <= 0 {
		return nil, errors.New("there must be at least 1 occurence in order to build a ComposeElement instance")
	}

	return createComposeElement(app.value, app.occurences), nil
}
