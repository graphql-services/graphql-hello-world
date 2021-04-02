package hello_world

import (
	"context"
	"fmt"
	"net/http"
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

func (r *queryResolver) Header(ctx context.Context, name string) (string, error) {
	headers := ctx.Value("headers").(http.Header)
	return headers.Get(name), nil
}
