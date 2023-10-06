import json

with open("res-time.json") as f:
    timestamps = json.load(f)

with open("res.json") as f:
    coins_detail = json.load(f)


level_map = {
    "24h": "D",
    "1m": "M",
    "1w": "W",
    "1y": "Y",
}
id_map = {
    "3": [
        "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee-1",
        "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee-10",
        "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee-138",
        "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee-288",
        "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee-288",
    ]
}
