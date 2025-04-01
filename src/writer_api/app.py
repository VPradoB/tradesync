from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def index():
    return {"te amo": "mi vida linda"}
