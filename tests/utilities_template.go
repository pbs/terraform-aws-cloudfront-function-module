package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testTemplate(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	arn := terraform.Output(t, terraformOptions, "arn")

	accountID := getAWSAccountID(t)
	expectedName := fmt.Sprintf("example-tf-cloudfront-function-%s", variant)

	expectedARN := fmt.Sprintf("arn:aws:cloudfront::%s:function/%s", accountID, expectedName)

	assert.Equal(t, expectedARN, arn)
}
