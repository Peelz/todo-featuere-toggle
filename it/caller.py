import time

import requests

if __name__ == '__main__':
    # for i in range(10):
    # print("call with dev")
    # res = requests.post("http://localhost:8080/todo", headers={'X-User-ID': 'dev'})
    # print(res.json())

    # for i in range(10):
    # print("call with qa")
    # res = requests.post("http://localhost:8080/todo", headers={'X-User-ID': 'qa'})
    # print(res.json())

    for i in range(1):
        time.sleep(0.3)
        print("######")
        # try:
        #     res = requests.get("http://localhost:8080/random?user=agentx")
        #     print("agentx", res.json())
        # except Exception as e:
        #     print(e)

        try:
            res = requests.get("http://localhost:8080/random?user=eddy")
            print("eddy", res.json())
        except Exception as e:
            print(e)

        # try:
        #     res = requests.get("http://localhost:8080/random?user=tony")
        #     print("tony", res.json())
        # except Exception as e:
        #     print(e)

        print("######")
        # time.sleep(0.3)
        # res2 = requests.get("http://localhost:8080/random?user=foo")
        # print("foo", res2.json())