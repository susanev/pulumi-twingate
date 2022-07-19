// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TwingateLabs.Twingate
{
    public static class GetTwingateGroup
    {
        public static Task<GetTwingateGroupResult> InvokeAsync(GetTwingateGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTwingateGroupResult>("twingate:index/getTwingateGroup:getTwingateGroup", args ?? new GetTwingateGroupArgs(), options.WithDefaults());

        public static Output<GetTwingateGroupResult> Invoke(GetTwingateGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTwingateGroupResult>("twingate:index/getTwingateGroup:getTwingateGroup", args ?? new GetTwingateGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTwingateGroupArgs()
        {
        }
    }

    public sealed class GetTwingateGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTwingateGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTwingateGroupResult
    {
        public readonly string Id;
        public readonly bool IsActive;
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetTwingateGroupResult(
            string id,

            bool isActive,

            string name,

            string type)
        {
            Id = id;
            IsActive = isActive;
            Name = name;
            Type = type;
        }
    }
}
