# Kubernetes Objects:

Anything you want to create in a Kubernetes cluster is a kubernetes object.

#### Ways to create a kubernetes object:
- Imperative way.
- Declarative way, using resources like YAML, JSON etc files.

` $ Kubectl create deploy nginx --image=nginx --port=80 --dry-run=client -oyaml > deploy.yaml `

The above command will create an object called deployment in kubernetes and it will manage it.

The yaml file will contain:
- apiVersion : to which the objects belong.
- kind : kind of object you want to deploy and kubernetes to manage that.
- metadat : define name of the tags and application name.
- spec : to define the desired state of the object.

#### To create an object in imperative way:

` kubectl create deploy nginx --image=nginx --port=80`

The above command will create the object.

#### To create an object using declarative way:

` kubectl create -f deploy.yaml`

The above command will create the object using the data in deploy.yaml

### The whole procedure:
- takes the YAML file and converts it into JSON.
- Give it to the API server.
- further it is deployed and persisted in ETCD.
