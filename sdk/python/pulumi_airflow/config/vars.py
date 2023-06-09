# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('airflow')


class _ExportableConfig(types.ModuleType):
    @property
    def base_endpoint(self) -> Optional[str]:
        return __config__.get('baseEndpoint') or _utilities.get_env('AIRFLOW_BASE_ENDPOINT')

    @property
    def disable_ssl_verification(self) -> Optional[bool]:
        """
        Disable SSL verification
        """
        return __config__.get_bool('disableSslVerification')

    @property
    def oauth2_token(self) -> Optional[str]:
        """
        The oauth to use for API authentication
        """
        return __config__.get('oauth2Token')

    @property
    def password(self) -> Optional[str]:
        """
        The password to use for API basic authentication
        """
        return __config__.get('password')

    @property
    def username(self) -> Optional[str]:
        """
        The username to use for API basic authentication
        """
        return __config__.get('username')

