from datetime import datetime
from blacksheep import Application
# ❗ средний фреймворк, медленнее sanic

# # Uvicorn:
# $ uvicorn example.py:main --host=localhost --port=8080
# # Hypercorn:
# $ hypercorn example.py:main --bind=localhost:8080
# # Daphne:
# $ daphne example:main --bind=localhost --port=8080

res='test'*5000

app = Application()

@app.route("/")
async def home():
    return res #f"Hello, World! {datetime.utcnow().isoformat()}"