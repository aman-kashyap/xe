from .. import db
from ..models import BaseClass
from flask import jsonify

""" models.py: app.projects.models """

__author__ = "Arvind Sharma"
__date__ = "23/11/16"


project_collection = db.project


class CredModel(BaseClass):
    def __init__(self):
        self.required_list = ["_id", 'ssh_username']
        self.allowed_list = ['_id', 'ssh_username', 'ssh_pass', 'ssh_key', 'sudo', 'project_id']
        self.either_one_required = ["ssh_pass", "ssh_key"]
        self.creds_collection = db.creds
        self.allowed_extension = {"_id": str,
                        "project_id": list,
                        "ssh_username": str,
                        "ssh_pass": str,
                        "ssh_key": str,
                        "sudo": bool
                        }
        super(CredModel, self).__init__(self.required_list, self.allowed_list, self.creds_collection,
                                        self.allowed_extension)
    
    def cred_insert_condition(self, request):
        form = request.get_json()
        project_id = self._check_header(request)
        if project_id == False:
            project_id = 'default'
        self._check_form_id(form)
        projects_with_id = project_collection.find_one({"_id":str(project_id)})
        if projects_with_id == None and project_id != 'default':
            self.errors.append("No project id is created with this name")
            return jsonify({"status": False, "errors": self.errors, "message":"your project id id not default and no id is already present with this name"})
        else:
            if form['_id'] == None:
                self.errors.append("_id is required")
                return jsonify({"status": False, "errors": self.errors, "message":"name field or _id field is nessceray"})
            else:
                if self._check_creds(form, self.either_one_required) == False:
                    self.errors.append("required field does not match")
                    return jsonify({"status": False, "errors": self.errors, "message":"type and total no. of required field must match" }) 
                return self.update_and_insert_check(form,project_id)
