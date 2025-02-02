# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .application import *
from .application_group import *
from .get_application import *
from .get_application_group import *
from .get_host_pool import *
from .get_msix_package import *
from .get_scaling_plan import *
from .get_workspace import *
from .host_pool import *
from .msix_package import *
from .scaling_plan import *
from .workspace import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.desktopvirtualization.v20190123preview as v20190123preview
    import pulumi_azure_native.desktopvirtualization.v20190924preview as v20190924preview
    import pulumi_azure_native.desktopvirtualization.v20191210preview as v20191210preview
    import pulumi_azure_native.desktopvirtualization.v20200921preview as v20200921preview
    import pulumi_azure_native.desktopvirtualization.v20201019preview as v20201019preview
    import pulumi_azure_native.desktopvirtualization.v20201102preview as v20201102preview
    import pulumi_azure_native.desktopvirtualization.v20201110preview as v20201110preview
    import pulumi_azure_native.desktopvirtualization.v20210114preview as v20210114preview
    import pulumi_azure_native.desktopvirtualization.v20210201preview as v20210201preview
    import pulumi_azure_native.desktopvirtualization.v20210309preview as v20210309preview
else:
    v20190123preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20190123preview')
    v20190924preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20190924preview')
    v20191210preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20191210preview')
    v20200921preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20200921preview')
    v20201019preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20201019preview')
    v20201102preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20201102preview')
    v20201110preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20201110preview')
    v20210114preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20210114preview')
    v20210201preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20210201preview')
    v20210309preview = _utilities.lazy_import('pulumi_azure_native.desktopvirtualization.v20210309preview')

