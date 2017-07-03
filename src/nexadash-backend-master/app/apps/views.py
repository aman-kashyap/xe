from . import apps
from flask import request,jsonify
# from flask_jwt import jwt_required, current_identity
# from .. auth.models import authorize
from models import AppsModel
import json


@apps.route('/v1/apps', methods=['GET', 'POST', 'PUT'])
def app_request_handler():
    apps_data = AppsModel()
    if request.method == 'GET':
        if 'limit' in request.args:
            fetched_value = apps_data.fetch(request, limit=True)
        else:
            fetched_value = apps_data.fetch(request, limit=False)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
        else:
            return fetched_value
    elif request.method == 'POST':
        post = apps_data.Make_app_name(request)
        if post is True:
            return jsonify({"status": True, "errors": None, "message": "data posted correctly"})
        else:
            return post
    elif request.method == 'PUT':
        if apps_data.update(request.form):
            return jsonify({"status": True, "errors": None, "message": "content_updated"})


@apps.route('/v1/apps/<id>', methods=['GET', 'DELETE'])
def get_specific(id):
    apps_data = AppsModel()
    if request.method == 'GET':
        fetched_value = apps_data.fetch_by_id(id)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    if request.method == 'DELETE':
        if apps_data.remove_by_id(id):
            return jsonify({"status": True, "errors": None, "message":"document deleted"})

