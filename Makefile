run:
	go run etc/main.go
plg:
	go build -o bin/demo.so -buildmode=plugin demo/plugin.go

clean:
	rm bin/demo.so
