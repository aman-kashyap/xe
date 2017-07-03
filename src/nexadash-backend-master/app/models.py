from flask import jsonify
from names import names, adverbs, adjectives
import random
from random import randint
from datetime import datetime


class BaseClass(object):
    
    def __init__(self, required_list, allowed_list, collection_name, allowed_extension):
        self.required_list = required_list
        self.allowed_list = allowed_list
        self.collection_name = collection_name
        self.allowed_extension = allowed_extension
        self.errors = []

    def _required_field(self, keys):
        for value in self.required_list:
            if value not in keys:
                return True
            else:
                pass
        return True

    def _allowed_field(self, keys):
        for key in keys:
            if key not in self.allowed_list:
                return True
            else:
                pass
        return True
    
    def _check_header(self, request):
        try:
            project_id = request.headers['Project']
            return project_id
        except:
            return False
    
    def _check_form_id(self, form):
        try:
            form['_id']
        except:
            form['_id'] = None
    
    def _allowed_field_type(self, dictionary):
        if dictionary.keys() == []:
            return True
        for keys in dictionary.keys():
            if type(dictionary[keys]) == list or type(dictionary[keys])==int or type(dictionary[keys])==dict or type(dictionary[keys])==datetime:
                if self.allowed_extension[keys] != type(dictionary[keys]):
                    return True
            else:
                if self.allowed_extension[keys] != type(dictionary[keys].encode('utf-8')):
                    return True
                else:
                    pass
        return True

    def _find_collections(self, form):
        app_name_present = self.collection_name.find_one({"_id": form['_id']})
        return app_name_present
    
    def _make_name(self):
        middle = []
        middle.append(adverbs)
        middle.append(adjectives)
        app_name = str(random.sample(names, 1)[0])+'-'+str(random.sample(random.choice(middle), 1)[0])+'-'+'app'+str(randint(0, 999))
        return app_name
    
    def _update_app(self, form, project_list):
        self.collection_name.update({"_id": form['_id']}, {"$set": {"project_id": project_list}}, upsert=False, multi=True)
        return True

    def _check_creds(self, form, either_one_required):
        if any(map(lambda v: v in either_one_required, form.keys())):
            for i in either_one_required:
                if i in form.keys():
                    self.required_list.append(i)
                else:
                    pass
            return True
        else:
            return False

    def update_and_insert_check(self, form, project_id):
        if form['_id'] == None:
            app_name = self._make_name()
            form['_id'] = app_name
            form['project_id'] = [project_id]
            return self.insert(form)
        else:
            collection_name_present = self._find_collections(form)
            if collection_name_present == None:
                form['project_id'] = [project_id]
                return self.insert(form)
            else:
                project_list = collection_name_present['project_id']
                if project_id in project_list:
                    self.errors.append("project id is already present with this name")
                    return jsonify({"status": False, "errors": self.errors, "message":"cannot append as is already present with associated _id"})
                else:
                    project_list.append(project_id)
                    return self._update_app(form, project_list)

    def insert(self, form):
        keys = form.keys()
        if self._required_field(keys) and self._allowed_field(keys)and self._allowed_field_type(form):
            try:
                post_data = self.collection_name.insert_one(form).inserted_id
                return True
            except Exception as e:
                self.errors.append("duplicate entry")
                return jsonify({"status": False, "errors": self.errors, "message":"make another id this id is already present"})
        else:
            self.errors.append("payload not matches the allowed extension")
            return jsonify({"status": False, "errors": self.errors, "message":"either the passsing argument =s are more then allowed or type is different"})
    
    def fetch(self, request, limit):
        collection_object = self.collection_name
        args = request.args
        keys = args.keys()
        data = []
        header_status = self._check_header(request)
        if header_status != False and keys == []:
            data.append(self.collection_name.find({'project_id':{"$in":[request.headers['Project']]}}))
        else:
            if keys == []:
                data.append(self.collection_name.find({}))
            else:
                if limit:
                    for i in keys:
                        if 'limit' == i:
                            pass
                        else:
                            try: 
                                data.append(self.collection_name.find({i: args[i]}).limit(int(args['limit'])))
                            except:
                                pass
                else:
                    for i in keys:
                        try:
                            data.append(self.collection_name.find({i: args[i]}))
                        except:
                            pass
        return data
    
    def fetch_by_id(self, id):
        data = []
        data.append(self.collection_name.find({'_id': str(id)}))
        return data
    
    def update(self, form):
        keys = form.keys()
        try:
            form['_id']
        except:
            self.errors.append("name of update field is not given")
            return jsonify({"status": False, "errors": self.errors, "message":"name field is neccesary"})
        already_present = self.collection_name.find_one({'_id':form['_id']})
        if already_present == None:
            self.errors.append("this _id field is not present")
            return jsonify({"status": False, "errors": self.errors, "message":"name of the collection is not present cant update"})
        elif self._required_field(keys) and self._allowed_field(keys)and self._allowed_field_type(form):
            post = {}
            for i in keys:
                post[i] = form[i]
            self.collection_name.update({}, {"$set": post}, upsert=False, multi=True)
            return True
        else:
            self.errors.append("payload not matches the allowed extension")
            return jsonify({"status": False, "errors": self.errors, "message":"arguments must match the allowed and required field and their type"})
    
    def remove_by_id(self, id):
        self.collection_name.remove({"_id": str(id)})
        return True
