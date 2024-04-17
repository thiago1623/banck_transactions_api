.PHONY: clean coverage docs help \
	quality requirements selfcheck test test-all upgrade validate

.DEFAULT_GOAL := help


help: ## display this help message
	@echo "Please use \`make <target>' where <target> is one of"
	@perl -nle'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'


build: ## build and install development environment requirements inside container
	docker compose up --build

test: clean ## run tests in the current virtualenv
	docker exec -it back bash -c "cd test_from_stori_card/ && python manage.py test"

migrate:
	docker exec -it back bash -c "cd test_from_stori_card/ && python manage.py migrate"

run:
	docker compose up

run-local:
	 go run main.go

run-local-air:
	 air

quality:
	pylint --max-line-length=120 test_from_stori_card/transaction_management_api/api/
	pylint --max-line-length=120 test_from_stori_card/transaction_management_api/utils/

runCLI:
	./stori_card_cli/cli stori_card_cli/transactions_info.csv

qualityDocker:
	docker exec -it back bash -c "pylint --max-line-length=120 test_from_stori_card/transaction_management_api/api/"
	docker exec -it back bash -c "pylint --max-line-length=120 test_from_stori_card/transaction_management_api/utils/"

selfcheck: ## check that the Makefile is well-formed
	@echo "The Makefile is well-formed."
