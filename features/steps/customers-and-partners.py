from behave import *

@given(u'partners')
def step_impl(context):
    raise NotImplementedError(u'STEP: Given partners')


@when(u'a customer located in "{city}" ("{long}", "{lat}") searches someone experienced with "{tool}"')
def step_impl(context, city, long, lat):
    raise NotImplementedError(u'STEP: When a customer located in ... searches someone experienced with "..."')


@then(u'they should receive a list of partners')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then they should receive a list of partners')


@then(u'they should receive an empty list of partners')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then they should receive an empty list of partners')


@when(u'a customer asks details about partner "{partner}"')
def step_impl(context, partner):
    raise NotImplementedError(u'STEP: When a customer asks details about partner "..."')


@then(u'they should receive parter details as such')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then they should receive parter details as such')


@then(u'they should receive error ({code})')
def step_impl(context, code):
    raise NotImplementedError(u'STEP: Then they should receive error (404)')
