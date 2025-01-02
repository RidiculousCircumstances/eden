wire:
	wire ./shared/wire
build:
	docker compose up --build -d
down:
	docker compose down
