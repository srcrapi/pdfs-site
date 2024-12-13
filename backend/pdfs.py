import os
from urllib.parse import quote

import dotenv
import httpx

dotenv.load_dotenv()


async def pdfs_handler():
    github_api_url = "https://api.github.com"
    pdfs_url = f"{github_api_url}/repos/kyotosplit/books/contents/pdfs?ref=main"
    github_token = os.getenv("GITHUB_TOKEN")

    headers = {"Authorization": f"token {github_token}"}

    resp = httpx.get(pdfs_url, headers=headers)

    if resp.status_code != 200:
        return {"error": resp.raise_for_status()}

    pdfs = [
        {"name": item["name"], "url": quote(item["download_url"], safe=":/")}
        for item in resp.json()
    ]

    return pdfs
