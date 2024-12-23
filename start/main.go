package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"stepcitemporal"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	// Define flags for email and YAML path
	email := flag.String("email", "", "Recipient email address")
	yamlPath := flag.String("path", "", "Path to the StepCI YAML script")

	// Parse command-line arguments
	flag.Parse()

	// Validate inputs
	if *email == "" || *yamlPath == "" {
		log.Fatalln("Both --email and --path are required")
	}

	// Create a Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Workflow options
	workflowOptions := client.StartWorkflowOptions{
		ID:        "stepci-workflow" + uuid.New().String(),
		TaskQueue: stepcitemporal.TaskQueueName,
	}

	// Workflow input
	input := stepcitemporal.WorkflowInput{
		StepCIPath:   *yamlPath,
		EmailAddress: *email,
	}

	// Execute the workflow
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, stepcitemporal.StepCIWorkflow, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	fmt.Printf("Started workflow: WorkflowID=%s, RunID=%s\n", we.GetID(), we.GetRunID())
}
