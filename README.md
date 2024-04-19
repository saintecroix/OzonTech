# OzonTech
### Инструкция к использованию
Команда запуска: `docker compose up -d --build`  
Параметр `memory` может быть только `inmemory` или `postgres`
Пример тела post-запроса `http://127.0.0.1:4000/post`:
```json
{
  "memory": "postgres",
  "data": "https://www.google.com/"
}
```
Пример тела ответа:
```json
{
  "url": "https://www.google.com/",
  "short-url": "WqiyWolLpL"
}
```
Пример телa get-запроса `http://127.0.0.1:4000/get`:
```json
{
  "memory": "postgres",
  "data": "WqiyWolLpL"
}
```
Пример тела ответа:
```json
{
  "url": "https://www.google.com/",
  "short-url": "WqiyWolLpL"
}
```
