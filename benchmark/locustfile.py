from locust import HttpUser, between, task

class WebUser(HttpUser):
    wait_time = between(0, 2)
    incr = 0

    @task
    def create(self):
        self.client.post("/api/v1/create", json={
            "url": "https://github.com/"+str(self.incr),
            "alias": "",
        })

        self.incr += 1