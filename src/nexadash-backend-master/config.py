

import os
basedir = os.path.abspath(os.path.dirname(__name__))


class Config(object):
    DEBUG = False
    TESTING = False

    @staticmethod
    def init_app(app):
        pass


class DevelopmentConfig(Config):
    SECRET_KEY = 's3cr3t'
    DEBUG = True
    UPLOAD_FOLDER = basedir+'/app/static/uploads'
    ALLOWED_EXTENSIONS = set(['jpg', 'jpeg'])
    ALLOWED_EXTENSIONS1 = set(['pdf', 'doc', 'docx'])
    MAX_CONTENT_LENGTH = 16 * 1024 * 1024
    MONGO_HOST = 'hm01'
    MONGO_PORT = 27017
    MONGO_DBNAME = 'nexastack'
    PORT = 5001
    CORS_HEADERS = 'Content-Type'
    NEXADASH_URL = (os.getenv('NEXADASH_URL') or 'http://203.100.70.58:5001/')
    PYMEMCACHE = {
        'server': ('172.31.4.105', 11211),
        'connect_timeout': 1.0,
        'timeout': 0.5,
        'no_delay': True,
        'key_prefix': b'nexadash-',
    }
    PLAYBOOKS = '/static/playbooks/'


class ProductionConfig(Config):
    SECRET_KEY = 'asdasdasd'
    DEBUG = False
    UPLOAD_FOLDER = basedir+'/app/static/uploads'
    ALLOWED_EXTENSIONS = set(['jpg', 'jpeg'])
    ALLOWED_EXTENSIONS1 = set(['pdf', 'doc', 'docx'])
    MAX_CONTENT_LENGTH = 16 * 1024 * 1024
    MONGO_HOST = '172.16.0.2'
    MONGO_PORT = 27017
    MONGO_DBNAME = 'nexastack'
    PORT = 5011
    CORS_HEADERS = 'Content-Type'
    NEXADASH_URL = (os.getenv('NEXADASH_URL') or 'http://172.16.0.19:5001')
    PYMEMCACHE = {
        'server': ('172.16.0.2', 11211),
        'connect_timeout': 1.0,
        'timeout': 0.5,
        'no_delay': True,
        'key_prefix': b'nexadash-',
    }


class StaggingConfig(Config):
    SECRET_KEY = 's3cr3tt'
    DEBUG = True
    UPLOAD_FOLDER = basedir+'/app/static/uploads'
    ALLOWED_EXTENSIONS = set(['jpg', 'jpeg'])
    ALLOWED_EXTENSIONS1 = set(['pdf', 'doc', 'docx'])
    MAX_CONTENT_LENGTH = 16 * 1024 * 1024
    MONGO_HOST = '172.16.0.2'
    MONGO_PORT = 27017
    MONGO_DBNAME = 'nexastack'
    PORT = 5000
    NEXADASH_URL = (os.getenv('NEXADASH_URL') or 'http://172.16.0.19:5001')
    PYMEMCACHE = {
        'server': ('memcached', 11211),
        'connect_timeout': 1.0,
        'timeout': 0.5,
        'no_delay': True,
        'key_prefix': b'nexadash-',
    }
    CORS_HEADERS = 'Content-Type'


config = {
    'devel': DevelopmentConfig,
    'prod': ProductionConfig,
    'stag': StaggingConfig,
    'default': DevelopmentConfig,
}
