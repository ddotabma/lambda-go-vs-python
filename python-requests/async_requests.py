import aiohttp
import asyncio
from aiohttp import ClientSession
from typing import List


async def get(session: ClientSession, url: str) -> dict:
    response = await session.get(url)  # Make single http request
    return await response.json()  # Obtain response body as dict


async def multiple_get_requests(url_list: List[str]):
    async with aiohttp.ClientSession() as session:
        return await asyncio.gather(*[get(session, i) for i in url_list])
