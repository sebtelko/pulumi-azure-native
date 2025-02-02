# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DiskCreateOption',
    'DiskEncryptionSetIdentityType',
    'DiskEncryptionSetType',
    'DiskStorageAccountTypes',
    'EncryptionType',
    'HyperVGeneration',
    'NetworkAccessPolicy',
    'OperatingSystemTypes',
    'SnapshotStorageAccountTypes',
]


class DiskCreateOption(str, Enum):
    """
    This enumerates the possible sources of a disk's creation.
    """
    EMPTY = "Empty"
    ATTACH = "Attach"
    FROM_IMAGE = "FromImage"
    IMPORT_ = "Import"
    COPY = "Copy"
    RESTORE = "Restore"
    UPLOAD = "Upload"


class DiskEncryptionSetIdentityType(str, Enum):
    """
    The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"


class DiskEncryptionSetType(str, Enum):
    """
    The type of key used to encrypt the data of the disk.
    """
    ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY = "EncryptionAtRestWithCustomerKey"
    ENCRYPTION_AT_REST_WITH_PLATFORM_AND_CUSTOMER_KEYS = "EncryptionAtRestWithPlatformAndCustomerKeys"


class DiskStorageAccountTypes(str, Enum):
    """
    The sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    PREMIUM_LRS = "Premium_LRS"
    STANDARD_SS_D_LRS = "StandardSSD_LRS"
    ULTRA_SS_D_LRS = "UltraSSD_LRS"


class EncryptionType(str, Enum):
    """
    The type of key used to encrypt the data of the disk.
    """
    ENCRYPTION_AT_REST_WITH_PLATFORM_KEY = "EncryptionAtRestWithPlatformKey"
    ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY = "EncryptionAtRestWithCustomerKey"
    ENCRYPTION_AT_REST_WITH_PLATFORM_AND_CUSTOMER_KEYS = "EncryptionAtRestWithPlatformAndCustomerKeys"


class HyperVGeneration(str, Enum):
    """
    The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
    """
    V1 = "V1"
    V2 = "V2"


class NetworkAccessPolicy(str, Enum):
    """
    Policy for accessing the disk via network.
    """
    ALLOW_ALL = "AllowAll"
    ALLOW_PRIVATE = "AllowPrivate"
    DENY_ALL = "DenyAll"


class OperatingSystemTypes(str, Enum):
    """
    The Operating System type.
    """
    WINDOWS = "Windows"
    LINUX = "Linux"


class SnapshotStorageAccountTypes(str, Enum):
    """
    The sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    PREMIUM_LRS = "Premium_LRS"
    STANDARD_ZRS = "Standard_ZRS"
