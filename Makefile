TS := $(shell /bin/date "+%Y-%m-%d--%H-%M-%S")
UTIL := @docker-compose -f docker-compose.yml -f docker-compose.util.yml run --rm

composer:
	${UTIL} composer ${C}

setup:
	@docker-compose up -d
	@make composer C=install
	@make migrations C=migrate

test:
	@docker-compose exec app ./vendor/bin/phpunit
	@docker-compose exec app ./vendor/bin/behat tests/behavioural/features/orderFulfillment.feature

standards:
	@docker-compose exec app ./vendor/bin/phpcs --standard=PSR12 --ignore=./src/Wallys/Infrastructure ./src

migrations:
	@docker-compose exec app /usr/local/bin/migrations ${C}
