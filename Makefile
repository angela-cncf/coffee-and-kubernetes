#####################################
# Play with the Hello K8s app locally
.PHONY: build-app
build-app:
	make -C app build
.PHONY: run-app
run-app:
	make -C app run
.PHONY: clean-app
clean-app:
	- make -C app clean

#######################################
# Play with the Hello K8s app in Docker
.PHONY: docker-container
docker-container:
	docker build --tag hello:docker --file docker/app.Dockerfile ./app
.PHONY: docker-run
docker-run: docker-container
	docker run --name hello-docker -p 30080:8080 hello:docker /go/bin/hello --config /go/bin/configuration.json
docker-run-again: 
	docker run --name hello-docker-again -p 30081:8080 hello:docker /go/bin/hello --config /go/bin/configuration.json
.PHONY: docker-clean
docker-clean:
	- docker stop hello-docker
	- docker stop hello-docker-again
	- docker rmi hello:docker -f

###############
# Play with K8s
.PHONY: k8s-container
k8s-container:
	docker build --tag hello:k8s --file docker/app.Dockerfile ./app
.PHONY: db
db:
	kubectl apply -f k8s/db.yaml
.PHONY: app
app: k8s-container db
	kubectl apply -f k8s/app.yaml

.PHONY: clean
clean:
	- kubectl delete -f k8s/app.yaml
	- kubectl delete -f k8s/db.yaml
	- docker rmi hello:k8s -f

#################
# Play with chaos
.PHONY: chaos
chaos:
	docker build --tag chaos-monkey:latest --file docker/chaos.Dockerfile k8s/
	kubectl apply -f k8s/chaos.yaml
.PHONY: order
order:
	- kubectl delete -f k8s/chaos.yaml
	- docker rmi chaos-monkey:latest -f