package main

import (
	"context"
)

type copyFunc func(context.Context, string, string, int) error
