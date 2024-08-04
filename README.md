# Kubernetes deployement 
1. For simplicity create variable for version `export version=v<bumped_version:int> name=<name:str>`
1. Build image 
```
docker build -t $name:$version .
```
2. Edit `./kube/$name-deployment.yml` image and put the new version
3. Save local image
```
docker save $name:$version > image.tar
```
4. Import image into microk8s registry
```
microk8s ctr image import image.tar
```
- Optional: verify the entry `microk8s ctr images ls | rg $name`
5. Apply deployment 
```
microk8s kubectl apply -f ./kube/$name-deployment.yml
```
 - If it's the first deploy apply service `microk8s kubectl apply -f kube/$name-service.yml`

## Deployment variables for django app
 - `export name=django version=v<bumped_version>`

## Deployment variables for gin
 - `export name=gin version=v<bumped_version>`

* note: endpoint for `content_tags` is not supported yet, due to missing db connection.

# Managing kubernetes
Run the dashboard `microk8s.dashboard-proxy`

Running app:
![Screenshot from 2023-10-29 12-20-40](https://github.com/tomekstrzeszkowski/mk8sdjango/assets/40120335/2ce1402d-1dea-45af-b0b8-aad6fb5b38bf)

dashboard:
![Screenshot from 2023-10-29 12-21-48](https://github.com/tomekstrzeszkowski/mk8sdjango/assets/40120335/ea28fa4a-8bc0-4f30-904d-8628aa8605c2)

# Running new services

## Postgres
Run `microk8s kubectl apply -f ./kube/postgres-<name>.yaml`:
 - `secrets`
 - `volume`
 - `configmap`
 - `deployment`
 - `service`
