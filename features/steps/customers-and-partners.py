from behave import *
from behave.model import Table
import uuid

@given(u'partners')
def step_impl(context):
    context.partners = dict()
    with context.db.cursor() as cur:
        for row in context.table:
            lat, long = row['location'].split(', ')
            cur.execute('insert into partner (name, geo, radius, rating) values (%s, ST_MakePoint(%s, %s), %s, %s) returning id',
                        (row['name'], float(long), float(lat), meters(row['radius']), row['rating']))
            id = cur.fetchone()[0]
            for exp in row['experience'].split(', '):
                cur.execute('insert into partner_skill (partner_id, code) values (%s, %s)',
                            (id, exp))
            context.partners[row['name']] = id

@when(u'a customer located in "{city}" ("{lat}", "{long}") searches someone experienced with "{tool}"')
def step_impl(context, city, long, lat, tool):
    context.last_response = context.rest_api.match_partners(tool, lat, long)

@then(u'they should receive a list of partners')
def check_partners(context):
    assert context.last_response.status_code == 200, f'{last_response.status_code} != 200'
    j = context.last_response.json()
    assert len(j) == len(context.table.rows), f'{len(j)} != {len(context.table.rows)}'
    for i, row in enumerate(context.table):
        assert row['name'] == j[i]['name'], f'{row} != {j[i]}'
        assert row['rating'] == j[i]['rating'], f'{row} != {j[i]}'
        assert abs(j[i]['distance'] - meters(row['distance'])) < 1000, f'{row} != {j[i]}'

@then(u'they should receive an empty list of partners')
def step_impl(context):
    context.table = Table([])
    return check_partners(context)

@when(u'a customer asks details about partner "{partner}"')
def step_impl(context, partner):
    context.last_response = context.rest_api.get_partner(context.partners.get(partner, str(uuid.uuid4())))

@then(u'they should receive parter details as such')
def step_impl(context):
    assert context.last_response.status_code == 200, f'{last_response.status_code} != 200'
    j = context.last_response.json()
    r = context.table.rows[0]
    assert j['name'] == r['name'], f'{j} != {r}'
    assert set(j['skills']) == set(r['experience'].split(', ')), f'{j} != {r}'
    assert j['address'].split(',') == r['location'].split(', '), f'{j} != {r}'
    assert j['radius'] == meters(r['radius']), f'{j} != {r}'
    assert j['rating'] == int(r['rating']), f'{j} != {r}'

@then(u'they should receive error ({code})')
def step_impl(context, code):
    assert context.last_response.status_code == int(code), f'{last_response.status_code} != {code}'

def meters(name):
    parts = name.split(' ')
    assert len(parts) == 2, f'bad distance: {name}'
    assert parts[1] == 'km', f'bad distance: {name}'
    return int(parts[0]) * 1000
