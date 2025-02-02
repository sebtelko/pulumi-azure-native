# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult',
    'AwaitableListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult',
    'list_billing_account_invoice_sections_by_create_subscription_permission',
]

@pulumi.output_type
class ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult:
    """
    The list of invoice section properties with create subscription permission.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> str:
        """
        The link (url) to the next page of results.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.InvoiceSectionWithCreateSubPermissionResponse']]:
        """
        The list of invoice section properties with create subscription permission.
        """
        return pulumi.get(self, "value")


class AwaitableListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult(
            next_link=self.next_link,
            value=self.value)


def list_billing_account_invoice_sections_by_create_subscription_permission(billing_account_name: Optional[str] = None,
                                                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult:
    """
    The list of invoice section properties with create subscription permission.


    :param str billing_account_name: The ID that uniquely identifies a billing account.
    """
    __args__ = dict()
    __args__['billingAccountName'] = billing_account_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:billing/v20200501:listBillingAccountInvoiceSectionsByCreateSubscriptionPermission', __args__, opts=opts, typ=ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult).value

    return AwaitableListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
