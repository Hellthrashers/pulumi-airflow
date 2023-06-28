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
    'RoleActionArgs',
]

@pulumi.input_type
class RoleActionArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 resource: pulumi.Input[str]):
        """
        :param pulumi.Input[str] action: The name of the permission.
        :param pulumi.Input[str] resource: The name of the resource.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "resource", resource)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        The name of the permission.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource", value)

