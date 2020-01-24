from dataclasses import dataclass
import asyncio
import aiohttp
from typing import List
import os

@dataclass
class Data:
    datetime: str
    values: List[int]

    def __post_init__(self):
        self.values = [int(i) for i in self.values]  # convert values to int
        assert isinstance(self.datetime, str)

def deserialize(d: dict):
    return Data(**d)


async def get(session, url: str) -> dict:
    response = await session.get(url)  # Make http requests
    return await response.json()       # Obtain body as dict


async def send_all_get_requests(url_list: List[str]):
    async with aiohttp.ClientSession() as session:
        return await asyncio.gather(*[get(session, i) for i in url_list])


def handler(event, context):
    results = asyncio.run(
        send_all_get_requests(
            [os.environ["API"] for _ in range(100)] # todo make var
        )
    )
    for i in results:
        print(deserialize(i))
    return f"{len(results)} responses"


if __name__ == "__main__":
    handler(None, None)
