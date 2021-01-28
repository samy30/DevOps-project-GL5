# Projet Devops
## Description du Projet 
L'objectif de ce projet est de se familiariser avec les notions, technologies et outils relatifs au déploiement et Devops. On s'intéressera principalement à : 

- Développer une application de gestion de produits en Golang.
- Créer des Tests unitaires. 
- Déploier l'application sur Kubernetes. 
- Créer un CI/CD pipeline sur Azure Devops. 
- Créer des métriques concernant l'application et les collecter via Prometheus. 
- Visualiser les métriques sur Grafana. 
- Créer des alerts.

## Makefile
Nous disposons d'un Makefile : 
    - Installer les dépendances : 
```sh
$ make install
```
    - Build de l'application : 
```sh
$ make build
```
    - Executer les tests : 
```sh
$ make test
```
    - Execution de l'appolication : 
```sh
$ make run
```
## Endpoints de l'application


| Methodes | API Endpoints | Description | Payload |
| ------ | ------ |------ | ------ |
| GET | "/api/products"| Afficher tous les produits |    -------
| POST | "/api/products"| Ajouter un nouveau produit |{"title": "title_value","price": price_value,"initial_quantity":qt_value,"category": {"name": name_value}}
| PUT | "/api/products/{id}"| Mise à jour d'un produit | {"title": "title_value","price": price_value,"initial_quantity":qt_value,"category": {"name": name_value}}
| DELETE |"/api/products/{id}" | Supprimer un produit | ------
| POST | "/api/products/buy" | Achat d'un produit |{"product_id":"id_value","quantity" :quantity_value}
| GET | "/api/products/transactions" | Afficher les transactions | ------

### Example : 
```sh
(GET) -> localhost:8000/api/products
```
## Métriques 
Afin de visualiser les métriques relatives à l'application, il suffit de taper : 
```sh
(GET) -> localhost:8000/metrics 
```    
Vous allez voir quelques choses de semblable : 
```sh
product_request_duration_milliseconds_bucket{code="200",method="GET",path="/api/products/transactions",service="Product Service",le="5000"} 2
# HELP product_requests_total How many HTTP requests processed, partitioned by status code, method and HTTP path.
# TYPE product_requests_total counter
product_requests_total{code="200",method="GET",path="/api/products",service="Product Service"} 12
``` 
## Monitoring 
Grafana nous permet de visualiser les métriques collectées par Prometheus. Tapez donc : 
```sh
(GET) -> localhost:3000
```    
Se connecter sur Grafana : 
|    Username | Password |
| ------ | ------ |
|admin | admin|



## Passer Au Cloud 

On a utilisé Azure Devops et le Azure Kubernetes Service (AKS)

|    | External IP |    Port | Replicas |
| ------ | ------ |------ | ------ |
| Backend | 52.151.229.9 | 8000 | 2
| Prometheus |52.142.34.131| 9090 | 1
| Grafana |52.149.232.84| 3000| 1