# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .get_guest_configuration_assignment import *
from .get_guest_configuration_hcrpassignment import *
from .guest_configuration_assignment import *
from .guest_configuration_hcrpassignment import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.guestconfiguration.v20180630preview as v20180630preview
    import pulumi_azure_native.guestconfiguration.v20181120 as v20181120
    import pulumi_azure_native.guestconfiguration.v20200625 as v20200625
    import pulumi_azure_native.guestconfiguration.v20210125 as v20210125
else:
    v20180630preview = _utilities.lazy_import('pulumi_azure_native.guestconfiguration.v20180630preview')
    v20181120 = _utilities.lazy_import('pulumi_azure_native.guestconfiguration.v20181120')
    v20200625 = _utilities.lazy_import('pulumi_azure_native.guestconfiguration.v20200625')
    v20210125 = _utilities.lazy_import('pulumi_azure_native.guestconfiguration.v20210125')

