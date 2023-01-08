package queries

import (
	"errors"
)

type builder struct {
	token  Token
	inside Inside
	fn     QueryFn
}

func createBuilder() Builder {
	out := builder{
		token:  nil,
		inside: nil,
		fn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token Token) Builder {
	app.token = token
	return app
}

// WithInside adds an inside to the builder
func (app *builder) WithInside(inside Inside) Builder {
	app.inside = inside
	return app
}

// WithFn adds a func to the builder
func (app *builder) WithFn(fn QueryFn) Builder {
	app.fn = fn
	return app
}

// Now builds a new Query instance
func (app *builder) Now() (Query, error) {
	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Query instance")
	}

	if app.inside == nil {
		return nil, errors.New("the inside is mandatory in order to build a Query instance")
	}

	if app.fn == nil {
		return nil, errors.New("the func is mandatory in order to build a Query instance")
	}

	return createQuery(app.token, app.inside, app.fn), nil
}
