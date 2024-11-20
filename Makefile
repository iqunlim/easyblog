run:
	@go run app/main.go

enter:
	@docker exec -w /go/src/github.com/iqunlim/easyblog/ -it easyblog-dev bash

dc:
	@docker compose up -d

dd:
	@docker compose down