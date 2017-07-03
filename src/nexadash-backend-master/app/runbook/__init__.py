from flask import Blueprint

""" __init__.py: app.runbook """

__author__ = "Arvind Sharma"
__date__ = "30/11/16"


runbook = Blueprint('runbook', __name__)

from . import views