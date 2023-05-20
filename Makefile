build:
	go build -o bin/bot cmd/bot/main.go

run: build
	./bin/bot

fclean:
	rm -rf bin/

re: fclean run