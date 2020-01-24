from flask import Flask

import datetime
import random
import time
app = Flask(__name__)


@app.route("/")
def hello():
    print("Paused 1 second", datetime.datetime.now().isoformat())
    time.sleep(1)
    return dict(datetime=str(datetime.datetime.now().isoformat()),
                values=[random.randint(0, 100_000) for _ in range(10)])


if __name__ == "__main__":
    app.run()
