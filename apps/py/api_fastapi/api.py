
import uvicorn, os, sys ; sys.path.append('./') ; sys.path.append('../')
from fastapi import FastAPI #, Request, Response , Form, Depends, UploadFile, File
#from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware

# отключениие публичной документации
app = FastAPI(docs_url=None, redoc_url=None, openapi_url=None); app.openapi = None # None
#app.mount("/static", StaticFiles(directory="dumps"), name="static");#templates = Jinja2Templates(directory="templates"); #app.mount("/static", StaticFiles(directory="static"), name="static")

app.add_middleware(CORSMiddleware,
    allow_origins=["*"], allow_credentials=True,
    allow_methods=["*"], allow_headers=["*"], )

########################################################
from services import Interfaces

@app.on_event("startup")
async def on_startup():
    await Interfaces.preload_settings() # подгрузка настроек
    i = Interfaces.RedirectVideoService
    x=i.StartAsThread_Worker( i.WorkerBlanaceForgotten, () ) # сервис контроля истории балансировки

from routes import API; API(app, Interfaces) # подключение api
########################################################

if __name__ == '__main__':
    uvicorn.run(app,  host='0.0.0.0', port=int(os.environ.get("API_PORT", 3000)))

