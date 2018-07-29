package bizrule

import "errors"

type Rule interface {
	Apply(*Context) (bool, error)
	CanBeApplied(*Context) bool
}

func New(in <-chan struct{}) (rule Rule, out <-chan struct{}) {
	return nil, nil
}

func NewContext() *Context {
	return &Context{storage: make(map[string]func(ctx *Context))}
}

type Context struct {
	storage map[string]func(ctx *Context)
}

func (ctx *Context) Describe(situation string, behaviour func(ctx *Context)) {
	ctx.storage[situation] = behaviour
}

func It(description string, action func() error) error {
	if err := action(); err != nil {
		return errors.New(description + ": " + err.Error())
	}
	return nil
}
