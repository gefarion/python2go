-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header

io.input("req.json")

wrk.method = "POST"
wrk.body   = io.read("*all")
wrk.headers["Content-Type"] = "application/json"
