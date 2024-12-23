package stepcitemporal_test

import (
	"stepcitemporal"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(stepcitemporal.RunStepCI, mock.Anything, mock.Anything).Return("PASS", nil)
	env.OnActivity(stepcitemporal.SendEmail, mock.Anything, mock.Anything, mock.Anything).Return("PASS send to example@test.com", nil)

	env.ExecuteWorkflow(stepcitemporal.StepCIWorkflow, stepcitemporal.WorkflowInput{EmailAddress: "example@test.com"})

	var result string
	assert.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, "Success message sent to example@test.com", result)
}
