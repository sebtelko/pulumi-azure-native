# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'ManagedServiceIdentityResponse',
    'SystemDataResponse',
    'UserAssignedIdentityResponse',
]

@pulumi.output_type
class ManagedServiceIdentityResponse(dict):
    """
    Managed identity generic object.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "tenantId":
            suggest = "tenant_id"
        elif key == "userAssignedIdentities":
            suggest = "user_assigned_identities"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedServiceIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedServiceIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedServiceIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 tenant_id: str,
                 type: Optional[str] = None,
                 user_assigned_identities: Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']] = None):
        """
        Managed identity generic object.
        :param str tenant_id: ID of the Azure Active Directory.
        :param str type: Type of the managed identity.
        :param Mapping[str, 'UserAssignedIdentityResponse'] user_assigned_identities: The list of user-assigned managed identities associated with the resource. Key is the Azure resource Id of the managed identity.
        """
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        ID of the Azure Active Directory.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Type of the managed identity.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']]:
        """
        The list of user-assigned managed identities associated with the resource. Key is the Azure resource Id of the managed identity.
        """
        return pulumi.get(self, "user_assigned_identities")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


@pulumi.output_type
class UserAssignedIdentityResponse(dict):
    """
    User-assigned managed identity.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientId":
            suggest = "client_id"
        elif key == "principalId":
            suggest = "principal_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in UserAssignedIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        UserAssignedIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        UserAssignedIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_id: str,
                 principal_id: str):
        """
        User-assigned managed identity.
        :param str client_id: Client App Id associated with this identity.
        :param str principal_id: Azure Active Directory principal ID associated with this identity.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "principal_id", principal_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        Client App Id associated with this identity.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        Azure Active Directory principal ID associated with this identity.
        """
        return pulumi.get(self, "principal_id")


