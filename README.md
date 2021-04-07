# Літературна карта Полтавщини

## [Відкрити застосунок](https://pltv.herokuapp.com/home/)
---
## Інструкція по встановленню на власному сервері (на прикладі Heroku)
1. Перш за все потрібно завантажити клієнт Heroku (https://devcenter.heroku.com/articles/heroku-cli)
2. Далі потрібно створити новий застосунок (пропустити цей пункт якщо застосунок уже створено)
```console
heroku create <ім'я застосунку>
```
3. Далі потрібно авторизуватися в Docker-репозиторії Heroku
```console
heroku container:login
```
4. Далі потрібно помітити Docker контейнер
```console
docker tag <image id> registry.heroku.com/<app name>/web
```
5. після цього потрібно відправити контейнер на сервер Heroku
```console
docker push registry.heroku.com/<app name>/web
```
6. після того як образ успішно відправлено
```console
heroku container:release web --app <app name>
```

