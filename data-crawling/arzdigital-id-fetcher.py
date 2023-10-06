import websocket
import json
import rel

f = open("d.txt", "a+")

i = 0

saved = []


def on_message(ws, message):
    f.write(message)
    f.write("\n--------------------\n")
    f.flush()
    try:
        global i
        new = json.loads(message)
        if "data" not in new:
            return
        if new["page"] in saved:
            return
        else:
            saved.append(new["page"])
        with open("arzdigital-coins.json", "r+") as r:
            res = json.loads(r.read() or '{"data":[]}')
            print(i, "::", len(res["data"]), "::", len(new["data"]))
            res["data"] = [*new["data"], *res["data"]]
        with open("arzdigital-coins.json", "w+") as r:
            json.dump(res, r)
        i += 1
        ws.send('{"action":"coins","key":1,"page":' + f'"{i}"}}')
        if i == 255:
            print("DONE")
            exit(0)
    except Exception as e:
        print(e)


def on_error(ws, error):
    print(error)


def on_close(ws, close_status_code, close_msg):
    print("### closed ###")


def on_open(ws):
    print("Opened connection")
    ws.send('{"action":"coins","key":1,"page":' + f'"{i}"}}')


if __name__ == "__main__":
    ws = websocket.WebSocketApp(
        "wss://ws2.arzdigital.com/",
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close,
    )

    ws.run_forever(
        dispatcher=rel, reconnect=5  # type:ignore
    )
    rel.signal(2, rel.abort)  # Keyboard Interrupt
    rel.dispatch()

# ws_app.run_forever()
