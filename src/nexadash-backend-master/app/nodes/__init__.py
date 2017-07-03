from flask import Blueprint

""" __init__.py: app.nodes """

__author__ = "Arvind Sharma"
__date__ = "26/11/16"

nodes = Blueprint('nodes', __name__)

from . import views