# -*- coding: utf-8 -*-

# Copyright (C) XenonStack, Pvt. Ltd - All Rights Reserved
# Unauthorized copying of this file, via any medium is strictly prohibited
# Proprietary and confidential

__author__ = "Gursimran Singh"
__copyright__ = "Copyright 2016, XenonStack Pvt. Ltd."
__license__ = "Proprietary"
__email__ = "gursimran@xenonstack.com"

from flask import jsonify
from ..models import BaseClass
from .. import db, memcache
import json
from bson import ObjectId
from datetime import datetime

class JSONEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, ObjectId):
            return str(o)
        return json.JSONEncoder.default(self, o)


class DeployModel(BaseClass):
    def __init__(self):
        self.required_list = []
        self.allowed_list = ['_id', 'name', 'plays', 'stats', 'completed_at', 'started_at', 'status', 'hosts', 'project_id']
        self.apps_collection = db.deployment
        self.allowed_extension = \
             {"_id": str,
              "name": str,
              "project_id": list,
              "plays": list,
              "stats": dict,
              "completed_at": datetime,
              "started_at": datetime,
              "status": str,
              "hosts": list
             }
        super(DeployModel, self).__init__(self.required_list, self.allowed_list, self.apps_collection, self.allowed_extension)

    def dynamic_name(self, request):
        form = request.get_json()
        project_id = self._check_header(request)
        if project_id == False:
            project_id = 'default'
        self._check_form_id(form)
        projects_with_id = db.deployment.find_one({"_id": str(project_id)})
        if projects_with_id == None and project_id != 'default':
            self.errors.append('No project id is created with this name')
            return jsonify({"status": False, "errors": self.errors, "message": "your project id id not default and no id is already present with this name"})
        else:
            return self.update_and_insert_check(form, project_id)


class Deploy(object):

    def __init__(self):
        pass

    @staticmethod
    def insert_log(form):
        db.runner.insert_one(form)

    @staticmethod
    def insert_deployment(form):
        db.deployment.insert_one(form)

    @staticmethod
    def set_nodes_cache(dictionary):
        for key, value in dictionary.iteritems():
            memcache.client.set(key, value)
        return True

    @staticmethod
    def get_nodes_cache(key):
        return memcache.client.get(key)

    @staticmethod
    def deployment_output(name):
        return db.runner.find({'name': name}).sort("timestamp", 1)

    @staticmethod
    def task_based_output(deplyment_name, task_slug):
        return db.runner.find({'name': deplyment_name, 'task-slug': task_slug}).sort("timestamp", 1)

    @staticmethod
    def task_status(deplyment_name, task_slug):
        return db.runner.find_one({'name': deplyment_name, 'task-slug': task_slug, "failed": True})

    @staticmethod
    def task_stats(deplyment_name):
        return db.runner.find({"name" : deplyment_name, "stats": {'$exists': True}})

    @staticmethod
    def task_names(deplyment_name):
        # aggregation = db.runner.aggregate([{
        #     '$group': {
        #         '_id': "$task-slug",
        #         'name': {
        #             '$first': "$task"
        #         },
        #         'failed': {
        #             '$min': "$res.failed"
        #         },
        #         'time_started': {
        #             '$min': "$timestamp"
        #         }
        #     }
        # },
        #     {'$match': {'name': deplyment_name}},
        # {
        #     '$sort': {
        #         'time_started': 1
        #     }
        # }])
        aggregation = db.runner.aggregate([{
            '$match': {
                    'name': deplyment_name
                }
            }, {
                '$group': {
                    '_id': "$task-slug",
                    'name': {
                        '$first': "$task"
                    },
                    'failed': {
                        '$min': "$res.failed"
                    },
                    'time_started': {
                        '$min': "$timestamp"
                    }
                }
            }, {
                '$sort': {
                    'time_started': 1
                }
            }
        ])
        return aggregation

    @staticmethod
    def set_current_deployment(form):
        return db.current_deployment.insert_one(form)

    @staticmethod
    def get_current_deployment():
        return db.current_deployment.find()

    @staticmethod
    def flush_current_deployment():
        return db.current_deployment.remove({})

    @staticmethod
    def deployment_competed(deplyment_name):
        return bool(db.runner.find_one({'name': deplyment_name, 'completed': True}))
