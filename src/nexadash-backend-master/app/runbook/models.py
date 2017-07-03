from .. import db
from ..models import BaseClass
from flask import jsonify

""" models.py: app.projects.models """

__author__ = "Arvind Sharma, Gursimran"
__date__ = "23/11/16"

project_collection = db.project


class RunBookModel(BaseClass):
    def __init__(self):
        self.required_list = ["_id", "desc", "type", "platforms", "url", "inventory"]
        self.allowed_list = ["_id", "desc", "type", "platforms", "url", "deploys", "tags", "inventory"]
        self.runbook_collection = db.runbook
        self.allowed_extension = {"_id": str,
                        "desc": int,
                        "type": list,
                        "platforms": list,
                        "url": str,
                        "deploys": int,
                        "tags" : list,
                        "inventory": dict
                        }
        super(RunBookModel, self).__init__(self.required_list, self.allowed_list, self.runbook_collection,
                                           self.allowed_extension)

    def runbook_insert_condition(self, request):
        form = request.get_json()
        self._check_form_id(form)
        if form['_id'] == None:
            self.errors.append("_id is required")
            return jsonify({"status": False, "errors": self.errors, "message":"name field or _id field is nessceray"})
        else:
            return self.insert(form)
