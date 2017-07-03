# -*- coding: utf-8 -*-

# Copyright (C) XenonStack, Pvt. Ltd - All Rights Reserved
# Unauthorized copying of this file, via any medium is strictly prohibited
# Proprietary and confidential

from .. import memcache
from datetime import datetime, timedelta


__author__ = "Gursimran Singh"
__copyright__ = "Copyright 2016, XenonStack Pvt. Ltd."
__license__ = "Proprietary"
__email__ = "gursimran@xenonstack.com"


class Health(object):
    def __init__(self):
        self.client = memcache.client

    def set_token(self):
        return self.client.set('token', '4E3665ADEBF2D378D6631278CBFBD')

    def get_token(self,):
        return self.client.get('token')

    def check_token(self, received):
        if self.get_token() == received:
            self.client.set('healthy', 'True')
            self.client.set('healthy_time', str(datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f')))
            return True
        self.client.set('healthy', 'False')
        return False

    def check_health(self, token):
        # if
        return self.client.get('token')

    def health_status(self):
        set_time = self.client.get('healthy_time')
        set_time = datetime.strptime(set_time, '%Y-%m-%d %H:%M:%S.%f')
        time_now = datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f')
        time_now = datetime.strptime(time_now, '%Y-%m-%d %H:%M:%S.%f')
        tdelta = time_now - set_time
        healthy = bool(self.client.get('healthy'))
        if healthy is True and int(tdelta.total_seconds()) < 15:
            return True
        return False
