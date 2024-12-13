from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from pdfs import pdfs_handler

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5173", "https://pdfs-site.vercel.app"],
    allow_credentials=True,
    allow_methods=["GET"],
    allow_headers=["*"],
)


@app.get("/")
async def index():
    pdfs_data = await pdfs_handler()

    return pdfs_data
