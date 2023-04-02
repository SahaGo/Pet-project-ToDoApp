# Создание веб-приложения по созданию ToDo списков на Golang, с помощью архитектуры REST-API

### Данный пет-проект написан по [видеокурсу](https://www.youtube.com/playlist?list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8).

В 1 уроке дано базовое представление чем будем заниматься и какие интрументы будут использоваться.

Во 2 уроке составлена структура приложения, подключены библиотеки автора и GIN, функция запуска веб-сервера,добавлены хендлеры и эндпоиты, структуры todo, user.
В процессе компиляции возникла ошибка `unix/syscall_darwin.1_13.go:25:3: //go:linkname must refer to declared function or variable`, удалось ее [исправить](https://stackoverflow.com/questions/71507321/go-1-18-build-error-on-mac-unix-syscall-darwin-1-13-go253-golinkname-mus).