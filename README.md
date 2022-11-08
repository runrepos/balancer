#### сделано на 5-8х фреймворках - Sanic, aiohttp, fastapi, fiber (go), fastrouter (go) 

 + sqlalchemy, gorm (go)

 + asyncmy/async_session 

 + percona:5.7

 + протестровано 2 версии python 3.10, 3.11, go 1.19
 
 + postman api


## 1) запуск бд

```
cd balancer/
docker-compose -f db.yml up -d

```

## 2) запуск сервисов (на выбор) - хост/порт на всех приложениях 0.0.0.0 и 3200


### Sanic - 12500 - 15275.80 rps
```

docker-compose -f api_sanic.yml up

```

### Go Fiber - 45000-56000.77 rps
```

docker-compose -f api_go_fiber.yml up

```

### aiohttp - 4812.64 rps
```

docker-compose -f api_aiohttp.yml up

```

### Go fastrouter - 38000-46000.77 rps
```

docker-compose -f api_go_fastrouter.yml up

```

### fastapi - 3800.24 rps
```

docker-compose -f api_fastapi.yml up

```

### (прототип, без бд) Rust actix - 95000-108000rps
```

docker-compose -f api_rust_actix.yml up

```

### (прототип, без бд)Rust warp - 68000-81579.94rps
```

docker-compose -f api_rust_warp.yml up

```

### (прототип, без бд) Rust rocket - 68000-81579.94rps
```

docker-compose -f api_rust_rocket.yml up

```

## проверка 301 в ручную (запустить несколько раз в терминале), иногда ссылка Location будет менятся
```
curl -i http://0.0.0.0:3200/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8
```


## тестирование процента % соотношения (показывается в терминале, запустить скрипт несколько раз), так же с помощью этого скрипта можно проверять изменение %, после изменения настроек через set_settings
```

python3 -m venv venv && . venv/bin/activate
pip3 install asyncio httpx pytz
python3 apps/py/test.py

```

Результат: {'main': 362, 'cdn': 638} 63.800000 %


## изменение %, cdn_host установка представлено в api (по умолчанию 30%)
[https://documenter.getpostman.com/view/23758491/2s8YRgqErM](https://documenter.getpostman.com/view/23758491/2s8YRgqErM)

либо, импортировать в postman - [docs/more_balancer.postman_collection.json]()

либо, импортировать в insomnia (v4) - [docs/insomnia.json]()

ссылка для POST/GET запроса (настроек):
```
POST http://localhost:3200/api/edit_settings
GET http://localhost:3200/api/get_settings

curl --location --request POST 'http://localhost:3200/api/edit_settings' \
--data-raw '{
	"tag": "VIDEO:CDN",
	"value": [{
            "host": "http://localhost:3200",
            "percent": 0.3
			 }]
}'
```


## тестирование RPS:

```
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8
```

(без принтов)


(прототип, без бд) rust/actix - (108314.21)

(прототип, без бд) rust/warp - (68000-81579.94)

(прототип, без бд) rust/rocket - (55730 - 67000)

📌 go/fiber Requests per second:    45000 - 56065.77 (56000)

📌 (нет Head метода -) go/fastrouter      70000-98539 (46862)

(не работает в докере) py3.10/robin - Requests per second:    21277.95 [#/sec] (на основе rust)

📌 py3.10/sanic - Requests per second:    12500 - 18675.80 [#/sec]

📌 py3.11/sanic - Requests per second:    20347.59 [#/sec] 

(не реализовано) py3.8/asgineer - Requests per second:    4909.31-6601

📌 py3.10/aiohttp - Requests per second:    4812.64

(нет 301) py3.10/apidaora Requests per second:    4598.42

📌 py3.10/fastapi - Requests per second:    3800.24-4500 [#/sec] 

(с принтами)
py3.10/sanic - Requests per second:    9229.51 [#/sec] (mean)

py3.11/sanic - Requests per second:    10070.64 [#/sec] (mean)


с репликами

swarm (3-4) sanic - 26877.76 (=233%)
swarm (2) sanic - 25000
swarm (6) sanic - 21000
swarm (6) aiohttp - 17543.09 (=341%) (65% от sanic)
swarm (6) aiohttp - 
swarm (3) fastapi - 6133 (=185%) (22% от sanic)
swarm (6) fastapi - 5131




Версию python можно изменить в ```dockers/python/Dockerfile```


## Очистка БД
```
docker-compose -f db.yml down
docker rm balancer_mysql
docker volume rm balancer_mysql
docker-compose -f db.yml up -d

# + перезапуск приложения
```


## дополнительные настройки проекта

```

/.env - пароль и имя от БД
apps/.env - настройки для сервисов, имя контейнера с бд, доп опции.

```

```
MYSQL_DATABASE=test
MYSQL_ROOT_PASSWORD=mysqlpass

MYSQL_HOST=balancer_mysql
MYSQL_PORT=3306

# процент остаточного лога запросов при изменении % балансировки
HISTORY_BALANCE_PERCENT_GLOBAL=.2
# процент забывания количества запросов %
BALANCE_AUTOMATICALLY_FORGOTTEN=.7
# интервал забывания количества запросов (Сек)
BALANCE_AUTOMATICALLY_FORGETTEN_TIME_INTERVAL=5

```


Струкрута файлов:

```
migrations - структура бд
docs - документации
dockers - докер файлы

py/api_* - варианты api на python
py/models - orm для настроек бд
py/routes - роуты
py/service - сервис редиректов

otherlang - реализации на других языках
```


https://web-frameworks-benchmark.netlify.app/result
