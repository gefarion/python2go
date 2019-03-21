all: test_go

%.so: %.go
	go build  -buildmode'='c-shared -o $@

test_go.py: test.py
	sed 's/module2migrate/a/' test.py > test_go.py

test_go: a.so test_go.py
	python test_go.py

clean:
	rm -fr *.a *.o *.so *.h test_go.py
