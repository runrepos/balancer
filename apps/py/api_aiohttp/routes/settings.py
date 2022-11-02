
import json as jsonutils
from aiohttp.web import Response, HTTPMovedPermanently, HTTPNotFound, json_response # 301
# https://docs.aiohttp.org/en/latest/web_exceptions.html

def def_response(content): return json_response(content)

class InitSettings:
    def __init__(self, app=None, interfaces=None): 
        Settings = interfaces.models.Settings
        RedirectVideoService = interfaces.RedirectVideoService


        async def get_settings(x): return def_response(await Settings.get_formated_settings())
        app.add_route('GET', '/api/get_settings', get_settings)

        async def edit_settings(request): 
            body = await request.json() ; body['value'] = jsonutils.dumps(body['value']) # validate to save as str
            updated = await Settings.edit_settings(body)
            # обновление настроек
            RedirectVideoService.updateSettings(await Settings.get_formated_settings()) # update settings
            return def_response(updated)
        app.add_route('POST', '/api/edit_settings', edit_settings)
