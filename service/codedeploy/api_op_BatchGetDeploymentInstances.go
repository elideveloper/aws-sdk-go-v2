// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a BatchGetDeploymentInstances operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetDeploymentInstancesInput
type BatchGetDeploymentInstancesInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment.
	//
	// DeploymentId is a required field
	DeploymentId *string `locationName:"deploymentId" type:"string" required:"true"`

	// The unique IDs of instances used in the deployment. The maximum number of
	// instance IDs you can specify is 25.
	//
	// InstanceIds is a required field
	InstanceIds []string `locationName:"instanceIds" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetDeploymentInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetDeploymentInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetDeploymentInstancesInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a BatchGetDeploymentInstances operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetDeploymentInstancesOutput
type BatchGetDeploymentInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Information about errors that might have occurred during the API call.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// Information about the instance.
	InstancesSummary []InstanceSummary `locationName:"instancesSummary" type:"list"`
}

// String returns the string representation
func (s BatchGetDeploymentInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetDeploymentInstances = "BatchGetDeploymentInstances"

// BatchGetDeploymentInstancesRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
//
// This method works, but is deprecated. Use BatchGetDeploymentTargets instead.
//
// Returns an array of one or more instances associated with a deployment. This
// method works with EC2/On-premises and AWS Lambda compute platforms. The newer
// BatchGetDeploymentTargets works with all compute platforms. The maximum number
// of instances that can be returned is 25.
//
//    // Example sending a request using BatchGetDeploymentInstancesRequest.
//    req := client.BatchGetDeploymentInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetDeploymentInstances
func (c *Client) BatchGetDeploymentInstancesRequest(input *BatchGetDeploymentInstancesInput) BatchGetDeploymentInstancesRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, BatchGetDeploymentInstances, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opBatchGetDeploymentInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetDeploymentInstancesInput{}
	}

	req := c.newRequest(op, input, &BatchGetDeploymentInstancesOutput{})
	return BatchGetDeploymentInstancesRequest{Request: req, Input: input, Copy: c.BatchGetDeploymentInstancesRequest}
}

// BatchGetDeploymentInstancesRequest is the request type for the
// BatchGetDeploymentInstances API operation.
type BatchGetDeploymentInstancesRequest struct {
	*aws.Request
	Input *BatchGetDeploymentInstancesInput
	Copy  func(*BatchGetDeploymentInstancesInput) BatchGetDeploymentInstancesRequest
}

// Send marshals and sends the BatchGetDeploymentInstances API request.
func (r BatchGetDeploymentInstancesRequest) Send(ctx context.Context) (*BatchGetDeploymentInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetDeploymentInstancesResponse{
		BatchGetDeploymentInstancesOutput: r.Request.Data.(*BatchGetDeploymentInstancesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetDeploymentInstancesResponse is the response type for the
// BatchGetDeploymentInstances API operation.
type BatchGetDeploymentInstancesResponse struct {
	*BatchGetDeploymentInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetDeploymentInstances request.
func (r *BatchGetDeploymentInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}