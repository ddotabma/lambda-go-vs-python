from dataclasses import dataclass
import asyncio
import aiohttp
from typing import List
import os
import pyarrow as pa



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

    datas = [deserialize(i) for i in results]
    dataclasses_to_parquet(datas)
    return f"{len(datas)} responses"



def dataclasses_to_parquet(dcs: List[Data]):
    timestamps = [i.datetime for i in dcs]
    timestamps_array = pa.array(timestamps, type=pa.string())

    values = [i.values for i in dcs]
    values_array = pa.array(values, type=pa.list_(pa.int32()))
    table = pa.Table.from_arrays(arrays=[timestamps_array, values_array], names=["timestamp", "values"])
    table.to_pandas().to_parquet("s3://bdr-go-blog/dump/python.parquet")

if __name__ == "__main__":
    handler(None, None)

    # print(deserialize({"datetime": "2020-01-27T07:07:27.744489", "values": [2210, 13260, 26837, 30913, 46491, 95062, 41528]}))
