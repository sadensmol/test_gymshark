.PHONY up:
up:
	go run github.com/cosmtrek/air@latest

.PHONY test:
test:
	go test -v ./...


.PHONY:build
build:
	export GOOS=linux && export GOARCH=amd64 && mkdir -p bin && go build -o ./bin/test_gymshark ./cmd/...

.PHONY:deploy
deploy: build
	scp -r ./bin/test_gymshark root@159.223.212.53:~/test_gymshark_new
	scp -rp ./templates root@159.223.212.53:~
	scp -rp ./infrastructure/prod/. root@159.223.212.53:~
	ssh root@159.223.212.53 "chmod +x start.sh && nohup ./start.sh &"
	