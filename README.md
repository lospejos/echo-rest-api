# echo-rest-api
В данном проекте показан пример создания REST API на Go с использованием Echo framework.

### В проекте использованы
#### HTTP
- `github.com/labstack/echo` - для создания rest сервиса; использованы: _router_, _data binding_ и _data rendering_, _logger middleware_

#### Валидация
- `gopkg.in/go-playground/validator.v9` - для _data validation_  

#### База данных
- `database/sql` - для работы с запросами и транзакционностью
- `github.com/lib/pq` - в качестве СУБД использовалась postgresql
- `github.com/rubenv/sql-migrate` - миграционная тулза для sql
- `docker postgres image` - для запуска postgresql
#### Конфигурация
- `flag` - для передачи параметров при запуске
- `github.com/jinzhu/configor` - для загрузки конфигурации из yaml

#### Логгирование
- `github.com/sirupsen/logrus` - для логов приложения 

#### Тестирование
- `testing` - для написания тестов
- `github.com/golang/mock/gomock` - для генерации моков
- `github.com/stretchr/testify/assert` - исключительно для _assert_ в тестах

#### Зависимости
`github.com/golang/dep` - менеджер зависимостей

### Для запуска
- `make init` - установит необходимые тулзы (dep, sql-migrate, mockgen), затем через `dep` установит зависимости проекта
- `make run`  - запустит приложение, при этом запустив docker контейнер с БД
- `make test` - запустит приложение, при этом запустив docker контейнер с БД
- `make spec` - сгенерирует open-api спецификацию в _spec/api.json_.
После запуска сервис отдает спеку как статику по пути _/spec/api.json_, можно также выполнить `swagger serve -F=swagger http://localhost:8081/spec/api.json`, что бы отобразить в html виде c запущенного сервиса.
- `spec-ui` - откроет спеку в html

### Подключение к контейнеру echo-rest-api-db:
`docker exec -i -t echo-rest-api-db /bin/sh`

