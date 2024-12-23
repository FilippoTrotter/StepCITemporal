package stepcitemporal

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func StepCIWorkflow(ctx workflow.Context, input WorkflowInput) (string, error) {

	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    5,
	}
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
		RetryPolicy:         retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, RunStepCI, input.StepCIPath).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("error runnig stepCI: %v", err)
	}

	err = workflow.ExecuteActivity(ctx, SendEmail, input.EmailAddress, "StepCI Succeeded").Get(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("error sending email: %v", err)
	}

	return fmt.Sprintf("Success message sent to %s", input.EmailAddress), nil
}
