# shop-service


### 1. Запуск:
```bash
make up
```


3. Приложение будет доступно по адресу: `http://localhost:8080`

## API Endpoints

###  **Авторизация:**
####  POST /auth/register — Регистрация пользователя  
- **Auth:** Не требуется  
- **Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
#### POST /auth/login — Авторизация пользователя  
- **Auth:** Не требуется  
- **Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
---
### **Информация:**
#### GET /api/info — Получить информацию о монетах, инвентаре и истории транзакций  
- **Auth:** Требуется (`Bearer Token` в заголовке Authorization)  
- **Пример ответа:**
  ```json
  {
    "coins": 1000,
    "inventory": [{"type": "item", "quantity": 2}],
    "coinHistory": {
      "received": [{"fromUser": "alice", "amount": 50}],
      "sent": [{"toUser": "bob", "amount": 30}]
    }
  }
  ```
---
###  **Переводы:**
#### POST /api/transfer — Отправить монеты другому пользователю  
- **Auth:** Требуется  
- **Body:**
  ```json
  {
    "toUsername": "string",
    "amount": 100
  }
  ```
---
### **Покупки:**
#### GET /api/buy/{item} — Купить предмет за монеты  
- **Auth:** Требуется  
- **Параметры:**
  - `item` (string) — Название предмета  
- **Пример ответа:**
  ```json
  {
    "message": "Item purchased successfully"
    "balance": 790
  }
  ```



## Покрытие тестами
✅ **Unit-тесты**: repository (77%), usecase (44%) (не успел написать тесты с transaction manager mock), handlers (77.6%)  
✅ **Интеграционный тест**: покупка мерча



### Интеграционные тесты:
```bash
make integration-test
```


docker compose  -f ./tests/integration/docker-compose.integration.yaml up -d 

### **Архитектура**  

Проект построен по принципу чистой архитектуры  

#### 📂 **Основные слои:**  
- **Handler:** Обрабатывает входящие HTTP-запросы, вызывает методы usecase и формирует ответы.  
- **UseCase:** Содержит основную бизнес-логику приложения. Здесь используется [Transaction Manager](https://github.com/avito-tech/go-transaction-manager) для управления транзакциями.  
- **Repository:** Отвечает за работу с базой данных, используя [Squirrel](https://github.com/Masterminds/squirrel) для построения SQL-запросов.  
- **Domain:** Определяет основные сущности и бизнес-ошибки.  

#### **Мапперы:**  
Для передачи данных между слоями используются мапперы


