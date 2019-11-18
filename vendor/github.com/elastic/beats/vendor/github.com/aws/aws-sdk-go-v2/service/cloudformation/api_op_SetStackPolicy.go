// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// The input for the SetStackPolicy action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/SetStackPolicyInput
type SetStackPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name or unique stack ID that you want to associate a policy with.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`

	// Structure containing the stack policy body. For more information, go to Prevent
	// Updates to Stack Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html)
	// in the AWS CloudFormation User Guide. You can specify either the StackPolicyBody
	// or the StackPolicyURL parameter, but not both.
	StackPolicyBody *string `min:"1" type:"string"`

	// Location of a file containing the stack policy. The URL must point to a policy
	// (maximum size: 16 KB) located in an S3 bucket in the same region as the stack.
	// You can specify either the StackPolicyBody or the StackPolicyURL parameter,
	// but not both.
	StackPolicyURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s SetStackPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetStackPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetStackPolicyInput"}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackPolicyBody != nil && len(*s.StackPolicyBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackPolicyBody", 1))
	}
	if s.StackPolicyURL != nil && len(*s.StackPolicyURL) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackPolicyURL", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/SetStackPolicyOutput
type SetStackPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetStackPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetStackPolicy = "SetStackPolicy"

// SetStackPolicyRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Sets a stack policy for a specified stack.
//
//    // Example sending a request using SetStackPolicyRequest.
//    req := client.SetStackPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/SetStackPolicy
func (c *Client) SetStackPolicyRequest(input *SetStackPolicyInput) SetStackPolicyRequest {
	op := &aws.Operation{
		Name:       opSetStackPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetStackPolicyInput{}
	}

	req := c.newRequest(op, input, &SetStackPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetStackPolicyRequest{Request: req, Input: input, Copy: c.SetStackPolicyRequest}
}

// SetStackPolicyRequest is the request type for the
// SetStackPolicy API operation.
type SetStackPolicyRequest struct {
	*aws.Request
	Input *SetStackPolicyInput
	Copy  func(*SetStackPolicyInput) SetStackPolicyRequest
}

// Send marshals and sends the SetStackPolicy API request.
func (r SetStackPolicyRequest) Send(ctx context.Context) (*SetStackPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetStackPolicyResponse{
		SetStackPolicyOutput: r.Request.Data.(*SetStackPolicyOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetStackPolicyResponse is the response type for the
// SetStackPolicy API operation.
type SetStackPolicyResponse struct {
	*SetStackPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetStackPolicy request.
func (r *SetStackPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
