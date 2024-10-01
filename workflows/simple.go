package workflows

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

// Workflow definition
func SimpleWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Simple workflow started.", "name", name)

	// Simulate some processing
	workflow.Sleep(ctx, time.Second*5)

	return "Hello, " + name, nil
}
