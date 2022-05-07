# Kubernetes Namespaces:

They are used to provide isolated environments to different team, groups within the same cluster. Consider it as multiple virtual cluster within the same physical cluster.

` kubectl get ns`

The above command is used to get all the current namespaces.

` kubectl create nd dev`

The above command is used to create new namespace called 'dev'.

There can be different resource quotas applied to the namespace to ensure, this namespace does not use too much cluster resources.
### Use cases:
1. Providing resource quota to differnt teams working on the same cluster.
    - you can have different teams like dev, testing etc.
    - you can provide them their seperate namespace with resource quota available to them.
    - In this way, none of them overuse the cluster resources to do their work.

2. Running differnt versions of same object in the same cluster.
    - `kubectl run nginx --image=nginx`
      - this will create a pod nginx in the default namespace, because we havent provided any ns.
    - `kubectl run nginx --image=nginx`
      - this will throw an error, that notifies that the pod is already running.
    - `kubectl run nginx --image=nginx -n dev`
      - this will create the pod nginx in the 'dev' namespace.
    - `kubectl get pods -n dev`
      - this command will return the pods running in the 'dev' namespace.

### Context:
Every command without any namespace specification is triggered in default namespace, it is because of 'context'. Specifying the namespace again and again for commands for other name can be hectic, so there is a way of changing the context.

`kubectl config set-context --current --namespace=dev`

this will change the context default namespace to dev. to see it:

`kubectl config view --minify | grep namespace`

