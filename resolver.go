package hello_world

import (
	"context"
	"fmt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "world", nil
}

func (r *queryResolver) Foo(ctx context.Context, blah string) (string, error) {
	return fmt.Sprintf("this is blah: %s", blah), nil
}
