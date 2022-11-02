


from routes.redirect import InitRedirect
from routes.settings import InitSettings
cls=[
        InitRedirect,
        InitSettings,
    ]
# interface
class API:
    def __init__(self, app, models):
        [init(app, models) for init in cls]