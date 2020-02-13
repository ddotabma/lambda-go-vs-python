import pyarrow as pa
from typing import List

from .model import Data


def dataclasses_to_parquet(dcs: List[Data]):
    timestamps = [i.datetime for i in dcs]
    timestamps_array = pa.array(timestamps, type=pa.string())

    values = [i.values for i in dcs]
    values_array = pa.array(values, type=pa.list_(pa.int32()))
    table = pa.Table.from_arrays(arrays=[timestamps_array, values_array],
                                 names=["timestamp", "values"])
    table.to_pandas().to_parquet("s3://bdr-go-blog/dump/python.parquet")
