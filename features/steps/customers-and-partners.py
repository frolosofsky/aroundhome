from behave import *
from behave.model import Table

@given(u'partners')
def step_impl(context):
    with context.db.cursor() as cur:
        for row in context.table:
            lat, long = row['location'].split(', ')
            cur.execute('insert into partner (name, geo, radius, rating) values (%s, ST_MakePoint(%s, %s), %s, %s) returning id',
                        (row['name'], float(long), float(lat), meters(row['radius']), row['rating']))
            id = cur.fetchone()[0]
            for exp in row['experience'].split(', '):
                cur.execute('insert into partner_skill (partner_id, code) values (%s, %s)',
                            (id, exp))


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
    raise NotImplementedError(u'STEP: When a customer asks details about partner "..."')


@then(u'they should receive parter details as such')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then they should receive parter details as such')


@then(u'they should receive error ({code})')
def step_impl(context, code):
    raise NotImplementedError(u'STEP: Then they should receive error (404)')

def meters(name):
    parts = name.split(' ')
    assert len(parts) == 2, f'bad distance: {name}'
    assert parts[1] == 'km', f'bad distance: {name}'
    return int(parts[0]) * 1000
