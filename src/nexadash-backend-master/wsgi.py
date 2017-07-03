
""" wsgi.py: wsgi app """


import os
from gevent.wsgi import WSGIServer
from app import create_app
# from app import jwt
# from flask_jwt import jwt_required, current_identity

__author__ = "Arvind Sharma"
__date__ = "18/11/16"

app = create_app(os.getenv('FLASK_CONFIG') or 'devel')
port = app.config['PORT']


if __name__ == '__main__':
    # if os.getenv('FLASK_CONFIG') == 'prod':
    http_server = WSGIServer(('0.0.0.0', port), app, None, 'default')
    http_server.serve_forever()
    # else:
    #     app.run(host='0.0.0.0', port=port)
