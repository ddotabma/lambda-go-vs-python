from dataclasses import dataclass
from typing import List


@dataclass
class Data:
    datetime: str
    values: List[int]

    def __post_init__(self):
        self.values = [int(i) for i in self.values]  # convert values to int
        assert isinstance(self.datetime, str)


def deserialize(d: dict):
    return Data(**d)
