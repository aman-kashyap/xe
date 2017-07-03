from flask import Blueprint

""" __init__.py: app.creds """

__author__ = "Arvind Sharma"
__date__ = "26/11/16"

creds = Blueprint('creds', __name__)

from . import views