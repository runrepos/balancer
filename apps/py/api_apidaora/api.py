#import sys ; sys.path.append('./') ; sys.path.append('../')
from http import HTTPStatus
from typing import Dict

from jsondaora import jsondaora

from apidaora import BadRequestError, Header, Response, appdaora, json, route
# ❗ Нет редиректов 301 !!!!

#from services import Interfaces

def app_factory():
    #Interfaces.preload_settings()
    #from routes import API; API(app.router, Interfaces) # подключение api

    @route.get('/')
    async def get_you_controller() -> Response:
        try: return 'test'
        except YouWereNotFoundError as error: raise BadRequestError(name=error.name, info=error.info) from error

    return appdaora([get_you_controller])

app = app_factory() #appdaora([get_you_controller])