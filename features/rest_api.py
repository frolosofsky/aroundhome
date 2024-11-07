import requests

class RestApi:
    def __init__(self, endpoint):
        self.endpoint = endpoint

    def healthcheck(self):
        assert requests.get(self.endpoint+'/health').status_code == 200
