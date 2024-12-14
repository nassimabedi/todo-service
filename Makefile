run:
	docker-compose up --build

migrate:
	docker-compose exec db psql -U user -d todo_service -f /db/migrations/001_create_todo_items_table.sql

