import os
import responder
import httpx

api = responder.API()

@api.route("/")
def index(req, res):
  res.text = "app-a"

@api.route("/debugkit")
def debugkit(req, res):
  r = httpx.get("http://debugkit/info/hostname")
  re.text = r.text

if __name__ == "__main__":
  port = os.getenv("APPA_PORT", "8080")
  api.run(address="0.0.0.0", port=int(port))
