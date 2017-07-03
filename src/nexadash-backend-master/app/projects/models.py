
from bson.objectid import ObjectId
import json
from ..models import BaseClass
from flask import jsonify
from .. import db

""" models.py: app.projects.models """

__author__ = "Arvind Sharma"
__date__ = "23/11/16"


class ProjectModel(BaseClass):

    def __init__(self):
        self.required_list = ["_id"]
        self.allowed_list = ['_id', 'tag']
        self.project_collection = db.project
        self.allowed_extension = {"_id": str,
                        "tag": list,
                        }
        super(ProjectModel, self).__init__(self.required_list, self.allowed_list, self.project_collection,
                                           self.allowed_extension)
