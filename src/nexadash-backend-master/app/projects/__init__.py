from flask import Blueprint

""" __init__.py: app.projects """

__author__ = "Arvind Sharma"
__date__ = "23/11/16"


projects = Blueprint('projects', __name__)

from . import views