# -*- coding: utf-8 -*-

# Copyright (C) XenonStack, Pvt. Ltd - All Rights Reserved
# Unauthorized copying of this file, via any medium is strictly prohibited
# Proprietary and confidential

from . import health
from flask import request, jsonify, current_app
# from flask_jwt import jwt_required, current_identity
# from .. auth.models import authorize
from models import Health
import json


__author__ = "Gursimran Singh"
__copyright__ = "Copyright 2016, XenonStack Pvt. Ltd."
__license__ = "Proprietary"
__email__ = "gursimran@xenonstack.com"


@health.before_request
def before_app_request():
    health_obj = Health()
    health_obj.set_token()


@health.route('/v1/health/token', methods=['POST', 'GET'])
def health_token():
    health_obj = Health()
    nexadash_url = current_app.config.get('NEXADASH_URL')

    if request.method == 'POST':
        token = request.get_json().get('token')
        # health = health_obj.check_token(token)
        return jsonify(token=bool(token), url=nexadash_url)
    if request.method == 'GET':
        health_status = True
        # health_status = health_obj.health_status()
        return jsonify(status=health_status)
