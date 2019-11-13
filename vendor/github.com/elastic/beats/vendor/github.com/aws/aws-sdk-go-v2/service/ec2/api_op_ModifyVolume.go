// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVolumeRequest
type ModifyVolumeInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The target IOPS rate of the volume.
	//
	// This is only valid for Provisioned IOPS SSD (io1) volumes. For more information,
	// see Provisioned IOPS SSD (io1) Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html#EBSVolumeTypes_piops).
	//
	// Default: If no IOPS value is specified, the existing value is retained.
	Iops *int64 `type:"integer"`

	// The target size of the volume, in GiB. The target volume size must be greater
	// than or equal to than the existing size of the volume. For information about
	// available EBS volume sizes, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html).
	//
	// Default: If no size is specified, the existing size is retained.
	Size *int64 `type:"integer"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`

	// The target EBS volume type of the volume.
	//
	// Default: If no type is specified, the existing type is retained.
	VolumeType VolumeType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ModifyVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVolumeResult
type ModifyVolumeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the volume modification.
	VolumeModification *VolumeModification `locationName:"volumeModification" type:"structure"`
}

// String returns the string representation
func (s ModifyVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVolume = "ModifyVolume"

// ModifyVolumeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// You can modify several parameters of an existing EBS volume, including volume
// size, volume type, and IOPS capacity. If your EBS volume is attached to a
// current-generation EC2 instance type, you may be able to apply these changes
// without stopping the instance or detaching the volume from it. For more information
// about modifying an EBS volume running Linux, see Modifying the Size, IOPS,
// or Type of an EBS Volume on Linux (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html).
// For more information about modifying an EBS volume running Windows, see Modifying
// the Size, IOPS, or Type of an EBS Volume on Windows (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html).
//
// When you complete a resize operation on your volume, you need to extend the
// volume's file-system size to take advantage of the new storage capacity.
// For information about extending a Linux file system, see Extending a Linux
// File System (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#recognize-expanded-volume-linux).
// For information about extending a Windows file system, see Extending a Windows
// File System (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html#recognize-expanded-volume-windows).
//
// You can use CloudWatch Events to check the status of a modification to an
// EBS volume. For information about CloudWatch Events, see the Amazon CloudWatch
// Events User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/).
// You can also track the status of a modification using the DescribeVolumesModifications
// API. For information about tracking status changes using either method, see
// Monitoring Volume Modifications (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#monitoring_mods).
//
// With previous-generation instance types, resizing an EBS volume may require
// detaching and reattaching the volume or stopping and restarting the instance.
// For more information, see Modifying the Size, IOPS, or Type of an EBS Volume
// on Linux (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html)
// and Modifying the Size, IOPS, or Type of an EBS Volume on Windows (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html).
//
// If you reach the maximum volume modification rate per volume limit, you will
// need to wait at least six hours before applying further modifications to
// the affected EBS volume.
//
//    // Example sending a request using ModifyVolumeRequest.
//    req := client.ModifyVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVolume
func (c *Client) ModifyVolumeRequest(input *ModifyVolumeInput) ModifyVolumeRequest {
	op := &aws.Operation{
		Name:       opModifyVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVolumeInput{}
	}

	req := c.newRequest(op, input, &ModifyVolumeOutput{})
	return ModifyVolumeRequest{Request: req, Input: input, Copy: c.ModifyVolumeRequest}
}

// ModifyVolumeRequest is the request type for the
// ModifyVolume API operation.
type ModifyVolumeRequest struct {
	*aws.Request
	Input *ModifyVolumeInput
	Copy  func(*ModifyVolumeInput) ModifyVolumeRequest
}

// Send marshals and sends the ModifyVolume API request.
func (r ModifyVolumeRequest) Send(ctx context.Context) (*ModifyVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVolumeResponse{
		ModifyVolumeOutput: r.Request.Data.(*ModifyVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVolumeResponse is the response type for the
// ModifyVolume API operation.
type ModifyVolumeResponse struct {
	*ModifyVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVolume request.
func (r *ModifyVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
