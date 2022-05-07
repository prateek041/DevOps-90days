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

What happens when you run `kubectl create -f "yaml file"`.

![](https://github.com/prateek041/DevOps-90days/blob/main/resources/Kubernetes/Day-2-Kubernetes(Civo)/assets/pods-lifecycle.png)

1. YAML file is converted to JSON file and sent to the API server.
2. Authentication using kubeconfig credentials.
3. when authenticated, authorization is done to check wether the user is authorized to execute that command.
4. when authorized, the data is sent to ETCD and persisted.
5. [STATE] **Pending**:
    - *Schedular* finds the best suited node, depending on various factors like resources etc.
    - When found, it will fill the label spec-mode-name and send it back to the API server.
    - This request is then stored in ETCD again.
    - *API server* then tells the decided kubelet to spin the container on the particular node.
    - *Kubelet* then fetches the image and other requirements using CRI. CNI gets the IP address to the pod.
    - that IP address is sent back to the API server and it's data is persisted in ETCD.
6. [STATE] **Container creating**:
    - Image is pulled and the container is created with the help *containerd* and *runc*.
7. [STATE] **Running state**:
    - The container starts running.
8. [STATE] **crashloopbackoff** and **succeded** state:
    - When the process dies too many times in the pod, it goes to crashloopbackoff.
    - when the process is successfull, it goes to Succeded state.

### Hooks:
1. Post-start-hook:
    - When you want some actions to be performed before the required container is started.
2. Pre-stop-hook:
    - When you want some actions to be performed before the container gets terminated.

### Init container:
It will run before starting the main container.
