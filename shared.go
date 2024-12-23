package stepcitemporal

type WorkflowInput struct {
	StepCIPath   string
	EmailAddress string
}

const TaskQueueName = "stepCI-test"
