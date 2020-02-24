# backend core 

#### Env config
Make sure you export this variables in your .env
```
DB_HOST=127.0.0.1
DB_DRIVER=postgres
API_SECRET=98hbun98h
DB_USER=user
DB_PASSWORD=pass
DB_NAME=core
DB_PORT=5432

PGADMIN_DEFAULT_EMAIL=email@gmail.com
PGADMIN_DEFAULT_PASSWORD=password
```

#### Docker Compose
Starting services
``` 
~ docker-compose up
~ docker-compose up -d
```
Stopping services
``` 
~ docker-compose rm
~ docker-compose down
~ docker-compose kill
```
Rebuilding
``` 
~ docker-compose rm
~ docker-compose down
~ docker rmi back_core_api

~ docker-compose up
```

Final
``` 
pgadmin     ->  http://localhost:5050
api         ->  http://localhost:420
```



