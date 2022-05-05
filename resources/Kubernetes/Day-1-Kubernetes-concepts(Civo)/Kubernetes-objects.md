# Kubernetes Objects:

Anything you want to create in a Kubernetes cluster is a kubernetes object.

#### Ways to create a kubernetes object:
- Imperative way.
- Declarative way, using resources like YAML, JSON etc files.

` $ Kubectl create deploy nginx --image=nginx --port=80 --dry-run=client -oyaml > deploy.yaml `

