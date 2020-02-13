import asyncio
import os

from .async_requests import send_all_get_requests
from .create_parquet import dataclasses_to_parquet
from .model import deserialize


def handler(event, context):
    results = asyncio.run(
        send_all_get_requests(
            [os.environ["API"] for _ in range(100)]  # todo make var
        )
    )

    datas = [deserialize(i) for i in results]
    dataclasses_to_parquet(datas)
    return f"{len(datas)} responses"


if __name__ == "__main__":
    handler(None, None)

    # print(deserialize({"datetime": "2020-01-27T07:07:27.744489", "values": [2210, 13260, 26837, 30913, 46491, 95062, 41528]}))
