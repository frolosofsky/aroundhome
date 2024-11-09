from behave import *
from rest_api import RestApi
import os
import psycopg

def before_all(context):
    try:
        context.rest_api = RestApi('http://' + os.getenv('bind', default='localhost:8080'))
        context.rest_api.healthcheck()
        context.db = psycopg.connect(conninfo=os.getenv('dbconn', default='postgres://postgres:pass@localhost/aroundhome?sslmode=disable'), autocommit=True)
    except:
        print()
        print('===== Make sure you run the server: docker compose up')
        print()
        raise    

def before_scenario(context, scenario):
    with context.db.cursor() as cur:
        cur.execute('delete from partner_skill')
        cur.execute('delete from partner')
