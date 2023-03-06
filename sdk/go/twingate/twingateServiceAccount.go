// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TwingateServiceAccount struct {
	pulumi.CustomResourceState

	// The name of the Service Account in Twingate
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTwingateServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewTwingateServiceAccount(ctx *pulumi.Context,
	name string, args *TwingateServiceAccountArgs, opts ...pulumi.ResourceOption) (*TwingateServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TwingateServiceAccount
	err := ctx.RegisterResource("twingate:index/twingateServiceAccount:TwingateServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTwingateServiceAccount gets an existing TwingateServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTwingateServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TwingateServiceAccountState, opts ...pulumi.ResourceOption) (*TwingateServiceAccount, error) {
	var resource TwingateServiceAccount
	err := ctx.ReadResource("twingate:index/twingateServiceAccount:TwingateServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TwingateServiceAccount resources.
type twingateServiceAccountState struct {
	// The name of the Service Account in Twingate
	Name *string `pulumi:"name"`
}

type TwingateServiceAccountState struct {
	// The name of the Service Account in Twingate
	Name pulumi.StringPtrInput
}

func (TwingateServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateServiceAccountState)(nil)).Elem()
}

type twingateServiceAccountArgs struct {
	// The name of the Service Account in Twingate
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a TwingateServiceAccount resource.
type TwingateServiceAccountArgs struct {
	// The name of the Service Account in Twingate
	Name pulumi.StringInput
}

func (TwingateServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateServiceAccountArgs)(nil)).Elem()
}

type TwingateServiceAccountInput interface {
	pulumi.Input

	ToTwingateServiceAccountOutput() TwingateServiceAccountOutput
	ToTwingateServiceAccountOutputWithContext(ctx context.Context) TwingateServiceAccountOutput
}

func (*TwingateServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateServiceAccount)(nil)).Elem()
}

func (i *TwingateServiceAccount) ToTwingateServiceAccountOutput() TwingateServiceAccountOutput {
	return i.ToTwingateServiceAccountOutputWithContext(context.Background())
}

func (i *TwingateServiceAccount) ToTwingateServiceAccountOutputWithContext(ctx context.Context) TwingateServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateServiceAccountOutput)
}

// TwingateServiceAccountArrayInput is an input type that accepts TwingateServiceAccountArray and TwingateServiceAccountArrayOutput values.
// You can construct a concrete instance of `TwingateServiceAccountArrayInput` via:
//
//	TwingateServiceAccountArray{ TwingateServiceAccountArgs{...} }
type TwingateServiceAccountArrayInput interface {
	pulumi.Input

	ToTwingateServiceAccountArrayOutput() TwingateServiceAccountArrayOutput
	ToTwingateServiceAccountArrayOutputWithContext(context.Context) TwingateServiceAccountArrayOutput
}

type TwingateServiceAccountArray []TwingateServiceAccountInput

func (TwingateServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateServiceAccount)(nil)).Elem()
}

func (i TwingateServiceAccountArray) ToTwingateServiceAccountArrayOutput() TwingateServiceAccountArrayOutput {
	return i.ToTwingateServiceAccountArrayOutputWithContext(context.Background())
}

func (i TwingateServiceAccountArray) ToTwingateServiceAccountArrayOutputWithContext(ctx context.Context) TwingateServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateServiceAccountArrayOutput)
}

// TwingateServiceAccountMapInput is an input type that accepts TwingateServiceAccountMap and TwingateServiceAccountMapOutput values.
// You can construct a concrete instance of `TwingateServiceAccountMapInput` via:
//
//	TwingateServiceAccountMap{ "key": TwingateServiceAccountArgs{...} }
type TwingateServiceAccountMapInput interface {
	pulumi.Input

	ToTwingateServiceAccountMapOutput() TwingateServiceAccountMapOutput
	ToTwingateServiceAccountMapOutputWithContext(context.Context) TwingateServiceAccountMapOutput
}

type TwingateServiceAccountMap map[string]TwingateServiceAccountInput

func (TwingateServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateServiceAccount)(nil)).Elem()
}

func (i TwingateServiceAccountMap) ToTwingateServiceAccountMapOutput() TwingateServiceAccountMapOutput {
	return i.ToTwingateServiceAccountMapOutputWithContext(context.Background())
}

func (i TwingateServiceAccountMap) ToTwingateServiceAccountMapOutputWithContext(ctx context.Context) TwingateServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateServiceAccountMapOutput)
}

type TwingateServiceAccountOutput struct{ *pulumi.OutputState }

func (TwingateServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateServiceAccount)(nil)).Elem()
}

func (o TwingateServiceAccountOutput) ToTwingateServiceAccountOutput() TwingateServiceAccountOutput {
	return o
}

func (o TwingateServiceAccountOutput) ToTwingateServiceAccountOutputWithContext(ctx context.Context) TwingateServiceAccountOutput {
	return o
}

// The name of the Service Account in Twingate
func (o TwingateServiceAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateServiceAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TwingateServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (TwingateServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateServiceAccount)(nil)).Elem()
}

func (o TwingateServiceAccountArrayOutput) ToTwingateServiceAccountArrayOutput() TwingateServiceAccountArrayOutput {
	return o
}

func (o TwingateServiceAccountArrayOutput) ToTwingateServiceAccountArrayOutputWithContext(ctx context.Context) TwingateServiceAccountArrayOutput {
	return o
}

func (o TwingateServiceAccountArrayOutput) Index(i pulumi.IntInput) TwingateServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TwingateServiceAccount {
		return vs[0].([]*TwingateServiceAccount)[vs[1].(int)]
	}).(TwingateServiceAccountOutput)
}

type TwingateServiceAccountMapOutput struct{ *pulumi.OutputState }

func (TwingateServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateServiceAccount)(nil)).Elem()
}

func (o TwingateServiceAccountMapOutput) ToTwingateServiceAccountMapOutput() TwingateServiceAccountMapOutput {
	return o
}

func (o TwingateServiceAccountMapOutput) ToTwingateServiceAccountMapOutputWithContext(ctx context.Context) TwingateServiceAccountMapOutput {
	return o
}

func (o TwingateServiceAccountMapOutput) MapIndex(k pulumi.StringInput) TwingateServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TwingateServiceAccount {
		return vs[0].(map[string]*TwingateServiceAccount)[vs[1].(string)]
	}).(TwingateServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateServiceAccountInput)(nil)).Elem(), &TwingateServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateServiceAccountArrayInput)(nil)).Elem(), TwingateServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateServiceAccountMapInput)(nil)).Elem(), TwingateServiceAccountMap{})
	pulumi.RegisterOutputType(TwingateServiceAccountOutput{})
	pulumi.RegisterOutputType(TwingateServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(TwingateServiceAccountMapOutput{})
}
