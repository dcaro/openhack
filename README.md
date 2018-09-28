# Demo application for the DevOps OpenHack

## Run as a container
1- Build the application with multistage container

`docker build . -t dohapp`

2- Run the application from the container

`docker run -p 8080:80 -it dohapp`


## Build and deploy on Kubernetes

1- Login to ACR

`docker login --username <user_name> --password <password> <registry_name>.azurecr.io`

2- Build the image

`docker build . -t <registry_name>.azurecr.io/demo/dohapp:1`

3- Push the image to ACR

`docker push <registry_name>.azurecr.io/demo/dohapp:1`

> Change code in the app (change line 15 and 16 for example).
> Repeat steps 2 and 3 to generate a new versions of the container.

4- Deploy to AKS

`helm install ./dohapp --name dohapp --set image.repository=<registry_name>.azurecr.io/demo/dohapp image.tag=1`

5- Wait until an IP is assigned to the service

`kubectl get svc dohapp -w`

Navigate to the obtained IP to show the website


5- Upgrade to version 2

`helm upgrade dohapp ./dohapp --set image.repository=<registry_name>.azurecr.io/demo/dohapp,image.tag=2`

Refresh the browser to show the updated version


6- Rollback to previous version

`helm rollback dohapp 1`

7- Refresh the browser to show the rollback
