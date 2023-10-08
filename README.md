## Docker
`docker-compose up --build`

## Kubernetes

- save local image
```docker save mk8sdjango-web > mk8sdjango-web.tar```
- Import image into microk8s registry
`microk8s ctr image import mk8sdjango-web.tar`
- Optional: verify the entry `microk8s ctr images ls | rg django`
 - Apply deployment 
`microk8s kubectl apply -f kube/deployment.yml`
 - Apply service `microk8s kubectl apply -f kube/service.yml`

## Managing kubernetes
Run the dashboard `microk8s.dashboard-proxy`
