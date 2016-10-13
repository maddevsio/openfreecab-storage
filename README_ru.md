# Хранилище данных для Openfreecabs.org 

Потокобезопасное хранилище геоданных в оперативной памяти для проекта https://openfreecabs.org с R деревом для геоиндеса с минимальным хттп апи

## Предварительные требования

1. [Go](https://golang.org/)
2. [Make](https://www.gnu.org/software/make/)

## Установка
с помощью make
```
mkdir -p $GOPATH/src/github.com/maddevsio/
cd $GOPATH/src/github.com/maddevsio
git clone https://github.com/maddevsio/openfreecab-storage
cd openfreecab-storage
make depends
make
```

Или по пути Go

```
mkdir -p $GOPATH/src/github.com/maddevsio/
cd $GOPATH/src/github.com/maddevsio
git clone https://github.com/maddevsio/openfreecab-storage
cd openfreecab-storage
go get -v
go build -v
go install
```

## Конфигурация

```
GLOBAL OPTIONS:
   --http_bind_addr value  Define custom http port to bind to (default: ":8090") [$HTTP_BIND_ADDR]
   --base_url value        Define custom base url for project (default: "http://localhost:8090") [$BASE_URL]
   --loglevel value        set log level (default: "debug") [$LOG_LEVEL]
   --test_mode             set test mode [$TEST_MODE]
   --help, -h              show help
   --version, -v           print the version
```

## Запуск

```
./openfreecab-storage
```

## HTTP API

Есть несколько методов в HTTP API

1. /nearest/:lat/:lon - Возвращает к ближайших водителей для вашего местоположения
2. /add/ - Используется для добавления данных краулером
3. /clean/:companyName/ - Очищает хранилище от данных той или иной компании.




## Участие в проекте
Вы можете без препятственно создавать таски, слать пулл-реквесты Порядок действий такой:

1. Форкните репозиторий
2. Сделайте ваши изменения
3. Сделайте Commit и Push
4. Отправьте pull request
