from . import projects
from flask import request,jsonify
# from flask_jwt import jwt_required, current_identity
# from .. auth.models import authorize
from models import ProjectModel
import json


""" views.py: app.projects.views """

__author__ = "Arvind Sharma"
__date__ = "23/11/16"


@projects.route('/v1/projects', methods=['GET', 'POST', 'PUT'])
def project():
    project_data = ProjectModel()
    if request.method == 'POST':
        post = project_data.insert(request.get_json())
        if post is True:
            return jsonify({"status": True, "errors": None, "message":"data posted correctly"})
        else:
            return post
    elif request.method == 'GET':
        if 'limit' in request.args:
            fetched_value = project_data.fetch(request, limit=True)
        else:
            fetched_value = project_data.fetch(request, limit=False)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    elif request.method == 'PUT':
        if project_data.update(request.form):
            return jsonify({"status": True, "errors": None, "message":"content_updated"})


@projects.route('/v1/projects/<id>', methods=['GET', 'DELETE'])
def get_specific(id):
    project_data = ProjectModel()
    if request.method == 'GET':
        fetched_value = project_data.fetch_by_id(id)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    if request.method == 'DELETE':
        if project_data.remove_by_id(id):
            return jsonify({"status": True, "errors": None, "message":"document deleted"})
