from flask import Blueprint

""" __init__.py: app.apps """

__author__ = "Arvind Sharma"
__date__ = "26/11/16"

apps = Blueprint('apps', __name__)

from . import views