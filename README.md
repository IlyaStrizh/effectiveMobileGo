<p align="center">
  <a href="https://pkg.go.dev/github.com/gorilla/mux"><img src="https://img.shields.io/badge/Gorilla-Mux-blue" alt="Gorilla" /></a>
  <a href='https://github.com/IlyaStrizh/effectiveMobileGo/blob/main/migrations/000001_create_cars_table.up.sql'><img src='https://img.shields.io/badge/migrations-PostgreSQL-blue' alt='PostgreSQL' /></a>
  <a href='https://github.com/IlyaStrizh/effectiveMobileGo/blob/main/docs/swagger.yaml'><img src='https://img.shields.io/badge/Swagger-2.0-green.svg' alt='Swagger' /></a>	
</p>

# Реализован каталог автомобилей
### 1. Выставлены rest методы
- 1. Получение данных с фильтрацией по всем полям и пагинацией 
- 2. Удаления по идентификатору
- 3. Изменение одного или нескольких полей по идентификатору
- 4. Добавления новых автомобилей в формате
```json
{
    "regNums": ["X123XX150"] // массив гос. номеров
}
```
### 2. При добавлении выполняется запрос во внешнее АПИ, описанного сваггером (это описание некоторого внешнего АПИ, которого нет, но к которому надо обращаться.)

```yaml
openapi: 3.0.3
info:
  title: Car info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: regNum
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Car'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    Car:
      required:
        - regNum
        - mark
        - model
        - owner
      type: object
      properties:
        regNum:
          type: string
          example: X123XX150
        mark:
          type: string
          example: Lada
        model:
          type: string
          example: Vesta
        year:
          type: integer
          example: 2002
        owner:
          $ref: '#/components/schemas/People'
    People:
      required:
        - name
        - surname
      type: object
      properties:
        name:
          type: string
        surname:
          type: string
        patronymic:
          type: string
```
### 3. Обогащенная информация сохраняется в БД postgres (структура БД создается путем миграций при старте сервиса)
### 4. Код покрыт debug- и info-логами
### 5. Конфигурационные данные вынесены в .env-файл
### 6. Сгенерирован сваггер на реализованное АПИ
