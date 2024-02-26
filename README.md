# Keysight Day 2024 - Coffee & Kubernetes Workshop

This is a small Go application that is able to run in a variety of environments. Its purpose is to showcase how the behavior of the exact same codebase evolves depending on how it is packaged and deployed.

## Required tools

To get the most out of the repository you'll want a text editor of some sort. The demos from the workshop use [`VSCode`](https://code.visualstudio.com/download) because the presenter likes it. You're perfectly fine to use any tool you're familiar with, from `Notepad` to `TextEdit` to `Emacs` or `vim`.

The demos will make liberal use of [`make`](https://www.gnu.org/software/make/), a timeless tool which you can get even [for Windows](https://gnuwin32.sourceforge.net/packages/make.htm).

The app exposes a REST API and nothing else. As a result of this, interacting with it will be via REST API calls. You can use a browser, [`curl`](https://curl.se/download.html) or the [REST Client VSCode Extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) and the handy examples included [here](test.http).

### Running the app locally

You can build and run the app locally on **any** platform as long as you install [Go](https://go.dev/dl/) on it.

### Running the containerized app

You can build and run the app as a container with **any** runtime that conforms to the [Container Runtime Interface](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md) The demo uses [`Rancher for Desktop`](https://github.com/rancher-sandbox/rancher-desktop/releases/tag/v1.9.1) which provides both `Docker` and `K8s` in an easy to use package. If you're new to this follow the yellow brick road. 

>Note: On Windows, [`Rancher for Desktop`](https://github.com/rancher-sandbox/rancher-desktop/releases/tag/v1.9.1) works best with the [Windows Subsystem for Linux](https://learn.microsoft.com/en-us/windows/wsl/install). Make sure that WSL is enabled before getting Rancher and do the recommended OS restart(s) when prompted.


### Running the K8s app

To play with the app in `Kubernetes` you'll clearly need `Kuebernetes`. The demo uses [`Rancher for Desktop`](https://github.com/rancher-sandbox/rancher-desktop/releases/tag/v1.9.1) which provides both `Docker` and `K8s` in an easy to use package. Stick to this recommendation if it's your first time playing with these technologies.

To interact with K8s, the standard option is to use the [`kubectl`](https://kubernetes.io/docs/tasks/tools/#kubectl) command. You'll need to install it and make sure it points to Rancher's cluster before you can use the `make` helpers related to K8s.

### Extras

A nice CLI tool for interacting with K8s is [`K9s`](https://github.com/derailed/k9s/releases). This is a primary tool for the K8s demos but remember that it is just a nice wrapper fro `kubectl` and doesn't provide additional functionality.

As a final note, keep in mind that K8s and containers are primarily Linux-based environments. It's a good idea to get friendly with the commandline, scripting and Linux OS internals. The more familiar you are with these topics the easier it will be for you to understand what is going on and to dive deeper into the concepts.

## How to use

The app exposes several relevant REST endpoints 
- `/hello` will respond with a nice message that identifies the caller and increments a visitor count
- `/howdy` will respond with a different message and won't increment the visitor count
- `/attack` will crash the app

### Classic app

You can run the application locally with `make run-app`. It will start listening on port 8080 and you can play with it by making calls to [`http://localhost:8080`](http://localhost:8080).

Take a look at the [Makefile](Makefile) to see how the application is started. Try running multiple copies of the application, crashing and restarting it. Do the results match your expectations? Think about what you would do to improve the behavior.

### Containerized app

>Keep the [`docker` cheatsheet](https://docs.docker.com/get-started/docker_cheatsheet.pdf) handy for this section.

You can run the application as a container with `make docker-run`. The app included in the container listens by default on port 8080, just like its local version. A port forward is needed to expose the port from the container outside the container runtime. The associated `make` target will forward container port 8080 to 30080 and you can play with it by making calls to [`http://localhost:30080`](http://localhost:30080).

Take a look at [`app.Dockerfile`](docker/app.Dockerfile) and read about [multi stage Docker builds](https://docs.docker.com/build/building/multi-stage/) and [`scratch`](https://hub.docker.com/_/scratch/). 

Take a look at the [Makefile](Makefile) to see how the container is run. Try running multiple containers, crashing and restarting them. Do the results match your expectations? How are the results different from running the app locally? Think about what you would do to improve the behavior.

### K8s app

>Keep the [K8s API reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#api-overview) and [`kubectl` cheatsheet](https://kubernetes.io/docs/reference/kubectl/quick-reference/#viewing-and-finding-resources) handy for this section.
> If you feel like taking a break, check out [The Illustrated Children's Guide to Kubernetes](https://youtu.be/3I9PkvZ80BQ?si=9ywZSfYFiSbXPdqj).

You can deploy and run the application in K8s with `make app`. This will use `docker` to build the app container and `kubectl` to create various resources in a cluster. Take a look at the [Makefile](Makefile) to see what happens.

The K8s magic is included in YAML definitions of K8s resources [here](k8s/app.yaml). We have `deployment`, `configmap`, `service` and `ingress`. Inspect each resource definition carefully and think about how they combine to build your app. Can you find out how to access the application? Is there anything you could improve here?

Try accessing the app at [`http://localhost/api`](http://localhost/api). Can you crash the app? How is the behavior different from the containerized or standalone app? Can you run multiple copies of the app? What happens when you do it?

Did you also notice the [`db.yaml`](k8s/db.yaml)? Take a look at it now and pay attention to the `persistentvolumeclaim` and `persistentvolume`definitions. How does this tie into the behavior of the K8s app? Why are we missing the `configmap` and `ingress` here?

## Chaos 🐒

Read about [chaos engineering](https://en.wikipedia.org/wiki/Chaos_engineering) and follow all the links you have the patience for. It's a big subject that boils down to: _resilience has many aspects which must be tested just like any other feature of your application_. 

We've included a poor man's chaos monkey which you can let loose in your cluster with `make chaos` and turn off with `make order`. It has its own container image, built using [chaos.Dockerfile](docker/chaos.Dockerfile) and if you're curious you can take a look at [the script](k8s/chaos.sh) that does the monkeying around. Notice that the script references a couple of ENV VARS that have fallback default values. This will come in handy later.

The monkey is unleashed via the K8s YAML definition found [here](k8s/chaos.yaml). In addition to a `deployment` we have a lot of permission related resources and no service or other means of exposing the functionality to the outside. Read about them in the [K8s API reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#api-overview).

Play with your app while the chaos monkey does its thing. Can you see any differences in its behavior? Monitor the pods in your cluster and see how they behave? Which ones are being affected by our monkey? Can you make others be affected as well?

## Additional reading and research topics 📚🐛

You can use the following lists of reading and research topics as a further **introduction to Kubernetes**. The content is neither exhaustive nor particularly cutting edge. It covers basic concepts that you can later build upon towards either a DevOps or K8s developer track. 

1. Popular container runtimes 
    - [Docker](https://www.docker.com/)
    - [containerd](https://containerd.io/)
    - [podman](https://podman.io/)
    - [cri-o](https://cri-o.io/)

2. K8s derivatives
    - [Open Shift](https://www.redhat.com/en/technologies/cloud-computing/openshift)
    - [Nomad](https://www.nomadproject.io/)
    - [Rancher](https://www.rancher.com/)

3. Alternative container orchestration projects
    - [Apache Mesos](https://mesos.apache.org/)
    - [Docker Swarm](https://docs.docker.com/engine/swarm/)
    - _[Helios](https://github.com/spotify/helios) (no longer developed or supported)_

4. Turnkey K8s
    - [Minikube](https://minikube.sigs.k8s.io/docs/start/) - the original turnkey K8s
    - [k3s](https://k3s.io/) - lightweight K8s for IoT
    - [kind](https://kind.sigs.k8s.io/) - lightweight local K8s with [Docker](https://www.docker.com/) or [podman](https://podman.io/)
    - [k3d](https://k3d.io/v5.6.0/) - lightweight k3s wrapper by Rancher
    - [microk8s](https://microk8s.io/) - turnkey, high availability K8s from Canonical
    - [K8s playground](https://labs.play-with-k8s.com/) - temporary tiny clusters in the cloud

5. K8s resources
    - the **real** [Chaos Monkey](https://netflix.github.io/chaosmonkey/)
    - [Prometheus](https://prometheus.io/docs/prometheus/latest/getting_started/) - the default monitoring solution for K8s
    - [Horizntal pod autoscaling](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) documentation
    - [Operator Framework](https://sdk.operatorframework.io/)
    - [Controller Pattern](https://kubernetes.io/docs/concepts/architecture/controller/)
