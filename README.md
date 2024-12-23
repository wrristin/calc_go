# calc_service
Этот проект представляет собой веб-сервис для вычисления арифмитических выражений.
## Как работает этот проект?
Все достаточно просто. Пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.
## Как запустить проект?
Сначала нужно клонировать репозиторий:
```
https://github.com/wrristin/calc_go.git
```
А потом переходим в папку с проектом:
```
cd C:\calc_service
```
А для запуска проекта необходимо в окне терминала ввести одну команду, которая будет прикреплена ниже.
```
go run ./cmd/...
```
## Информация о сервисе
У сервиса 1 endpoint с url-ом /api/v1/calculate. Пользователь отправляет на этот url POST-запрос с телом:
```
{
    "expression": "выражение, которое ввёл пользователь"
}
```
В ответ пользователь получает HTTP-ответ с телом:
```
{
    "result": "результат выражения"
}
```
## Примеры использования, ошибки и ожидаемые ответы
Всего есть 3 исхода завершения программы.
- Правильный запрос - 200(OK)
- Некорректный запрос — 422 (Необрабатываемый объект)
- Некорректный запрос — 500 (Внутренняя ошибка сервера)
- ## ---------------------
Чтобы получить ответ, нужно отправить запрос через Postman или cURL:
- URL: http://localhost:8080/api/v1/calculate
- Метод: POST
- Заголовок: Content-Type: application/json
- И сам пример, который будет приведен ниже.
###### Для этого нужен сайт Postman. Выберите метод POST, введите URL: http://localhost:8080/api/v1/calculate. Перейдите на вкладку Headers, и добавьте новый заголовок Content-Type: application/json. Потом перейдите во вкладку Body, выберите raw и выберите JSON данные.
```
{
    "expression": "2 + 2 * 2"
}
```
В итоге мы получим следующее:
```
{"result":"6.000000"}
```
Если мы попробуем вставить
```
{
    "expression": "2 + a"
}
```
То получим следующее:
```
{"error": "Internal server error"}
```
То есть ошибку 500: Internal Server Error
Или например если вставить вот такое выражение:
```
{
    "expression": "3 / 0"
}
```
То выдаст ошибку 500: Internal server error
