// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteQueryDefinitionInput struct {
	_ struct{} `type:"structure"`

	// QueryDefinitionId is a required field
	QueryDefinitionId *string `locationName:"queryDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteQueryDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteQueryDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteQueryDefinitionInput"}

	if s.QueryDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteQueryDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Success *bool `locationName:"success" type:"boolean"`
}

// String returns the string representation
func (s DeleteQueryDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteQueryDefinition = "DeleteQueryDefinition"

// DeleteQueryDefinitionRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
//    // Example sending a request using DeleteQueryDefinitionRequest.
//    req := client.DeleteQueryDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteQueryDefinition
func (c *Client) DeleteQueryDefinitionRequest(input *DeleteQueryDefinitionInput) DeleteQueryDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteQueryDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteQueryDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteQueryDefinitionOutput{})

	return DeleteQueryDefinitionRequest{Request: req, Input: input, Copy: c.DeleteQueryDefinitionRequest}
}

// DeleteQueryDefinitionRequest is the request type for the
// DeleteQueryDefinition API operation.
type DeleteQueryDefinitionRequest struct {
	*aws.Request
	Input *DeleteQueryDefinitionInput
	Copy  func(*DeleteQueryDefinitionInput) DeleteQueryDefinitionRequest
}

// Send marshals and sends the DeleteQueryDefinition API request.
func (r DeleteQueryDefinitionRequest) Send(ctx context.Context) (*DeleteQueryDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteQueryDefinitionResponse{
		DeleteQueryDefinitionOutput: r.Request.Data.(*DeleteQueryDefinitionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteQueryDefinitionResponse is the response type for the
// DeleteQueryDefinition API operation.
type DeleteQueryDefinitionResponse struct {
	*DeleteQueryDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteQueryDefinition request.
func (r *DeleteQueryDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}