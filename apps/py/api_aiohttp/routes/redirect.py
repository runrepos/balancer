
from aiohttp.web import Response, HTTPMovedPermanently, HTTPNotFound # 301
# https://docs.aiohttp.org/en/latest/web_exceptions.html

def initRoute(router):
    async def test(x): return Response(text='test')
    router.add_route('*', '/test', test)


async def root_redirect(request): #video_id, m3u8_file):
    global RedirectVideoService 
    try: 
        url = request.rel_url.query.get('video', None)
        if not url: return HTTPNotFound()
    except: return HTTPNotFound()
    #print(url)
    if RedirectVideoService.max_requests_percent(): # случайный % Percent -> CDN
        url = RedirectVideoService.formatVideoRedirectUrl(url) 
        if not url: return HTTPNotFound() # не найдено по паттерну
        return HTTPMovedPermanently(url )  # редирект на CDN
    # print(url, RedirectVideoService.get_percent())
    return HTTPMovedPermanently(url) # редирект на основной сервер # -> Origin 


class InitRedirect:
    def __init__(self, app=None, interfaces=None): 
        global RedirectVideoService; RedirectVideoService = interfaces.RedirectVideoService

        app.add_route('*', '/', root_redirect)

        # tests
        async def rps_test_json(x): return Response(text='test'*5000)
        app.add_route('*', '/api/rps_test_json', rps_test_json)
        async def rps_test_redirect(x): raise HTTPMovedPermanently('http://')
        app.add_route('*', '/api/rps_test_redirect', rps_test_redirect)



       #\f"/video/<video_id>/<m3u8_file>", methods=["HEAD", "GET"]
