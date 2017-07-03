from werkzeug.security import safe_str_cmp
from flask_jwt import current_identity


""" models.py: app.auth.models """

__author__ = "Arvind Sharma"
__date__ = "18/11/16"


class User(object):
    def __init__(self, id, username, password, role):
        self.id = id
        self.username = username
        self.password = password
        self.role = role

    def __str__(self):
        return "User(id='%s')" % self.id

users = [
    User(1, 'user1@gmail.com', 'abcxyz', 'admin'),
    User(2, 'user2@gmail.com', 'abcxyz', 'user'),
]

username_table = {u.username: u for u in users}
userid_table = {u.id: u for u in users}


def authenticate(username, password):
    user = username_table.get(username, None)
    if user and safe_str_cmp(user.password.encode('utf-8'), password.encode('utf-8')):
        return user


def identity(payload):
    user_id = payload['identity']
    print payload
    return userid_table.get(user_id, None)


class Authorize(object):
    def __init__(self, role):
        self.role = role

    def __call__(self, func):
        def authorize_and_call(*args, **kwargs):
            if not current_identity.has_key(role): 
                raise Exception('Unauthorized Access!')
            return func(*args, **kwargs)
        return authorize_and_call


def authorize(role):
    def wrapper(func):
        def authorize_and_call(*args, **kwargs):
            if not role in current_identity.__dict__['role']: 
                raise Exception('Unauthorized Access!')
            return func(*args, **kwargs)
        return authorize_and_call
    return wrapper

