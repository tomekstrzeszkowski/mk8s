# Before first run
 - Enable ingress addon which will add nginx ingress controller
 ```
microk8s enable ingress
 ```
 - Create namespaces
 ```
 microk8s kubectl apply -f ./kube/namespaces.yml
 ```
 - Create all services, ingress, secret and other items
 
# New app version deployement 
1. For simplicity create variable for version `export version=v<bumped_version:int> name=<name:gin|django>`
2. Build image, for gin use `./gin/.`
```
docker build -t ${name}_web:$version .
```
3. Edit `./kube/${name}-deployment.yml` image and put the new version
4. Save local image
```
docker save ${name}_web:$version > image_tmp.tar
```
* for django Edit `./kub/nginx-deployment.yml` as well
5. Import image into microk8s registry
```
microk8s ctr image import image_tmp.tar
```
- Optional: verify the entry `microk8s ctr images ls | rg $name`
6. Apply deployment 
```
microk8s kubectl apply -f ./kube/${name}-deployment.yml
```

## Deployment variables for django app
 - `export name=django version=v<bumped_version>`

## Deployment variables for gin
 - `export name=gin version=v<bumped_version>`

* note: endpoint for `content_tags` is not supported yet, due to missing db connection.

# Managing kubernetes
Run the dashboard `microk8s.dashboard-proxy`

# Running app:

## Known issue

Sometime when pods are starting you might see django connection issues and nginx error message.
To solve that issue restart pods:

```
kubectl rollout restart -n backend-django deployment mk8sdjango
```

![Screenshot from 2023-10-29 12-20-40](https://github.com/tomekstrzeszkowski/mk8sdjango/assets/40120335/2ce1402d-1dea-45af-b0b8-aad6fb5b38bf)

dashboard:
![Screenshot from 2023-10-29 12-21-48](https://github.com/tomekstrzeszkowski/mk8sdjango/assets/40120335/ea28fa4a-8bc0-4f30-904d-8628aa8605c2)

# Running new services

## Postgres
Run `microk8s kubectl apply -f ./kube/postgres-<name>.yml`:
 - `secrets`
 - `volume`
 - `configmap`
 - `deployment`
 - `service`
