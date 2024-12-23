package stepcitemporal

import (
	"context"
	"fmt"
	"os/exec"

	"go.temporal.io/sdk/activity"
)

func RunStepCI(ctx context.Context, yamlPath string) (string, error) {
	cmd := exec.CommandContext(ctx, "npx", "stepci", "run", yamlPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Include StepCI's output in the error message
		return "", fmt.Errorf("stepci failed: %w\nOutput: %s", err, string(output))
	}
	return string(output), nil
}
func SendEmail(ctx context.Context, emailAdress, content string) (string, error) {
	activity.GetLogger(ctx).Info("Sending email to user", "EmailAddress", emailAdress)
	return content + " sent to " + emailAdress, nil
}
