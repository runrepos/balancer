# модули
import models # init DB
from services.redirect_video import RedirectVideoService 

storage={}

# интерфейсы передаваемые
class Interfaces:
    models = models
    RedirectVideoService = RedirectVideoService

    # preload db
    @staticmethod
    async def preload_settings():
        settings=await models.Settings.get_formated_settings()
        RedirectVideoService.updateSettings(settings) ; print('settings releaded:',settings) # check settings

    # for apidaora
    # @staticmethod
    # def send_callback(id, obj, callback):
    #     if not id in storage: storage[id]=callback
    #     else: 
    #         if isinstance(obj, dict): storage[id].update(callback)
    #         if isinstance(obj, list): storage[id].append(callback)
    # @staticmethod
    # def get_callbacks(): return storage