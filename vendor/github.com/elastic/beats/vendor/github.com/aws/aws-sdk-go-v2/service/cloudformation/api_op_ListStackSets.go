// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSetsInput
type ListStackSetsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListStackSets again and assign that token to
	// the request object's NextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string `min:"1" type:"string"`

	// The status of the stack sets that you want to get summary information about.
	Status StackSetStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListStackSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStackSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStackSetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSetsOutput
type ListStackSetsOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all of the remaining results, NextToken is
	// set to a token. To retrieve the next set of results, call ListStackInstances
	// again and assign that token to the request object's NextToken parameter.
	// If the request returns all results, NextToken is set to null.
	NextToken *string `min:"1" type:"string"`

	// A list of StackSetSummary structures that contain information about the user's
	// stack sets.
	Summaries []StackSetSummary `type:"list"`
}

// String returns the string representation
func (s ListStackSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStackSets = "ListStackSets"

// ListStackSetsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns summary information about stack sets that are associated with the
// user.
//
//    // Example sending a request using ListStackSetsRequest.
//    req := client.ListStackSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSets
func (c *Client) ListStackSetsRequest(input *ListStackSetsInput) ListStackSetsRequest {
	op := &aws.Operation{
		Name:       opListStackSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListStackSetsInput{}
	}

	req := c.newRequest(op, input, &ListStackSetsOutput{})
	return ListStackSetsRequest{Request: req, Input: input, Copy: c.ListStackSetsRequest}
}

// ListStackSetsRequest is the request type for the
// ListStackSets API operation.
type ListStackSetsRequest struct {
	*aws.Request
	Input *ListStackSetsInput
	Copy  func(*ListStackSetsInput) ListStackSetsRequest
}

// Send marshals and sends the ListStackSets API request.
func (r ListStackSetsRequest) Send(ctx context.Context) (*ListStackSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStackSetsResponse{
		ListStackSetsOutput: r.Request.Data.(*ListStackSetsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListStackSetsResponse is the response type for the
// ListStackSets API operation.
type ListStackSetsResponse struct {
	*ListStackSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStackSets request.
func (r *ListStackSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}