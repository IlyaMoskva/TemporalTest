package activities

import (
	"context"
	"fmt"
)

// Activity definition
func GreetActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
