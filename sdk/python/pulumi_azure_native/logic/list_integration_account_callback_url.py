# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ListIntegrationAccountCallbackUrlResult',
    'AwaitableListIntegrationAccountCallbackUrlResult',
    'list_integration_account_callback_url',
]

@pulumi.output_type
class ListIntegrationAccountCallbackUrlResult:
    """
    The callback url.
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The URL value.
        """
        return pulumi.get(self, "value")


class AwaitableListIntegrationAccountCallbackUrlResult(ListIntegrationAccountCallbackUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIntegrationAccountCallbackUrlResult(
            value=self.value)


def list_integration_account_callback_url(integration_account_name: Optional[str] = None,
                                          key_type: Optional[Union[str, 'KeyType']] = None,
                                          not_after: Optional[str] = None,
                                          resource_group_name: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationAccountCallbackUrlResult:
    """
    The callback url.
    API Version: 2019-05-01.


    :param str integration_account_name: The integration account name.
    :param Union[str, 'KeyType'] key_type: The key type.
    :param str not_after: The expiry time.
    :param str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['keyType'] = key_type
    __args__['notAfter'] = not_after
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:logic:listIntegrationAccountCallbackUrl', __args__, opts=opts, typ=ListIntegrationAccountCallbackUrlResult).value

    return AwaitableListIntegrationAccountCallbackUrlResult(
        value=__ret__.value)
