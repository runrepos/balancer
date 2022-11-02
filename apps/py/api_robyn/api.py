
# https://github.com/sansyrox/robyn
# треш фреймфорк (написан энтузиастом на rust)

from robyn import Robyn

app = Robyn(__file__)

@app.get("/")
async def h(request):
    return "Hello, world!"

if __name__ == '__main__': app.start(port=5000, url="localhost")
# ❗ не устанавлияается host 0.0.0.0, по этому нельзя использоваь в докере