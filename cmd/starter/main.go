package main

import (
	"TemporalTest/activities"
	"TemporalTest/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	// Create a Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create a worker on a specific task queue
	w := worker.New(c, "my-task-queue", worker.Options{})

	// Register workflow and activities
	w.RegisterWorkflow(workflows.SimpleWorkflow)
	w.RegisterActivity(activities.GreetActivity)

	// Start the worker to listen to task queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
