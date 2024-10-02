package workflows

import (
	"TemporalTest/activities"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
	"testing"
)

// TestSimpleWorkflow unit tests the SimpleWorkflow function
func TestSimpleWorkflow(t *testing.T) {
	// Initialize Temporal test suite
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Register the GreetActivity for the workflow
	env.RegisterActivity(activities.GreetActivity)

	// Execute the workflow
	env.ExecuteWorkflow(SimpleWorkflow, "John Doe")

	// Check if the workflow completed successfully
	assert.True(t, env.IsWorkflowCompleted())

	// Verify the result
	var result string
	err := env.GetWorkflowResult(&result)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, John Doe", result)
}
