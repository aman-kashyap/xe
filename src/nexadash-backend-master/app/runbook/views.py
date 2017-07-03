from . import runbook
from flask import request,jsonify
# from flask_jwt import jwt_required, current_identity
# from .. auth.models import authorize
from models import RunBookModel
import json


@runbook.route('/v1/runbook', methods=['GET', 'POST', 'PUT'])
def runbooks():
    runbook_data = RunBookModel()
    if request.method == 'GET':
        if 'limit' in request.args:
            fetched_value = runbook_data.fetch(request, limit=True)
        else:
            fetched_value = runbook_data.fetch(request, limit=False)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    elif request.method == 'POST':
        post = runbook_data.runbook_insert_condition(request)
        if post == True:
            return jsonify({"status": True, "errors": None, "message":"data posted correctly"})
        else:
            return post
    elif request.method == 'PUT':
        if runbook_data.update(request.form):
            return jsonify({"status": True, "errors": None, "message":"content_updated"})


@runbook.route('/v1/runbook/<id>', methods=['GET', 'DELETE'])
def get_specific(id):
    runbook_data = RunBookModel()
    if request.method == 'GET':
        fetched_value = runbook_data.fetch_by_id(id)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    if request.method == 'DELETE':
        if runbook_data.remove_by_id(id):
            return jsonify({"status": True, "errors": None, "message": "document deleted"})
