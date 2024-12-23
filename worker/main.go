package main

import (
	"log"
	"stepcitemporal"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, stepcitemporal.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(stepcitemporal.StepCIWorkflow)
	w.RegisterActivity(stepcitemporal.RunStepCI)
	w.RegisterActivity(stepcitemporal.SendEmail)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
