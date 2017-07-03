# -*- coding: utf-8 -*-

# Copyright (C) XenonStack, Pvt. Ltd - All Rights Reserved
# Unauthorized copying of this file, via any medium is strictly prohibited
# Proprietary and confidential

from flask import Blueprint

""" __init__.py: app.deploy """

__author__ = "Gursimran Singh"
__copyright__ = "Copyright 2016, XenonStack Pvt. Ltd."
__license__ = "Proprietary"
__email__ = "gursimran@xenonstack.com"

deploy = Blueprint('deploy', __name__)

from . import views