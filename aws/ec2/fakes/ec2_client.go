package fakes

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Client struct {
	DescribeVolumesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeVolumesInput
		}
		Returns struct {
			Output *ec2.DescribeVolumesOutput
			Error  error
		}
	}

	DeleteVolumeCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteVolumeInput
		}
		Returns struct {
			Output *ec2.DeleteVolumeOutput
			Error  error
		}
	}

	DescribeTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeTagsInput
		}
		Returns struct {
			Output *ec2.DescribeTagsOutput
			Error  error
		}
	}

	DeleteTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteTagsInput
		}
		Returns struct {
			Output *ec2.DeleteTagsOutput
			Error  error
		}
	}
}

func (e *EC2Client) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	e.DescribeVolumesCall.CallCount++
	e.DescribeVolumesCall.Receives.Input = input

	return e.DescribeVolumesCall.Returns.Output, e.DescribeVolumesCall.Returns.Error
}

func (e *EC2Client) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	e.DeleteVolumeCall.CallCount++
	e.DeleteVolumeCall.Receives.Input = input

	return e.DeleteVolumeCall.Returns.Output, e.DeleteVolumeCall.Returns.Error
}

func (e *EC2Client) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	e.DescribeTagsCall.CallCount++
	e.DescribeTagsCall.Receives.Input = input

	return e.DescribeTagsCall.Returns.Output, e.DescribeTagsCall.Returns.Error
}

func (e *EC2Client) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	e.DeleteTagsCall.CallCount++
	e.DeleteTagsCall.Receives.Input = input

	return e.DeleteTagsCall.Returns.Output, e.DeleteTagsCall.Returns.Error
}