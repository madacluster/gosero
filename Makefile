build:
	env GOOS=linux GOARCH=arm64 go build

run:
	go run main.go

deploy:
	npm run build --prefix frontend
	cp -r  frontend/public  pkg/webserver/
	env GOOS=linux GOARCH=arm64 go build
	scp gosero ubuntu@${RASPBERRY_IP}:~/
