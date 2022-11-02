# СКРИПТ ТЕСТИРОВАНИЯ ПРОПОРЦИИ Balancer

from utils.async_lib import init_async
import time 

# # http://balancer-domain/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8
# docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8
# curl http://localhost:3200/api/get_settings
# curl -I http://localhost:3200/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u


url='http://0.0.0.0:3200/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8'
urls = [[url] for i in range(0,1000)]

# асинхронно загружает N
t0=time.time()
res=init_async(urls)

codes={}

# print(res[0])
# exit()
def count_codes(c):
    c=str(c); codes[c]=1 if not c in codes else codes[c]+1
test=[count_codes('main' if ("//s1." in el['res']['headers']['location']) else 'cdn') for el in res]  #el['res']['code']) for el in res]

print(len(res), '{:.2f}'.format(time.time()-t0), 'secs')
print(codes, '{:2f}'.format(codes['cdn']/(len(res))*100),'%') # вывод пропроции %