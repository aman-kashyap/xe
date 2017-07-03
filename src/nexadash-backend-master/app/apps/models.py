from .. import db
from flask import jsonify
from ..models import BaseClass

""" models.py: app.projects.models """

__author__ = "Arvind Sharma"
__date__ = "23/11/16"


project_collection = db.project


class AppsModel(BaseClass):

    def __init__(self):
        self.required_list = []
        self.allowed_list = ['_id', 'type', 'project_id']
        self.apps_collection = db.apps
        self.allowed_extension = {"_id": str,
                            "project_id": list,
                            "type": str
                                  }
        super(AppsModel, self).__init__(self.required_list, self.allowed_list, self.apps_collection,
                                        self.allowed_extension)

    def Make_app_name(self, request):
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
            return self.update_and_insert_check(form, project_id)


