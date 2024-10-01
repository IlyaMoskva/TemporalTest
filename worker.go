package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
)

func main() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Workflow options
	options := client.StartWorkflowOptions{
		ID:        "simple-workflow",
		TaskQueue: "my-task-queue",
	}

	// Start the workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, "SimpleWorkflow", "Temporal")
	if err != nil {
		panic(err)
	}

	// Wait for the result
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Workflow result: %s\n", result)
}
