from locust import HttpUser, between, task

class WebUser(HttpUser):
    wait_time = between(0, 2)
    incr = 0

    @task
    def create(self):
        packet = {
            "url": "https://github.com/"+str(self.incr),
            "alias": "",
        }

        with self.client.post("/api/v1/create", json=packet, catch_response=True) as response:
            if response.status_code != 200 and response.status_code != 400:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()

        self.incr += 1
    
    @task
    def redirect(self):
        with self.client.get("/api/v1/jian", allow_redirects=False, catch_response=True) as response:
            if response.status_code != 302:
                response.failure("Got unexpected response code: " + str(response.status_code) + " Error: " + str(response.text))
            else:
                response.success()