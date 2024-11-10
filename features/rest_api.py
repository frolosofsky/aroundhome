import requests

class RestApi:
    def __init__(self, endpoint):
        self.endpoint = endpoint

    def healthcheck(self):
        assert requests.get(self.endpoint+'/health').status_code == 200

    def match_partners(self, material, lat, long):
        params = {
            'material': material,
            'address': f'{lat};{long}'
        }
        return requests.get(self.endpoint + '/match', headers={'Content-Type': 'application/json'}, params=params)

    def get_partner(self, id):
        return requests.get(self.endpoint + f'/partners/{id}', headers={'Content-Type': 'application/json'})
