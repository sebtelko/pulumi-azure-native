# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AccessTier',
    'Action',
    'BlobAccessTier',
    'BlobType',
    'Bypass',
    'DefaultAction',
    'DirectoryServiceOptions',
    'EnabledProtocols',
    'EncryptionScopeSource',
    'EncryptionScopeState',
    'ExpirationAction',
    'ExtendedLocationTypes',
    'HttpProtocol',
    'IdentityType',
    'InventoryRuleType',
    'KeySource',
    'KeyType',
    'Kind',
    'LargeFileSharesState',
    'MinimumTlsVersion',
    'Name',
    'Permissions',
    'PrivateEndpointServiceConnectionStatus',
    'PublicAccess',
    'RootSquashType',
    'RoutingChoice',
    'RuleType',
    'Services',
    'ShareAccessTier',
    'SignedResource',
    'SignedResourceTypes',
    'SkuName',
    'State',
]


class AccessTier(str, Enum):
    """
    Required for storage accounts where kind = BlobStorage. The access tier used for billing.
    """
    HOT = "Hot"
    COOL = "Cool"


class Action(str, Enum):
    """
    The action of virtual network rule.
    """
    ALLOW = "Allow"


class BlobAccessTier(str, Enum):
    """
    The access tier of a storage blob.
    """
    HOT = "Hot"
    COOL = "Cool"
    ARCHIVE = "Archive"


class BlobType(str, Enum):
    """
    The type of a storage blob to be created.
    """
    BLOCK = "Block"
    APPEND = "Append"


class Bypass(str, Enum):
    """
    Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
    """
    NONE = "None"
    LOGGING = "Logging"
    METRICS = "Metrics"
    AZURE_SERVICES = "AzureServices"


class DefaultAction(str, Enum):
    """
    Specifies the default action of allow or deny when no other rules match.
    """
    ALLOW = "Allow"
    DENY = "Deny"


class DirectoryServiceOptions(str, Enum):
    """
    Indicates the directory service used.
    """
    NONE = "None"
    AADDS = "AADDS"
    AD = "AD"


class EnabledProtocols(str, Enum):
    """
    The authentication protocol that is used for the file share. Can only be specified when creating a share.
    """
    SMB = "SMB"
    NFS = "NFS"


class EncryptionScopeSource(str, Enum):
    """
    The provider for the encryption scope. Possible values (case-insensitive):  Microsoft.Storage, Microsoft.KeyVault.
    """
    MICROSOFT_STORAGE = "Microsoft.Storage"
    MICROSOFT_KEY_VAULT = "Microsoft.KeyVault"


class EncryptionScopeState(str, Enum):
    """
    The state of the encryption scope. Possible values (case-insensitive):  Enabled, Disabled.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class ExpirationAction(str, Enum):
    """
    The SAS expiration action. Can only be Log.
    """
    LOG = "Log"


class ExtendedLocationTypes(str, Enum):
    """
    The type of the extended location.
    """
    EDGE_ZONE = "EdgeZone"


class HttpProtocol(str, Enum):
    """
    The protocol permitted for a request made with the account SAS.
    """
    HTTPS_HTTP = "https,http"
    HTTPS = "https"


class IdentityType(str, Enum):
    """
    The identity type.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


class InventoryRuleType(str, Enum):
    """
    The valid value is Inventory
    """
    INVENTORY = "Inventory"


class KeySource(str, Enum):
    """
    The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage, Microsoft.Keyvault
    """
    MICROSOFT_STORAGE = "Microsoft.Storage"
    MICROSOFT_KEYVAULT = "Microsoft.Keyvault"


class KeyType(str, Enum):
    """
    Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped encryption key will be used. 'Service' key type implies that a default service key is used.
    """
    SERVICE = "Service"
    ACCOUNT = "Account"


class Kind(str, Enum):
    """
    Required. Indicates the type of storage account.
    """
    STORAGE = "Storage"
    STORAGE_V2 = "StorageV2"
    BLOB_STORAGE = "BlobStorage"
    FILE_STORAGE = "FileStorage"
    BLOCK_BLOB_STORAGE = "BlockBlobStorage"


class LargeFileSharesState(str, Enum):
    """
    Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class MinimumTlsVersion(str, Enum):
    """
    Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
    """
    TLS1_0 = "TLS1_0"
    TLS1_1 = "TLS1_1"
    TLS1_2 = "TLS1_2"


class Name(str, Enum):
    """
    Name of the policy. The valid value is AccessTimeTracking. This field is currently read only
    """
    ACCESS_TIME_TRACKING = "AccessTimeTracking"


class Permissions(str, Enum):
    """
    The signed permissions for the service SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p).
    """
    R = "r"
    D = "d"
    W = "w"
    L = "l"
    A = "a"
    C = "c"
    U = "u"
    P = "p"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"


class PublicAccess(str, Enum):
    """
    Specifies whether data in the container may be accessed publicly and the level of access.
    """
    CONTAINER = "Container"
    BLOB = "Blob"
    NONE = "None"


class RootSquashType(str, Enum):
    """
    The property is for NFS share only. The default is NoRootSquash.
    """
    NO_ROOT_SQUASH = "NoRootSquash"
    ROOT_SQUASH = "RootSquash"
    ALL_SQUASH = "AllSquash"


class RoutingChoice(str, Enum):
    """
    Routing Choice defines the kind of network routing opted by the user.
    """
    MICROSOFT_ROUTING = "MicrosoftRouting"
    INTERNET_ROUTING = "InternetRouting"


class RuleType(str, Enum):
    """
    The valid value is Lifecycle
    """
    LIFECYCLE = "Lifecycle"


class Services(str, Enum):
    """
    The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
    """
    B = "b"
    Q = "q"
    T = "t"
    F = "f"


class ShareAccessTier(str, Enum):
    """
    Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
    """
    TRANSACTION_OPTIMIZED = "TransactionOptimized"
    HOT = "Hot"
    COOL = "Cool"
    PREMIUM = "Premium"


class SignedResource(str, Enum):
    """
    The signed services accessible with the service SAS. Possible values include: Blob (b), Container (c), File (f), Share (s).
    """
    B = "b"
    C = "c"
    F = "f"
    S = "s"


class SignedResourceTypes(str, Enum):
    """
    The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs for blobs, queue messages, table entities, and files.
    """
    S = "s"
    C = "c"
    O = "o"


class SkuName(str, Enum):
    """
    The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called accountType.
    """
    STANDARD_LRS = "Standard_LRS"
    STANDARD_GRS = "Standard_GRS"
    STANDARD_RAGRS = "Standard_RAGRS"
    STANDARD_ZRS = "Standard_ZRS"
    PREMIUM_LRS = "Premium_LRS"
    PREMIUM_ZRS = "Premium_ZRS"
    STANDARD_GZRS = "Standard_GZRS"
    STANDARD_RAGZRS = "Standard_RAGZRS"


class State(str, Enum):
    """
    Gets the state of virtual network rule.
    """
    PROVISIONING = "provisioning"
    DEPROVISIONING = "deprovisioning"
    SUCCEEDED = "succeeded"
    FAILED = "failed"
    NETWORK_SOURCE_DELETED = "networkSourceDeleted"
