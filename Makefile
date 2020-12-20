build:
	go build .

test:
	@make build
	bash test.sh

clean:
	@rm ltime
