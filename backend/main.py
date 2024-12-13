from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from pdfs import pdfs_handler

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5173"],
    allow_credentials=True,
    allow_methods=["GET"],
    allow_headers=["*"],
)


@app.get("/pdfs")
async def pdfs():
    pdfs_data = await pdfs_handler()

    return pdfs_data
