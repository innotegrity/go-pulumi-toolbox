package gcp

import (
	"fmt"
	"strings"

	googlecompute "github.com/pulumi/pulumi-google-native/sdk/go/google/compute/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// Subnetwork is used for creating a subnet resource.
type Subnetwork struct {
	ID                      string               `json:"id"`
	Description             *string              `json:"description"`
	EnableFlowLogs          *bool                `json:"enableFlowLogs"`
	Region                  *string              `json:"region"`
	ExternalIPv6Prefix      *string              `json:"externalIPv6Prefix"`
	IPRange                 *string              `json:"ipRange"`
	IPv6AccessType          *string              `json:"ipv6AccessType"`
	LogConfig               *subnetworkLogConfig `json:"logConfig"`
	Name                    *string              `json:"name"`
	Network                 *string              `json:"network"`
	PrivateIPGoogleAccess   *bool                `json:"privateIPGoogleAccess"`
	PrivateIPv6GoogleAccess *string              `json:"privateIPv6GoogleAccess"`
	Project                 *string              `json:"project"`
	Purpose                 *string              `json:"purpose"`
	RequestID               *string              `json:"requestID"`
	Role                    *string              `json:"role"`
	SecondaryIPRanges       []subnetworkIPRange  `json:"secondaryIPRanges"`
	StackType               *string              `json:"ipStackType"`
}

// subnetworkLogConfig holds the logging configuration for a subnet.
type subnetworkLogConfig struct {
	AggregationInterval *string  `json:"aggregationInterval"`
	Enable              *bool    `json:"enable"`
	FilterExpr          *string  `json:"filterExpr"`
	FlowSampling        *float64 `json:"flowSampling"`
	Metadata            *string  `json:"metadata"`
	MetadataFields      []string `json:"metadataFields"`
}

// subnetworkIPRange contains information about a secondary IP range for a subnet.
type subnetworkIPRange struct {
	IPRange *string `json:"ipRange"`
	Name    *string `json:"name"`
}

// NewSubnetwork returns a new, empty object.
func NewSubnetwork() config.ResourceTypeProvisioner {
	return &Subnetwork{}
}

// GetID returns the ID of the object.
func (s Subnetwork) GetID() string {
	return s.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (s Subnetwork) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*googlecompute.Network)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", s.GetType())
	}

	// return the output for the given property
	switch property {
	case "ID":
		return resource.ID(), nil
	default:
		return nil, fmt.Errorf("'%s': unknown resource property", property)
	}
}

// GetType returns the type of the object.
func (s Subnetwork) GetType() string {
	return SUBNETWORK_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (s Subnetwork) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, s.GetType(), s.GetID())
	stackState := state.Get()

	// check required arguments
	if s.Region == nil || *s.Region == "" {
		return nil, fmt.Errorf("'region' is a required setting")
	}
	if s.Network == nil || *s.Network == "" {
		return nil, fmt.Errorf("'network' is a required setting")
	}

	// initialize args
	args := googlecompute.SubnetworkArgs{
		Region: pulumi.String(*s.Region),
	}
	if s.Name != nil {
		args.Name = pulumi.String(*s.Name)
	}
	if s.Description != nil {
		args.Description = pulumi.String(*s.Description)
	}
	if s.EnableFlowLogs != nil {
		args.EnableFlowLogs = pulumi.Bool(*s.EnableFlowLogs)
	}
	if s.ExternalIPv6Prefix != nil {
		args.ExternalIpv6Prefix = pulumi.String(*s.ExternalIPv6Prefix)
	}
	if s.IPRange != nil {
		args.IpCidrRange = pulumi.String(*s.IPRange)
	}
	if s.IPv6AccessType != nil {
		args.Ipv6AccessType = googlecompute.SubnetworkIpv6AccessType(strings.ToUpper(*s.IPv6AccessType))
	}
	if s.PrivateIPGoogleAccess != nil {
		args.PrivateIpGoogleAccess = pulumi.Bool(*s.PrivateIPGoogleAccess)
	}
	if s.PrivateIPv6GoogleAccess != nil {
		args.PrivateIpv6GoogleAccess = googlecompute.SubnetworkPrivateIpv6GoogleAccess(
			strings.ToUpper(*s.PrivateIPv6GoogleAccess))
	}
	if s.Project != nil {
		args.Project = pulumi.String(*s.Project)
	}
	if s.Purpose != nil {
		args.Purpose = googlecompute.SubnetworkPurpose(strings.ToUpper(*s.Purpose))
	}
	if s.RequestID != nil {
		args.RequestId = pulumi.String(*s.RequestID)
	}
	if s.Role != nil {
		args.Role = googlecompute.SubnetworkRole(strings.ToUpper(*s.Role))
	}
	if s.StackType != nil {
		args.StackType = googlecompute.SubnetworkStackType(strings.ToUpper(*s.StackType))
	}

	// lookup network ID
	id, err := stackState.GetOutput(ctx, *s.Network)
	if err != nil {
		return nil, err
	}
	args.Network = id.ApplyT(func(id string) string {
		ctx.Log.Debug(fmt.Sprintf("'%s': network ID retrieved", *s.Network), nil)
		return id
	}).(pulumi.StringOutput)

	// setup logging configuration
	if s.LogConfig != nil {
		config := googlecompute.SubnetworkLogConfigArgs{}
		if s.LogConfig.AggregationInterval != nil {
			config.AggregationInterval = googlecompute.SubnetworkLogConfigAggregationInterval(
				strings.ToUpper(*s.LogConfig.AggregationInterval))
		}
		if s.LogConfig.Enable != nil {
			config.Enable = pulumi.Bool(*s.LogConfig.Enable)
		}
		if s.LogConfig.FilterExpr != nil {
			config.FilterExpr = pulumi.String(*s.LogConfig.FilterExpr)
		}
		if s.LogConfig.FlowSampling != nil {
			config.FlowSampling = pulumi.Float64(*s.LogConfig.FlowSampling)
		}
		if s.LogConfig.Metadata != nil {
			config.Metadata = googlecompute.SubnetworkLogConfigMetadata(strings.ToUpper(*s.LogConfig.Metadata))
		}
		if len(s.LogConfig.MetadataFields) > 0 {
			config.MetadataFields = pulumi.ToStringArray(s.LogConfig.MetadataFields)
		}
		args.LogConfig = config
	}

	// setup secondary IP ranges
	if len(s.SecondaryIPRanges) > 0 {
		ranges := googlecompute.SubnetworkSecondaryRangeArray{}
		for _, r := range s.SecondaryIPRanges {
			config := googlecompute.SubnetworkSecondaryRangeArgs{}
			if r.IPRange != nil {
				config.IpCidrRange = pulumi.String(*r.IPRange)
			}
			if r.Name != nil {
				config.RangeName = pulumi.String(*r.Name)
			}
			ranges = append(ranges, config)
		}
		args.SecondaryIpRanges = ranges
	}

	// provision the resource
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := googlecompute.NewSubnetwork(ctx, s.ID, &args, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
