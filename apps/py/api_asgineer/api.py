import asgineer
# ❗ средний фреймворк, медленнее sanic

# # Uvicorn:
# $ uvicorn example.py:main --host=localhost --port=8080
# # Hypercorn:
# $ hypercorn example.py:main --bind=localhost:8080
# # Daphne:
# $ daphne example:main --bind=localhost --port=8080

@asgineer.to_asgi
async def main(request):
     return 200, {}, f"<html>You requested <b>{request.path}</b></html>"

#if __name__ == '__main__':
#asgineer.run('uvicorn', main, 'localhost:3000')