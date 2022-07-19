# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .get_twingate_connector import *
from .get_twingate_connectors import *
from .get_twingate_group import *
from .get_twingate_groups import *
from .get_twingate_remote_network import *
from .get_twingate_resource import *
from .get_twingate_resources import *
from .provider import *
from .twingate_connector import *
from .twingate_connector_tokens import *
from .twingate_group import *
from .twingate_remote_network import *
from .twingate_resource import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_twingate.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_twingate.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "twingate",
  "mod": "index/twingateConnector",
  "fqn": "pulumi_twingate",
  "classes": {
   "twingate:index/twingateConnector:TwingateConnector": "TwingateConnector"
  }
 },
 {
  "pkg": "twingate",
  "mod": "index/twingateConnectorTokens",
  "fqn": "pulumi_twingate",
  "classes": {
   "twingate:index/twingateConnectorTokens:TwingateConnectorTokens": "TwingateConnectorTokens"
  }
 },
 {
  "pkg": "twingate",
  "mod": "index/twingateGroup",
  "fqn": "pulumi_twingate",
  "classes": {
   "twingate:index/twingateGroup:TwingateGroup": "TwingateGroup"
  }
 },
 {
  "pkg": "twingate",
  "mod": "index/twingateRemoteNetwork",
  "fqn": "pulumi_twingate",
  "classes": {
   "twingate:index/twingateRemoteNetwork:TwingateRemoteNetwork": "TwingateRemoteNetwork"
  }
 },
 {
  "pkg": "twingate",
  "mod": "index/twingateResource",
  "fqn": "pulumi_twingate",
  "classes": {
   "twingate:index/twingateResource:TwingateResource": "TwingateResource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "twingate",
  "token": "pulumi:providers:twingate",
  "fqn": "pulumi_twingate",
  "class": "Provider"
 }
]
"""
)
