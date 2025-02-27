Веб-сервер на Go, который реализует функциональность простейшего планировщика задач. Планировщик хранит задачи, каждая из них содержит дату дедлайна и заголовок с комментарием. Задачи могут повторяться по заданному правилу: например, ежегодно или через какое-то количество дней. В качестве базы данных используется sqlite3.

В директории cmd/service находится main.go

В директории internal/database файлы подключения к db, а также функция проверки существования файла db и в случае отсутствия - создание db и таблицы.

В директории error реализована пользовательская функция обработки сообщения об ошибке в формате JSON.

В директории hendler находятся обработчики приложения для создания/редактирования/получения и удаления задач.

В директории next_date реализована функция NextDate которая вычисляет следующую дату для задачи в соответствии с указанным правилом.

Директория repository содержит структуру репозитория и функцию, котрая создает экземпляр репозитория и отвечает за выполнение запросов к db. А также методы создания таблицы Scheduler и запросы к db (создания/редактирования/получения и удаления задач.)

Файл .env содержит 2 переменные окружения TODO_PORT=port и TODO_DBFILE=name.db

В директории tests находятся тесты для проверки API, которые реализованы в веб-сервере.

Директория web содержит файлы фронтенда.

Для использования приложения go run cmd/service/main.go, в браузере http://localhost:7540/ .

Для запуска всех тестов go test ./tests (порт: 7540), для проверки запросов по отдельности:

проверка веб-сервера: go test -run ^TestApp$ ./tests 

проверка DB: go test -run ^TestDB$ ./tests

проверка даты: go test -run ^TestNextDate$ ./tests

проверка добавления задачи: go test -run ^TestAddTask$ ./tests 

проверка списка ближайших задач go test -run ^TestTasks$ ./tests 

проверка редактирования задачи go test -run ^TestTask$ ./tests 

проверка обновления параметров задачи go test -run ^TestEditTask$ ./tests 

проверка выполнения задания go test -run ^TestDone$ ./tests 

проверка удаления задачи go test -run ^TestDelTask$ ./tests 

