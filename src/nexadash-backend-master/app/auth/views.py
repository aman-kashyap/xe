from . import auth
from flask_jwt import jwt_required, current_identity
from models import authorize

""" views.py: app.auth.views """

__author__ = "Arvind Sharma"
__date__ = "18/11/16"


@auth.route('/protected')
@jwt_required()
@authorize('admin')
def protected():
    return '%s' % current_identity



