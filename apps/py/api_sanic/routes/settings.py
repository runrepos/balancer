from sanic.response import text, json as json, redirect
import json as jsonutils

class InitSettings:
    def __init__(self, app=None, interfaces=None):

        Settings = interfaces.models.Settings
        RedirectVideoService = interfaces.RedirectVideoService

        @app.get("/api/get_settings")
        async def get_settings(request):
            return json(await Settings.get_formated_settings())

        @app.post("/api/edit_settings")
        async def edit_settings(request): 
            body = request.json ; body['value'] = jsonutils.dumps(body['value']) # validate to save as str
            updated = await Settings.edit_settings(body)
            # обновление настроек сервиса
            RedirectVideoService.updateSettings(await Settings.get_formated_settings()) # update settings
            return json(updated)
