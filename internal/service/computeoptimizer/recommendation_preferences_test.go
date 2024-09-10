// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeoptimizer_test

import (
	"context"
	"fmt"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfcomputeoptimizer "github.com/hashicorp/terraform-provider-aws/internal/service/computeoptimizer"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testAccRecommendationPreferences_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.RecommendationPreferencesDetail
	resourceName := "aws_computeoptimizer_recommendation_preferences.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.ComputeOptimizerEndpointID)
			testAccPreCheckEnrollmentStatus(ctx, t, "Active")
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.ComputeOptimizerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRecommendationPreferencesDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRecommendationPreferencesConfig_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckRecommendationPreferencesExists(ctx, resourceName, &v),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enhanced_infrastructure_metrics"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("external_metrics_preference"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("inferred_workload_types"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("look_back_period"), knownvalue.StringExact("DAYS_32")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("preferred_resource"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("resource_type"), knownvalue.StringExact("Ec2Instance")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("savings_estimation_mode"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.ListExact(
						[]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact("AccountId"),
							}),
						}),
					),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("utilization_preference"), knownvalue.ListSizeExact(0)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckRecommendationPreferencesExists(ctx context.Context, n string, v *awstypes.RecommendationPreferencesDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ComputeOptimizerClient(ctx)

		output, err := tfcomputeoptimizer.FindRecommendationPreferencesByThreePartKey(ctx, conn, rs.Primary.Attributes["resource_type"], rs.Primary.Attributes["scope.0.name"], rs.Primary.Attributes["scope.0.value"])

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckRecommendationPreferencesDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).ComputeOptimizerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_computeoptimizer_recommendation_preferences" {
				continue
			}

			_, err := tfcomputeoptimizer.FindRecommendationPreferencesByThreePartKey(ctx, conn, rs.Primary.Attributes["resource_type"], rs.Primary.Attributes["scope.0.name"], rs.Primary.Attributes["scope.0.value"])

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Compute Optimizer Recommendation Preferences %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

const testAccRecommendationPreferencesConfig_basic = `
data "aws_caller_identity" "current" {}

resource "aws_computeoptimizer_recommendation_preferences" "test" {
  resource_type = "Ec2Instance"
  scope {
    name  = "AccountId"
    value = data.aws_caller_identity.current.account_id
  }

  look_back_period = "DAYS_32"
}
`
