# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetTwingateGroupResult',
    'AwaitableGetTwingateGroupResult',
    'get_twingate_group',
    'get_twingate_group_output',
]

@pulumi.output_type
class GetTwingateGroupResult:
    """
    A collection of values returned by getTwingateGroup.
    """
    def __init__(__self__, id=None, is_active=None, name=None, security_policy_id=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if security_policy_id and not isinstance(security_policy_id, str):
            raise TypeError("Expected argument 'security_policy_id' to be a str")
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> str:
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetTwingateGroupResult(GetTwingateGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTwingateGroupResult(
            id=self.id,
            is_active=self.is_active,
            name=self.name,
            security_policy_id=self.security_policy_id,
            type=self.type)


def get_twingate_group(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTwingateGroupResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('twingate:index/getTwingateGroup:getTwingateGroup', __args__, opts=opts, typ=GetTwingateGroupResult).value

    return AwaitableGetTwingateGroupResult(
        id=__ret__.id,
        is_active=__ret__.is_active,
        name=__ret__.name,
        security_policy_id=__ret__.security_policy_id,
        type=__ret__.type)


@_utilities.lift_output_func(get_twingate_group)
def get_twingate_group_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTwingateGroupResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
