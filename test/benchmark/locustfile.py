import random
from locust import HttpUser, between, task

create_incr = 0
redirect_incr = 0

class create_tinyurl_using_increasing_number(HttpUser):
    wait_time = between(0, 2)

    @task
    def create(self):
        global create_incr
        packet = {
            "url": "https://github.com/"+str(create_incr),
            "alias": "",
        }

        with self.client.post("/api/v1/create", json=packet, catch_response=True) as response:
            if response.status_code != 200 and response.status_code != 400:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()
                create_incr += 1

class create_tinyurl_using_same_resource(HttpUser):
    wait_time = between(0, 2)

    @task
    def create(self):
        packet = {
            "url": "https://github.com/JianLiu666",
            "alias": "",
        }

        with self.client.post("/api/v1/create", json=packet, catch_response=True) as response:
            if response.status_code != 200 and response.status_code != 400:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()

class redirect_using_increasing_number(HttpUser):
    wait_time = between(0, 2)

    @task
    def redirect(self):
        global redirect_incr

        with self.client.get("/api/v1/"+str(redirect_incr), allow_redirects=False, catch_response=True) as response:
            if response.status_code != 400:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()
                redirect_incr += 1

class redirect_using_same_tinyurl(HttpUser):
    wait_time = between(0, 2)

    @task
    def redirect(self):
        with self.client.get("/api/v1/xaJxi", allow_redirects=False, catch_response=True) as response:
            if response.status_code != 302:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()

class randomly_case(HttpUser):
    wait_time = between(0, 1)
    tinyurls = ["jian"]

    @task
    def create(self):
        global create_incr
        packet = {
            "url": "https://github.com/"+str(create_incr),
            "alias": "",
        }

        with self.client.post("/api/v1/create", json=packet, catch_response=True) as response:
            if response.status_code != 200 and response.status_code != 400:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()
                create_incr += 1
                # 避免無止盡增長
                if response.status_code == 200 and len(self.tinyurls) < 1000000:
                    self.tinyurls.append(response.json()['tiny'].split("/")[-1])

    @task(100)
    def redirect(self):
        with self.client.get("/api/v1/"+random.choice(self.tinyurls), allow_redirects=False, catch_response=True) as response:
            if response.status_code != 302:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()