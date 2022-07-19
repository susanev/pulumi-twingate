// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TwingateGroup struct {
	pulumi.CustomResourceState

	// The name of the group
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTwingateGroup registers a new resource with the given unique name, arguments, and options.
func NewTwingateGroup(ctx *pulumi.Context,
	name string, args *TwingateGroupArgs, opts ...pulumi.ResourceOption) (*TwingateGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TwingateGroup
	err := ctx.RegisterResource("twingate:index/twingateGroup:TwingateGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTwingateGroup gets an existing TwingateGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTwingateGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TwingateGroupState, opts ...pulumi.ResourceOption) (*TwingateGroup, error) {
	var resource TwingateGroup
	err := ctx.ReadResource("twingate:index/twingateGroup:TwingateGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TwingateGroup resources.
type twingateGroupState struct {
	// The name of the group
	Name *string `pulumi:"name"`
}

type TwingateGroupState struct {
	// The name of the group
	Name pulumi.StringPtrInput
}

func (TwingateGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateGroupState)(nil)).Elem()
}

type twingateGroupArgs struct {
	// The name of the group
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a TwingateGroup resource.
type TwingateGroupArgs struct {
	// The name of the group
	Name pulumi.StringInput
}

func (TwingateGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateGroupArgs)(nil)).Elem()
}

type TwingateGroupInput interface {
	pulumi.Input

	ToTwingateGroupOutput() TwingateGroupOutput
	ToTwingateGroupOutputWithContext(ctx context.Context) TwingateGroupOutput
}

func (*TwingateGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateGroup)(nil)).Elem()
}

func (i *TwingateGroup) ToTwingateGroupOutput() TwingateGroupOutput {
	return i.ToTwingateGroupOutputWithContext(context.Background())
}

func (i *TwingateGroup) ToTwingateGroupOutputWithContext(ctx context.Context) TwingateGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateGroupOutput)
}

// TwingateGroupArrayInput is an input type that accepts TwingateGroupArray and TwingateGroupArrayOutput values.
// You can construct a concrete instance of `TwingateGroupArrayInput` via:
//
//          TwingateGroupArray{ TwingateGroupArgs{...} }
type TwingateGroupArrayInput interface {
	pulumi.Input

	ToTwingateGroupArrayOutput() TwingateGroupArrayOutput
	ToTwingateGroupArrayOutputWithContext(context.Context) TwingateGroupArrayOutput
}

type TwingateGroupArray []TwingateGroupInput

func (TwingateGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateGroup)(nil)).Elem()
}

func (i TwingateGroupArray) ToTwingateGroupArrayOutput() TwingateGroupArrayOutput {
	return i.ToTwingateGroupArrayOutputWithContext(context.Background())
}

func (i TwingateGroupArray) ToTwingateGroupArrayOutputWithContext(ctx context.Context) TwingateGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateGroupArrayOutput)
}

// TwingateGroupMapInput is an input type that accepts TwingateGroupMap and TwingateGroupMapOutput values.
// You can construct a concrete instance of `TwingateGroupMapInput` via:
//
//          TwingateGroupMap{ "key": TwingateGroupArgs{...} }
type TwingateGroupMapInput interface {
	pulumi.Input

	ToTwingateGroupMapOutput() TwingateGroupMapOutput
	ToTwingateGroupMapOutputWithContext(context.Context) TwingateGroupMapOutput
}

type TwingateGroupMap map[string]TwingateGroupInput

func (TwingateGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateGroup)(nil)).Elem()
}

func (i TwingateGroupMap) ToTwingateGroupMapOutput() TwingateGroupMapOutput {
	return i.ToTwingateGroupMapOutputWithContext(context.Background())
}

func (i TwingateGroupMap) ToTwingateGroupMapOutputWithContext(ctx context.Context) TwingateGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateGroupMapOutput)
}

type TwingateGroupOutput struct{ *pulumi.OutputState }

func (TwingateGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateGroup)(nil)).Elem()
}

func (o TwingateGroupOutput) ToTwingateGroupOutput() TwingateGroupOutput {
	return o
}

func (o TwingateGroupOutput) ToTwingateGroupOutputWithContext(ctx context.Context) TwingateGroupOutput {
	return o
}

// The name of the group
func (o TwingateGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TwingateGroupArrayOutput struct{ *pulumi.OutputState }

func (TwingateGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateGroup)(nil)).Elem()
}

func (o TwingateGroupArrayOutput) ToTwingateGroupArrayOutput() TwingateGroupArrayOutput {
	return o
}

func (o TwingateGroupArrayOutput) ToTwingateGroupArrayOutputWithContext(ctx context.Context) TwingateGroupArrayOutput {
	return o
}

func (o TwingateGroupArrayOutput) Index(i pulumi.IntInput) TwingateGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TwingateGroup {
		return vs[0].([]*TwingateGroup)[vs[1].(int)]
	}).(TwingateGroupOutput)
}

type TwingateGroupMapOutput struct{ *pulumi.OutputState }

func (TwingateGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateGroup)(nil)).Elem()
}

func (o TwingateGroupMapOutput) ToTwingateGroupMapOutput() TwingateGroupMapOutput {
	return o
}

func (o TwingateGroupMapOutput) ToTwingateGroupMapOutputWithContext(ctx context.Context) TwingateGroupMapOutput {
	return o
}

func (o TwingateGroupMapOutput) MapIndex(k pulumi.StringInput) TwingateGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TwingateGroup {
		return vs[0].(map[string]*TwingateGroup)[vs[1].(string)]
	}).(TwingateGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateGroupInput)(nil)).Elem(), &TwingateGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateGroupArrayInput)(nil)).Elem(), TwingateGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateGroupMapInput)(nil)).Elem(), TwingateGroupMap{})
	pulumi.RegisterOutputType(TwingateGroupOutput{})
	pulumi.RegisterOutputType(TwingateGroupArrayOutput{})
	pulumi.RegisterOutputType(TwingateGroupMapOutput{})
}
