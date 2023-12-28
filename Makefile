compose_up:
	docker compose --project-name hex -f ./deployments/docker-compose.yaml up --build -d

compose_down:
	docker compose --project-name hex -f ./deployments/docker-compose.yaml down

test-commit:
	pre-commit run --all-files

setup:
	brew install pre-commit yamllint
