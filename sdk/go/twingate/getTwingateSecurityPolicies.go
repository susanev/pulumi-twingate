// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTwingateSecurityPolicies(ctx *pulumi.Context, args *GetTwingateSecurityPoliciesArgs, opts ...pulumi.InvokeOption) (*GetTwingateSecurityPoliciesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTwingateSecurityPoliciesResult
	err := ctx.Invoke("twingate:index/getTwingateSecurityPolicies:getTwingateSecurityPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateSecurityPolicies.
type GetTwingateSecurityPoliciesArgs struct {
	SecurityPolicies []GetTwingateSecurityPoliciesSecurityPolicy `pulumi:"securityPolicies"`
}

// A collection of values returned by getTwingateSecurityPolicies.
type GetTwingateSecurityPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                                      `pulumi:"id"`
	SecurityPolicies []GetTwingateSecurityPoliciesSecurityPolicy `pulumi:"securityPolicies"`
}

func GetTwingateSecurityPoliciesOutput(ctx *pulumi.Context, args GetTwingateSecurityPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetTwingateSecurityPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTwingateSecurityPoliciesResult, error) {
			args := v.(GetTwingateSecurityPoliciesArgs)
			r, err := GetTwingateSecurityPolicies(ctx, &args, opts...)
			var s GetTwingateSecurityPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTwingateSecurityPoliciesResultOutput)
}

// A collection of arguments for invoking getTwingateSecurityPolicies.
type GetTwingateSecurityPoliciesOutputArgs struct {
	SecurityPolicies GetTwingateSecurityPoliciesSecurityPolicyArrayInput `pulumi:"securityPolicies"`
}

func (GetTwingateSecurityPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateSecurityPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateSecurityPolicies.
type GetTwingateSecurityPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetTwingateSecurityPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateSecurityPoliciesResult)(nil)).Elem()
}

func (o GetTwingateSecurityPoliciesResultOutput) ToGetTwingateSecurityPoliciesResultOutput() GetTwingateSecurityPoliciesResultOutput {
	return o
}

func (o GetTwingateSecurityPoliciesResultOutput) ToGetTwingateSecurityPoliciesResultOutputWithContext(ctx context.Context) GetTwingateSecurityPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTwingateSecurityPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateSecurityPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTwingateSecurityPoliciesResultOutput) SecurityPolicies() GetTwingateSecurityPoliciesSecurityPolicyArrayOutput {
	return o.ApplyT(func(v GetTwingateSecurityPoliciesResult) []GetTwingateSecurityPoliciesSecurityPolicy {
		return v.SecurityPolicies
	}).(GetTwingateSecurityPoliciesSecurityPolicyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateSecurityPoliciesResultOutput{})
}
