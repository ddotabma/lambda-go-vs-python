from dataclasses import dataclass
import asyncio
import aiohttp
from typing import List


@dataclass
class OperationData:
    datetime: str
    values: List[int]

    def __post_init__(self):
        self.values = [int(i) for i in self.values]  # convert values to int
        assert isinstance(self.datetime, str)


async def get(session, url: str) -> dict:
    response = await session.get(url)  # Make requests
    return await response.json()  # Dictionaries


async def gather_all_get_requests(url_list: List[str]):
    async with aiohttp.ClientSession() as session:
        return await asyncio.gather(*[get(session, i) for i in url_list])


def handler(event, context):
    results = asyncio.run(
        gather_all_get_requests(
            ["https://566pcoo3hl.execute-api.eu-west-1.amazonaws.com/dev" for _ in range(100)] # todo make var
        )
    )
    for i in results:
        print(OperationData(**i))
    return f"{len(results)} responses"


if __name__ == "__main__":
    handler(None, None)
