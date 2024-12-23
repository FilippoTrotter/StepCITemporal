package stepcitemporal_test

import (
	"os"
	"stepcitemporal"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

func TestRunStepCI(t *testing.T) {
	// Set up test environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()

	// Register the RunStepCI activity
	env.RegisterActivity(stepcitemporal.RunStepCI)

	// Success case: valid Temporal YAML file
	t.Run("Success", func(t *testing.T) {
		// Create a temporary valid YAML file for success case
		validYamlContent := `version: 1
tests:
  example:
    steps:
      - name: GET request
        http:
          url: https://httpbin.org/status/200
          method: GET
          check:
            status: 200`
		tmpFile, err := os.CreateTemp("", "valid-temporal-*.yaml")
		assert.NoError(t, err)
		defer os.Remove(tmpFile.Name()) // Clean up the file after the test

		// Write content to the temporary file
		_, err = tmpFile.Write([]byte(validYamlContent))
		assert.NoError(t, err)

		// Execute the activity with the valid YAML path
		val, err := env.ExecuteActivity(stepcitemporal.RunStepCI, tmpFile.Name())
		assert.NoError(t, err)

		var result string
		val.Get(&result)

		// Assert the successful result
		assert.Contains(t, result, "PASS")
	})

	// Failure case: invalid Temporal YAML file
	t.Run("Failure", func(t *testing.T) {
		// Create a temporary invalid YAML file for failure case
		invalidYamlContent := `version: 1
tests:
  example:
    steps:
      - name: GET request
        http:
          url: nofound
          method: GET
          check:
            status: 200`
		tmpFile, err := os.CreateTemp("", "invalid-temporal-*.yaml")
		assert.NoError(t, err)
		defer os.Remove(tmpFile.Name()) // Clean up the file after the test

		// Write content to the temporary file
		_, err = tmpFile.Write([]byte(invalidYamlContent))
		assert.NoError(t, err)

		// Execute the activity with the invalid YAML path
		_, err = env.ExecuteActivity(stepcitemporal.RunStepCI, tmpFile.Name())
		assert.Error(t, err)
	})
}

// TestGetLocationInfo tests the GetLocationInfo activity with a mock server.
func TestSendMail(t *testing.T) {
	// set up test environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()

	env.RegisterActivity(stepcitemporal.SendEmail)

	val, err := env.ExecuteActivity(stepcitemporal.SendEmail, "test@example.com", "Test pass")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var location string
	val.Get(&location)

	expectedLocation := "Test pass sent to test@example.com"
	assert.Equal(t, location, expectedLocation)
}
