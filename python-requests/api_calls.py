import asyncio
import os

from api_calls_utils.async_requests import multiple_get_requests
from api_calls_utils.create_parquet import dataclasses_to_parquet
from api_calls_utils.model import deserialize


def handler(_, __):
    results = asyncio.run(
        multiple_get_requests(
            [os.environ["API"] for _ in range(100)]  # todo make var
        )
    )

    datas = [deserialize(i) for i in results]
    dataclasses_to_parquet(datas)
    return f"{len(datas)} responses"


if __name__ == "__main__":
    handler(None, None)

    # print(deserialize({"datetime": "2020-01-27T07:07:27.744489", "values": [2210, 13260, 26837, 30913, 46491, 95062, 41528]}))
