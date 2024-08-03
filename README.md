# Docker
`docker compose up --build`

# Migration for django app

`docker compose run web python manage.py migrate`

# Kubernetes deployement 
1. build image `docker build -t <img_name>:v<bumped_version:int> .`
2. Edit `./kube/<name>-deployment.yml` image and put the new version
3. save local image
```docker save <img_name>:v<bumped_version:int> > <img_name>.tar```
4. Import image into microk8s registry
`microk8s ctr image import <img_name>.tar`
- Optional: verify the entry `microk8s ctr images ls | rg <img_name>`
5. Apply deployment 
`microk8s kubectl apply -f ./kube/<name>-deployment.yml`
 - If it's the first deploy apply service `microk8s kubectl apply -f kube/<name>-service.yml`

## Example for django app

1. `docker build -t mk8sdjango_web:v<bumped_version:int> .`
2. Edit `./kube/django-deployment.yml` 
3. 
```docker save mk8sdjango_web:v<bumped_version:int> > mk8sdjango_web.tar```
4. 
`microk8s ctr image import mk8sdjango_web.tar`
- `microk8s ctr images ls | rg django`
5. Apply deployment 
`microk8s kubectl apply -f ./kube/django-deployment.yml`

 # Kubernetes deployment go `mk8sdjango/gin`
 The process is like in django app just use: 
  - `<img_name>` -> `gin_web`
  - `<name>` -> `gin`

* note: endpoint for `content_tags` is not supported yet, due to missing db connection.

# Managing kubernetes
Run the dashboard `microk8s.dashboard-proxy`

Running app:
![Screenshot from 2023-10-29 12-20-40](https://github.com/tomekstrzeszkowski/mk8sdjango/assets/40120335/2ce1402d-1dea-45af-b0b8-aad6fb5b38bf)

dashboard:
![Screenshot from 2023-10-29 12-21-48](https://github.com/tomekstrzeszkowski/mk8sdjango/assets/40120335/ea28fa4a-8bc0-4f30-904d-8628aa8605c2)


# Known issues

When running Django app you might get 502 error. To solve that issue restart deployment
 - `microk8s kubectl rollout restart -n backend-django deployment mk8sdjango`
or delete pods
 - `microk8s kubectl delete -n backend-django pod mk8sdjango-[django-pod-id]` 


# Running new services

## Postgres
Run `microk8s kubectl apply -f ./kube/postgres-[name].yaml` in this order:
 1. `volume`
 2. `volume-claim`
then run the rest in any order
 - `configmap`
 - `deployment`
 - `service`

## Django App
 - Edit `.env-deploy` file and put PG's IP
 - build image and deploy new version