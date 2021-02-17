# Troubleshooting Runbook

## Creating a fresh docker image

If creating a fresh docker image, then the init.sh script will need to run. 

Note that, you have to set the script as executable because Docker copies over permissions.

If, looking at the `postgis` container logs (`docker-compose logs --follow postgis`), you see see the error:
```
postgis    | /usr/local/bin/docker-entrypoint.sh: running /docker-entrypoint-initdb.d/init.sh
postgis    | /usr/local/bin/docker-entrypoint.sh: /docker-entrypoint-initdb.d/init.sh: /bin/bash: bad interpreter: Permission denied
```

then make the initialization script executable by running, 
```
chmod +x migration/bin/init/init.sh
```

## Removing an existing contain
If intending to start fresh; and so removing a container. It is important to also remove the existing postres data volume. Otherwise, this intialization script will not run; and thus flyway will not be able to successfully initialize the table. 

In this case, find the volume by running 
```
docker volume ls -q
```

Look for 
```
breadcrumbs_pg_data
```

If it exists, delete it by running 
```
docker volume rm breadcrumbs_pg_data
```

If you get the error: 
```
Error response from daemon: remove breadcrumbs_pg_data: volume is in use - [d473cd3b40293fc1023f1108f2b09118a59df03950178fa24e1f55c636246da8]
```

Then remove the indicated container. In this case: 
```
docker rm d473cd3b4029
```


Note that this removes the DB data; so be careful. 