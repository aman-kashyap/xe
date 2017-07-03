from .. import db, memcache
from ..models import BaseClass
from flask import jsonify

""" models.py: app.projects.models """

__author__ = "Arvind Sharma"
__date__ = "23/11/16"

project_collection = db.project


class NodeModel(BaseClass):
    def __init__(self):
        self.required_list = ["_id", "ip" "creds"]
        self.allowed_list = ['_id', 'cluster', 'tag', 'project_id', 'ip', 'creds']
        self.node_collection = db.node

        self.allowed_extension = {"_id": str,
                        "cluster": int,
                        "ip": str,
                        "creds": dict,
                        "tag": list,
                        "project_id": list
                        }
        super(NodeModel, self).__init__(self.required_list, self.allowed_list, self.node_collection,
                                        self.allowed_extension)
    
    def node_insert_condition(self, request):
        form = request.get_json()
        project_id = self._check_header(request)
        if project_id == False:
            project_id = 'default'
        self._check_form_id(form)
        projects_with_id = project_collection.find_one({"_id": str(project_id)})
        if projects_with_id == None and project_id != 'default':
            self.errors.append('No project id is created with this name')
            return jsonify({"status": False, "errors": self.errors, "message": "your project id id not default and no id is already present with this name"})
        else:
            if form['_id'] == None:
                self.errors.append("_id is required")
                return jsonify({"status": False, "errors": self.errors, "message":"name field or _id field is nessceray"})
            else:
                return self.update_and_insert_check(form, project_id)


class Node(object):
    def __init__(self):
        pass

    @staticmethod
    def set_nodes_cache(dictionary):
        for key, value in dictionary.iteritems():
            memcache.client.set(key, value)
        return True

    @staticmethod
    def get_nodes_cache(key):
        return memcache.client.get(key)

