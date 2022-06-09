##Docker-compose.yml
Перед запуском контейнеров почистить папки вывода docs, src, client, server

- protoc-gen-doc-md Генерация документации markdown. Вывод в папку docs
- protoc-gen-doc-html Генерация документации в статическом html. Вывод в папку docs
- mkdocs Сервер с web-документацией, генерация из markdown, экспозиция на 8080 порт
- grpc-gateway-swagger-ui Сервер с интерактивной документацией swagger. Генерация swagger спецификации. Вывод в папку docs. Экспозиция на 3000
- protoc-all-go Компиляция protobuf описания в классы Golang. Вывод в папку src
- protoc-all-csharp Компиляция protobuf описания в классы C#. Вывод в папку src
- docker-go-swagger-client Генерация клиента Golang по swagger спецификации. Вывод в папку client
- docker-go-swagger-server Генерация сервера Golang по swagger спецификации. Вывод в папку server