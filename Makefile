TS := $(shell /bin/date "+%Y-%m-%d--%H-%M-%S")
UTIL := @docker-compose -f docker-compose.yml -f docker-compose.util.yml run --rm

composer:
	${UTIL} composer ${C}

setup:
	@docker-compose up -d
	@make composer C=install
	@make migrations C=migrate

gobuild:
	@cd ./services/${S} && wire ./config/ && env GOOS=linux GOARCH=amd64 go build cmd/${S}/main.go && cd -
	@docker-compose stop ${S}-service
	@docker-compose build ${S}-service
	@docker-compose up -d ${S}-service

test:
	@docker-compose exec app ./vendor/bin/phpunit
	@docker-compose exec app ./vendor/bin/behat tests/behavioural/features/orderFulfillment.feature

standards:
	@docker-compose exec app ./vendor/bin/phpcs --standard=PSR12 --ignore=./src/Wallys/Infrastructure ./src

migrations:
	@docker-compose exec app /usr/local/bin/migrations ${C}
