// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An assessment created for a group in the Migration project.
// API Version: 2019-10-01.
func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:migrate:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	// Unique name of an assessment within a project.
	AssessmentName string `pulumi:"assessmentName"`
	// Unique name of a group within a project.
	GroupName string `pulumi:"groupName"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An assessment created for a group in the Migration project.
type LookupAssessmentResult struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this assessment. /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessment/{assessmentName}
	Id string `pulumi:"id"`
	// Unique name of an assessment.
	Name string `pulumi:"name"`
	// Properties of the assessment.
	Properties AssessmentPropertiesResponse `pulumi:"properties"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects/groups/assessments].
	Type string `pulumi:"type"`
}
