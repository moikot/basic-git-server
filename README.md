# Basic Git server
![ci](https://github.com/moikot/basic-git-server/workflows/ci/badge.svg)
An very basic [GITTP library](https://github.com/adamveld12/gittp) based Git server packaged in a multi-platform Docker image.

Supported platfroms include:
* linux/amd64 - For running on Intell or AMD based systems.
* linux/arm64 - For running on Raspberry Pi and other ARM v8 64 bit based SBCs.

## Install

Basic Git sever Docker container doesn't need anything except a mapped volume. The server automatically creates a subdirectory when you do a `git push`.

### Run it as a standalone Docker container

```bash
$ mkdir srv-repos 
$ docker run -d -p 8080:8080 -v $(pwd)/srv-repos:/repos moikot/basic-git-server -d /repos
```

Try to create a test repo and push your changes to the server exposed on http://localhost:8080:
```bash
$ mkdir test-repo && cd test-repo && touch test.txt
$ git init && git add . && git commit -m "The initial commit"
$ git remote add origin http://localhost:8080/test-repo
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Writing objects: 100% (3/3), 212 bytes | 212.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To http://localhost:8080/test-repo
 * [new branch]      master -> master
Branch 'master' set up to track remote branch 'master' from 'origin'.
```

### Run it as a Kubernetes service 

Basic Git server comes with a Helm chart. To use the chart Helm must be configured for your Kuberntes cluser. Please refer to the Kuberntes and Helm documentation for the details.

The versions required are:
* **Helm 3.0+**
* **Kubernetes 1.8+**

*Note: It is possible it works with earlier versions but this chart is untested for any other versions.*

To install the latest version of this chart, add the `moikot` Helm repository and run `helm install`:

```bash
$ helm repo add moikot https://moikot.github.io/helm-charts
"moikot" has been added to your repositories

$ helm install basic-git-server moikot/basic-git-server
```

Please see the options supported in the [`values.yaml`](https://github.com/moikot/helm-charts/blob/master/charts/basic-git-server/values.yaml) file.