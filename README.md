# Shortbread Parking

```
docker-compose up
docker-compose run web buffalo db migrate
docker-compose run web buffalo task db:seed
```

## Deployment

```
heroku container:push --recursive web
heroku container:release web
```
