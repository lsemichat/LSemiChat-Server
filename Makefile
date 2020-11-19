help: ## display help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev-up: ## up develop env
	docker-compose up -d

dev-start: ## start develop env
	docker-compose start

dev-stop: ## stop develop env
	docker-compose stop

dev-down: ## down develop env
	docker-compose down --rmi local --volumes

dev-log-api: ## watch api's log
	docker-compose logs -f api

dev-log-db: ## watch db's log
	docker-compose logs -f db