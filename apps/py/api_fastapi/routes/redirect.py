
#from fastapi import APIRouter , Depends, HTTPException, Header
from fastapi.responses import Response, RedirectResponse #, JSONResponse , HTMLResponse

# router = APIRouter( prefix="", tags=[], dependencies=[], responses={} )

async def root_redirect(url): #video_id, m3u8_file):
    global RedirectVideoService 
    if RedirectVideoService.max_requests_percent(): # случайный % Percent -> CDN
        url = RedirectVideoService.formatVideoRedirectUrl(url) 
        if not url: return Response('not found', status_code=404) # не найдено по паттерну
        return RedirectResponse(url, status_code=301)  # редирект на CDN
    # print(url, RedirectVideoService.get_percent())
    return RedirectResponse(url, status_code=301) # редирект на основной сервер # -> Origin Server


class InitRedirect:
    def __init__(self, app=None, interfaces=None): 
        global RedirectVideoService; RedirectVideoService = interfaces.RedirectVideoService

        @app.get('/')
        async def answer(video): return await root_redirect(video)
        @app.head('/')
        async def answer(video): return await root_redirect(video)


        # tests
        @app.get('/api/rps_test_json')
        async def answer(): return Response("test"*5000, media_type="application/json") 
        @app.head('/api/rps_test_redirect')
        async def answer(): return RedirectResponse("http://", status_code=301)
        @app.get('/api/rps_test_redirect')
        async def answer(): return RedirectResponse("http://", status_code=301)


       #\f"/video/<video_id>/<m3u8_file>", methods=["HEAD", "GET"]
