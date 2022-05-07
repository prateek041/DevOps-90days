# Pods:

### What are pods ?
- smallest unit in kubernetes.
- collection of one or more containers.
- specification, telling how the containers will be running inside the cluster.

![](https://github.com/prateek041/DevOps-90days/blob/main/resources/Kubernetes/Day-2-Kubernetes(Civo)/assets/pods.png)

## Hands on using minikube and yaml files:

1. Install minikube using the video [Installing minikube on your local Kubernetes cluster - civo academy](https://www.youtube.com/watch?v=HxPqUf1_bIk&list=PLhc-GEHI0F7_wsxeSNqniQqSforzp92mV&index=12).
2. Run command: `minikube start`, to start a local kubernetes cluster.
3. On your desktop, create a file sample-pod.yaml and put the following [code](), into the file you created.
4. run : `kubectl create -f "file path"`, where file path is the path of "sample-code.yaml" file on your desktop.

      ![](https://github.com/prateek041/DevOps-90days/blob/main/resources/Kubernetes/Day-2-Kubernetes(Civo)/assets/create-pod.png)

5. your pod is created, to get the info, run `kubectl get pods`.
6. creating a pod in imperative way, run: `kubectl run demo --image=nginx --port=80`

    ![](https://github.com/prateek041/DevOps-90days/blob/main/resources/Kubernetes/Day-2-Kubernetes(Civo)/assets/pods-imperative-way.png).
    
7. check the logs using `kubectl logs "pod name"`
8. check events and description of pod using `kubectl describe pod "podname"` 
9. to delete the pod, use: `kubectl delete pod "pod name"`
10. to get inside a pod, use: `kubectl exec -it "pod name" bash`

# Pods lifecycle:

