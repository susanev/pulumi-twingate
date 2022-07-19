// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getTwingateConnectors(opts?: pulumi.InvokeOptions): Promise<GetTwingateConnectorsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("twingate:index/getTwingateConnectors:getTwingateConnectors", {
    }, opts);
}

/**
 * A collection of values returned by getTwingateConnectors.
 */
export interface GetTwingateConnectorsResult {
    readonly connectors: outputs.GetTwingateConnectorsConnector[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
