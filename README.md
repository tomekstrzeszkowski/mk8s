## Docker
`docker-compose up --build`

### Migration

`docker-compose run --rm web python manage.py migrate`

## Kubernetes

- build image `docker build -t mk8sdjango_web:v<bumped_version:int> .`
- Edit `kube/deployment.yml` image and put the new version
- save local image
```docker save mk8sdjango_web:v<bumped_version:int> > mk8sdjango_web.tar```
- Import image into microk8s registry
`microk8s ctr image import mk8sdjango_web.tar`
- Optional: verify the entry `microk8s ctr images ls | rg django`
 - Apply deployment 
`microk8s kubectl apply -f kube/deployment.yml`
 - If it's the first deploy apply service `microk8s kubectl apply -f kube/service.yml`

### Managing kubernetes
Run the dashboard `microk8s.dashboard-proxy`
