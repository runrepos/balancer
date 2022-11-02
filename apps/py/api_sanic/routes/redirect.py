from sanic.response import text, json, redirect


async def root_redirect(request): #video_id, m3u8_file):
    # global RedirectVideoService
    # if not 'video' in request.args: return text('not found',404)
    try: url = request.args['video'][0] # валидация
    except: return text('not found',status=404)
    # Percent -> CDN
    if RedirectVideoService.max_requests_percent(): # случайный %
        url = RedirectVideoService.formatVideoRedirectUrl(url) ; 
        # print(url, RedirectVideoService.get_percent()) #; return text(url, status=301)
        if not url: return text('not found', status=404) # не найдено по паттерну
        return redirect(url, status=301)  # редирект на CDN
    # -> Origin Server
    # print(url, RedirectVideoService.get_percent())
    return redirect(url, status=301) # редирект на основной сервер

    #return text(RedirectVideoService.getFromServerM3u8()) #


class InitRedirect:
    def __init__(self, app=None, interfaces=None): 
        global RedirectVideoService; RedirectVideoService = interfaces.RedirectVideoService
        # основной роут для редиректа
        app.add_route( root_redirect, f"/" , 
                       methods=["HEAD", "GET"] )

        # Для сверки RPS
        @app.get("/api/rps_test_json")
        async def app_test_json(request): 
            return json("test"*5000) 

        @app.get("/api/rps_test_redirect")
        async def app_redirect(request):
            return redirect("http://", status=301)

        #app.add_route(root, f"/video/<video_id>/<m3u8_file>", methods=["HEAD", "GET"])
