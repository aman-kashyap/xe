import os
from flask import Flask, current_app
from config import config
from pymongo import MongoClient
from flask_jwt import JWT
from auth.models import authenticate, identity
from flask_cors import CORS, cross_origin
from flask_pymemcache import FlaskPyMemcache

""" __init__.py: app """

__author__ = "Arvind Sharma"
__date__ = "18/11/16"

app_config = config[os.getenv('FLASK_CONFIG') or 'devel']
client = MongoClient(app_config.MONGO_HOST, app_config.MONGO_PORT)
db = client.nexastack
jwt = JWT()
cors = CORS()
memcache = FlaskPyMemcache()


def create_app(configuration, register_blueprints=True):
    app = Flask(__name__)
    app.config.from_object(config[configuration])
    config[configuration].init_app(app)
    cors.init_app(app)
    jwt = JWT(app,authenticate, identity)
    memcache.init_app(app)

    if register_blueprints:
        from .auth import auth as auth_blueprint
        app.register_blueprint(auth_blueprint)
        from .projects import projects as projects_blueprint
        app.register_blueprint(projects_blueprint)
        from .nodes import nodes as nodes_blueprint
        app.register_blueprint(nodes_blueprint)
        from .creds import creds as creds_blueprint
        app.register_blueprint(creds_blueprint)
        from .apps import apps as apps_blueprint
        app.register_blueprint(apps_blueprint)
        from .runbook import runbook as runbook_blueprint
        app.register_blueprint(runbook_blueprint)
        from .health import health as health_blueprint
        app.register_blueprint(health_blueprint)
        from .deploy import deploy as deploy_blueprint
        app.register_blueprint(deploy_blueprint)
    return app


