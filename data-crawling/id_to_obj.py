import json
from tqdm import tqdm


with open("./arzdigital-coins.json") as f:
    coins = json.load(f).get("data", {})

with open("./all_tokens.json") as f:
    all_tokens = json.load(f)

# matches = {}
# for coin in tqdm(coins):
#     matches[coin["id"]] = set()
#     for token_id, token in all_tokens.items():
#         det = token["detail"]
#         matched = False
#         for d_k, d_v in det.items():
#             if matched:
#                 break
#             for c_k, c_v in coin.items():
#                 if matched:
#                     break
#                 if c_v is not None and c_v == d_v:
#                     matches[coin["id"]].add(token_id)
#                     matched = True
#     matches[coin["id"]] = list(matches[coin["id"]])

matches = {}
for coin in tqdm(coins):
    matches[coin["id"]] = list(
        {
            tokenId
            for tokenId, token in all_tokens.items()
            for c_v in coin.values()
            if c_v is not None
            for d_v in token.get("detail", {}).values()
            if c_v == d_v
        }
    )
with open("res.json", "w") as f:
    json.dump(matches, f)
