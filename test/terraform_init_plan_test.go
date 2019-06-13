package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformInitPlan(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		VarFiles: []string{"test/terraform.tfvars"},

        EnvVars: map[string]string{
            "AWS_DEFAULT_REGION": "eu-central-1",
      },
	}

	// This will run `terraform init` and `terraform plan` and fail the test if there are any errors
	terraform.InitAndPlan(t, terraformOptions)
}
