import os, sys, random, time, threading

CDN_HOST=os.environ.get('DEF_CDN_HOST', 'http://localhost:3200')
SERVER_ID=os.environ.get('DEF_SERVER_ID', 's1')

REDIRECT_PERCENT = .29 # 30%
# DEFAULT_RETURN_m3u8 = 'test'*50 # return data
HISTORY_BALANCE_PERCENT =  float(os.environ.get('HISTORY_BALANCE_PERCENT_GLOBAL', .2)) # функция балансировки истории, при изменении %
BALANCE_AUTOMATICALLY_FORGOTTEN = float(os.environ.get('HISTORY_BALANCE_PERCENT_GLOBAL', .7)) # // %
BALANCE_AUTOMATICALLY_FORGETTEN_TIME = float(os.environ.get('HISTORY_BALANCE_PERCENT_GLOBAL_INTERVAL', 5)) # сек


total_requests=[0,0] # total, 301




class RedirectVideoService:


    # расчет % для редиректа
    @staticmethod # вычисляет %
    def max_requests_percent():
        total_requests[0] += 1 # подсчет всех
        try: # проверка случайного %
            if random.random() < REDIRECT_PERCENT and \
            total_requests[1]/total_requests[0] < REDIRECT_PERCENT: # проверка переполнения %
                total_requests[1]+=1 # подсчет 301
                return True
        except: return

    # парсинг входящего урл
    # http://balancer-domain/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8
    @staticmethod # парсинг url
    def parseRequestUrl(url): # print(url)
        # для CDN:VIDEO
        if '/video/' in url:
            url = url.split('/video/')
            try: server_id = url[0].split('://')[1].split('.')[0]
            except: server_id = SERVER_ID or 's1'
            if url[1]: return server_id, url[1]
            #if url[1] and len(url[1])>1: return server_id, url[1]

    # формирование ссылки для редиректа
    @staticmethod # формат урл
    def formatVideoRedirectUrl(redirect_url): # video_id, m3u8_file): # "/video/1488/xcg2djHckad.m3u8"
        redirect_url = RedirectVideoService.parseRequestUrl(redirect_url)
        if not redirect_url: return # not found in pattern
        return f'{CDN_HOST}/{redirect_url[0]}/video/'+redirect_url[1]  #{video_id}/{m3u8_file}'

    @staticmethod # получение инфо о % вероятности перехода
    def get_percent():
        return '{0:.2f}'.format( total_requests[1] / total_requests[0]*100)+ \
               f'% = {total_requests[1]} / {total_requests[0]}'

    # обновление настроек
    @staticmethod
    def updateSettings(settings):
        if 'VIDEO:CDN' in settings:
            settings=settings['VIDEO:CDN']
            print("Updating settings ...", settings)
            global REDIRECT_PERCENT, CDN_HOST;
            try: REDIRECT_PERCENT = float(settings['value'][0]['percent'])
            except: pass
            try: CDN_HOST = float(settings['value'][0]['host'])
            except: pass
            total_requests[0] *= HISTORY_BALANCE_PERCENT # вывравнивание пропорций истории
            total_requests[1] *= HISTORY_BALANCE_PERCENT 

    # воркер для регулирования баланса количества редиректов
    @staticmethod
    def StartAsThread_Worker(func, args): # def (1,)
        x = threading.Thread(target=func, args=args) ; x.start() ; return x  # x.join()

    @staticmethod
    def WorkerBlanaceForgotten():
        while True:
            time.sleep(BALANCE_AUTOMATICALLY_FORGETTEN_TIME) # pause
            RedirectVideoService.CorrectBalance() ; print("Correct balance", 
                                        BALANCE_AUTOMATICALLY_FORGOTTEN, 
                                        total_requests[0]+total_requests[1] ) # сокращение истории
    # коррекция баланса счетчика количества переходов
    @staticmethod
    def CorrectBalance():
        total_requests[1] *= BALANCE_AUTOMATICALLY_FORGOTTEN
        total_requests[0] *= BALANCE_AUTOMATICALLY_FORGOTTEN


    # дополнительный интерфейс, можно выводить M3u8 сразу
    # @staticmethod
    # def getFromServerM3u8():
    #     return DEFAULT_RETURN_m3u8

