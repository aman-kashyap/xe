from . import nodes
from flask import request,jsonify
from models import NodeModel, Node
import json


@nodes.route('/v1/nodes', methods = ['GET','POST','PUT'])
def node():
    node_data = NodeModel()
    node_obj = Node()
    if request.method == 'GET':
        if 'limit' in request.args:
            fetched_value = node_data.fetch(request, limit = True)
        else:
            fetched_value = node_data.fetch(request, limit = False)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    value['live'] = node_obj.get_nodes_cache(value['ip'])
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    elif request.method == 'POST':
        post = node_data.node_insert_condition(request)
        if post == True:
            return jsonify({"status": True, "errors": None, "message":"data posted correctly"})
        else:
            return post
    elif request.method == 'PUT':
        if node_data.update(request.form):
            return jsonify({"status": True, "errors": None, "message":"content_updated"})


@nodes.route('/v1/nodes/<id>', methods=['GET', 'DELETE'])
def get_specific(id):
    node_data = NodeModel()
    if request.method == 'GET':
        fetched_value = node_data.fetch_by_id(id)
        if fetched_value:
            fetched_list = []
            for obj in fetched_value:
                for value in obj:
                    fetched_list.append(value)
            return json.dumps(fetched_list)
    if request.method == 'DELETE':
        if node_data.remove_by_id(id):
            return jsonify({"status": True, "errors": None, "message":"document deleted"})


@nodes.route('/v1/nodes/ping', methods=['POST'])
def nodes_status():
    node_obj = Node()
    if request.method == 'POST':
        node_data = NodeModel()
        np_list = []
        np_dict = {}
        np_dict['name'] = "test"
        np_dict['hosts'] = []
        fetched_value = node_data.fetch(request, limit=False)
        dictionary = request.get_json()
        node_obj.set_nodes_cache(dictionary)
        if fetched_value:
            for obj in fetched_value:
                for value in obj:
                    n_dict = {}
                    n_dict['hostname'] = value['ip']
                    np_dict['hosts'].append(n_dict)
            np_list.append(np_dict)
            return jsonify(np_list)
