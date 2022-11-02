
import json as jsonutils
from fastapi import Request
from fastapi.responses import Response #, RedirectResponse

def def_response(content):
    return Response(jsonutils.dumps(content), media_type="application/json") 
class InitSettings:
    def __init__(self, app=None, interfaces=None): 
        Settings = interfaces.models.Settings
        RedirectVideoService = interfaces.RedirectVideoService

        @app.get("/api/get_settings")
        async def answer(): 
            return def_response(await Settings.get_formated_settings())
        @app.post("/api/edit_settings")
        async def answer(request: Request): 
            body = await request.json() ; body['value'] = jsonutils.dumps(body['value']) # validate to save as str
            updated = await Settings.edit_settings(body)
            # обновление настроек
            RedirectVideoService.updateSettings(await Settings.get_formated_settings()) # update settings
            return def_response(updated)
