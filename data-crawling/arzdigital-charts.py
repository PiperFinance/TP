import json
import httpx
import asyncio


time_ranges = ["24h", "1w", "1m", "3m", "1y", "all"]
ids = [
    2,
    3,
    4,
    11,
    768885,
    12,
    769858,
    1765,
    30,
    9,
    29,
    7,
    8,
    2237,
    174,
    189,
    768861,
    768861,
    2725,
    31,
    18,
    770187,
    770232,
    1659,
]


async def main():
    async with httpx.AsyncClient() as client:

        reqUrl = "https://api.arzdigital.com/history/"

        headersList = {
            "User-Agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0",
            "Accept": "application/json, text/javascript, */*; q=0.01",
            "Accept-Language": "en-US,en;q=0.5",
            "Accept-Encoding": "gzip, deflate, br",
            "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
            "Origin": "https://arzdigital.com",
            "Connection": "keep-alive",
            "Referer": "https://arzdigital.com/",
            "Sec-Fetch-Dest": "empty",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Site": "same-site",
            "TE": "trailers",
        }
        tasks = []
        resp = {}
        for id in ids:
            for tr in time_ranges:

                async def _l(id, tr):
                    try:
                        payload = f"action=arzajax2&gethistory={id}&dollar=1&range={tr}"
                        data = await client.post(reqUrl, data=payload, headers=headersList)  # type: ignore
                        if id not in resp:
                            resp[id] = {}
                        resp[id][tr] = data.json()
                        with open("res-time.json", "w+") as f:
                            json.dump(resp, f)
                    except Exception as e:
                        print(e)

                tasks.append(_l(id, tr))
        await asyncio.gather(*tasks)
        with open("res-time.json", "w+") as f:
            json.dump(resp, f)


asyncio.run(main())
