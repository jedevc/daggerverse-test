package main

import (
	"context"
)

type Test struct{}

func (m *Test) Hello(ctx context.Context) (string, error) {
	res, err := dag.Foo().MyFunction("foo").Stdout(ctx)
	if err != nil {
		return "", err
	}

	res2, err := dag.Bar().MyFunction("bar").Stdout(ctx)
	if err != nil {
		return "", err
	}

	return res + res2, nil
}
