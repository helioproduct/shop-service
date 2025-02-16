# shop-service


### 📌 1. Запуск через Docker Compose:
```bash
docker-compose up --build
```


3. Приложение будет доступно по адресу: `http://localhost:8080`

## 📡 Доступные ручки API

### 1️⃣ **Авторизация:**
#### ➡️ `POST /auth/register` — Регистрация пользователя  
- **Auth:** Не требуется  
- **Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
#### ➡️ `POST /auth/login` — Авторизация пользователя  
- **Auth:** Не требуется  
- **Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
---
### 2️⃣ **Информация:**
#### ➡️ `GET /api/info` — Получить информацию о монетах, инвентаре и истории транзакций  
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
### 3️⃣ **Переводы:**
#### ➡️ `POST /api/transfer` — Отправить монеты другому пользователю  
- **Auth:** Требуется  
- **Body:**
  ```json
  {
    "toUsername": "string",
    "amount": 100
  }
  ```
---
### 4️⃣ **Покупки:**
#### ➡️ `GET /api/buy/{item}` — Купить предмет за монеты  
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



## 🧪 Покрытие тестами
✅ **Unit-тесты**: repository, usecase, handlers  
✅ **Интеграционные тесты**: ...
