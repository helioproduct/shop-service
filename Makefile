### Переменные
DOCKER_COMPOSE = docker-compose -f tests/integration/docker-compose.integration.yaml

### 1. Запуск приложения через Docker Compose
.PHONY: up
up:
	@echo "Запуск приложения через Docker Compose..."
	@docker-compose up --build

.PHONY: down
down:
	@echo "🛑 Остановка сервисов..."
	@docker-compose down -v

### 3. Перезапуск с пересборкой
.PHONY: restart
restart: down up

.PHONY: integration-test
integration-test:
	@echo "Запуск интеграционных тестов..."
	@$(DOCKER_COMPOSE) up --build --abort-on-container-exit
	@$(DOCKER_COMPOSE) down -v


.PHONY: clean
clean:
	@echo "🧹 Очистка Docker..."
	@docker system prune -af
