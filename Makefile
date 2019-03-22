all: test_a test_b

srv: srv_go

srv_go: srv_go.py c.so
	python srv_go.py

%.so: %.go
	go build  -buildmode'='c-shared -o $@ $*.go

test_a.py: test.py
	sed 's/module2migrate/a/' test.py > test_a.py

test_b.py: test.py module2migrate.py
	sed 's/entities/b_go/' module2migrate.py > module2migrate_go.py
	sed 's/entities/b_go/;s/module2migrate/&_go/' test.py > test_b.py

test_%: %.so test_%.py
	python test_$*.py

clean:
	rm -fr *.a *.o *.so *.h test_*.py module2migrate_go*.py *.pyc

req:
	wrk -s req.lua -c 1 -t 1 -d 1s http://0.0.0.0:8080/req

req-kill:
	wrk -s req.lua -c 100 -t 50 -d 2m http://0.0.0.0:8080/req
