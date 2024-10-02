package activities

import (
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
	"testing"
)

// TestGreetActivity tests the GreetActivity function
func TestGreetActivity(t *testing.T) {
	// Initialize Temporal test suite
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()

	// Register the GreetActivity
	env.RegisterActivity(GreetActivity)

	// Execute the activity
	result, err := env.ExecuteActivity(GreetActivity, "John Doe")

	// Check if the activity completed successfully
	assert.NoError(t, err)

	// Verify the result
	var greeting string
	err = result.Get(&greeting)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, John Doe!", greeting)
}
