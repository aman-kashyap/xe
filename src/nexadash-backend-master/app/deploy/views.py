# -*- coding: utf-8 -*-

# Copyright (C) XenonStack, Pvt. Ltd - All Rights Reserved
# Unauthorized copying of this file, via any medium is strictly prohibited
# Proprietary and confidential

from . import deploy
from .. import memcache
from flask import request, jsonify, current_app, Response, url_for
# from flask_jwt import jwt_required, current_identity
# from .. auth.models import authorize
from models import Deploy, JSONEncoder, DeployModel
from datetime import datetime
import json
from bson import ObjectId
from bson.json_util import dumps
from slugify import slugify
import requests

__author__ = "Gursimran Singh"
__copyright__ = "Copyright 2016, XenonStack Pvt. Ltd."
__license__ = "Proprietary"
__email__ = "gursimran@xenonstack.com"


@deploy.route('/v1/deploy', methods=['POST', 'GET'])
def deploy_index():
    deploy_data = DeployModel()
    if request.method == 'GET':
        if 'limit' in request.args:
            fetched_value = deploy_data.fetch(request, limit=True)
        else:
            fetched_value = deploy_data.fetch(request, limit=False)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
        else:
            return fetched_value
    elif request.method == 'POST':
        print request.get_json()
        post = deploy_data.dynamic_name(request)
        if post is True:
            return jsonify({"status": True, "errors": None, "message": "Deployment posted correctly."})
        else:
            return post


@deploy.route('/v1/deploy/<id>', methods=['GET', 'DELETE'])
def get_deployment(id):
    deploy_data = DeployModel()
    if request.method == 'GET':
        fetched_value = deploy_data.fetch_by_id(id)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    if request.method == 'DELETE':
        if deploy_data.remove_by_id(id):
            return jsonify({"status": True, "errors": None, "message":"Deployment deleted."})


@deploy.route('/v1/runner/run', methods=['POST', 'GET'])
def deploy_pulse():
    dep_obj = Deploy()
    if request.method == 'POST':
        req = request.get_json()
        dep_obj.set_current_deployment(req)
        return jsonify(status=True)
    if request.method == 'GET':
        output = [output for output in dep_obj.get_current_deployment()]
        output = json.loads(dumps(output))
        print output
        dep_obj.flush_current_deployment()
        return jsonify(output)


@deploy.route('/v1/runner/<deployment_name>', methods=['POST', 'GET'])
def deploy_run(deployment_name):
    dep_obj = Deploy()
    if request.method == 'POST':
        dictionary2 = request.get_json()
        for d in dictionary2:
            if d.get('completeoutput'):
                payload = {}
                payload['output'] = d.get('completeoutput')
                payload['_id'] = slugify(deployment_name)
                payload['name'] = deployment_name
                dep_obj.insert_deployment(payload)
                del d['completeoutput']
            current_task = d.get('task', None)
            if current_task:
                memcache.client.set('current_task', current_task)
            else:
                d['name'] = deployment_name
                d['timestamp'] = datetime.now()
                if d.get('stats'):
                    d['completed'] = True
                    d['completed_time'] = datetime.now()
                if d.get('res', None):
                    if d.get('res', None).get('msg') == 'All items completed':
                        memcache.client.set('current_task', None)
                if memcache.client.get('current_task', None) != "None":
                    d['task'] = memcache.client.get('current_task', None)
                    d['task-slug'] = slugify(d['task'])
                dep_obj.insert_log(d)
        return jsonify(status=True)
    if request.method == 'GET':
        output = json.loads(dumps(dep_obj.deployment_output(deployment_name)))
        return jsonify(status=True, output=output)


@deploy.route('/v1/runner/<deployment_name>/complete')
def completed(deployment_name):
    dep_obj = Deploy()
    completed = dep_obj.deployment_competed(deployment_name)
    if completed:
        return jsonify(status=True)
    return jsonify(status=False)


@deploy.route('/v1/runner/<deployment_name>/tasks')
def task_names(deployment_name):
    dep_obj = Deploy()
    output = json.loads(dumps(dep_obj.task_names(deployment_name)))
    completed = dep_obj.deployment_competed(deployment_name)
    if completed:
        return jsonify(status=True, output=output)
    return jsonify(status=False, output=output)


@deploy.route('/v1/runner/<deployment_name>/<task_slug>')
def task_run(deployment_name, task_slug):
    dep_obj = Deploy()
    output = json.loads(dumps(dep_obj.task_based_output(deployment_name, task_slug)))
    return jsonify(status=True, output=output)


@deploy.route('/v1/runner/<deployment_name>/<task_slug>/status')
def task_status(deployment_name, task_slug):
    dep_obj = Deploy()
    output = json.loads(dumps(dep_obj.task_status(deployment_name, task_slug)))
    if output:
        return jsonify(status=True)
    return jsonify(status=False)


@deploy.route('/v1/runner/<deployment_name>/stats')
def task_statistics(deployment_name):
    dep_obj = Deploy()
    output = json.loads(dumps(dep_obj.task_stats(deployment_name)))
    print output
    if output:
        return jsonify(status=True, output=output)
    return jsonify(status=False, output=output)
