from . import creds
from flask import request,jsonify
# from flask_jwt import jwt_required, current_identity
# from .. auth.models import authorize
from models import CredModel
import json


@creds.route('/v1/creds', methods = ['GET','POST','PUT'])
def node():
    cred_data = CredModel()
    if request.method == 'GET':
        if 'limit' in request.args:
            fetched_value = cred_data.fetch(request, limit = True)
        else:
            fetched_value = cred_data.fetch(request, limit = False)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    elif request.method == 'POST':
        post = cred_data.cred_insert_condition(request)
        if post == True:
            return jsonify({"status": True, "errors": None, "message":"data posted correctly"})
        else:
            return post
    elif request.method == 'PUT':
        if cred_data.update(request.form):
            return jsonify({"status": True, "errors": None, "message":"content_updated"})

@creds.route('/v1/creds/<id>', methods=['GET', 'DELETE'])
def get_specific(id):
    cred_data = CredModel()
    if request.method == 'GET':
        fetched_value = cred_data.fetch_by_id(id)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    if request.method == 'DELETE':
        if cred_data.remove_by_id(id):
            return jsonify({"status": True, "errors": None, "message":"document deleted"})
