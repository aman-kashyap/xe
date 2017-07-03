from flask import Blueprint

""" __init__.py: app.auth """

__author__ = "Arvind Sharma"
__date__ = "18/11/16"


auth = Blueprint('auth', __name__)

from . import views
