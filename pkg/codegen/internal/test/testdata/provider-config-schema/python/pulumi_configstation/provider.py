# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 favorite_color: Optional[pulumi.Input[Union[str, 'Color']]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[Union[str, 'Color']] favorite_color: this is a relaxed string enum which can also be set via env var
        """
        if favorite_color is None:
            favorite_color = _utilities.get_env('FAVE_COLOR')
        if favorite_color is not None:
            pulumi.set(__self__, "favorite_color", favorite_color)

    @property
    @pulumi.getter(name="favoriteColor")
    def favorite_color(self) -> Optional[pulumi.Input[Union[str, 'Color']]]:
        """
        this is a relaxed string enum which can also be set via env var
        """
        return pulumi.get(self, "favorite_color")

    @favorite_color.setter
    def favorite_color(self, value: Optional[pulumi.Input[Union[str, 'Color']]]):
        pulumi.set(self, "favorite_color", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 favorite_color: Optional[pulumi.Input[Union[str, 'Color']]] = None,
                 __props__=None):
        """
        Create a Configstation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[str, 'Color']] favorite_color: this is a relaxed string enum which can also be set via env var
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Configstation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 favorite_color: Optional[pulumi.Input[Union[str, 'Color']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if favorite_color is None:
                favorite_color = _utilities.get_env('FAVE_COLOR')
            __props__.__dict__["favorite_color"] = pulumi.Output.from_input(favorite_color).apply(pulumi.runtime.to_json) if favorite_color is not None else None
        super(Provider, __self__).__init__(
            'configstation',
            resource_name,
            __props__,
            opts)

