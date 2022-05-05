# Kubernetes Architecture

![Kubernetes-Architecture](https://github.com/prateek041/DevOps-90days/blob/main/resources/Kubernetes/Day-1-Kubernetes-concepts(Civo)/Kubernetes-Architecture.png)

## What happens here ?

1. ***Master node*** : responsible for managing the worker nodes.

2. ***Kubectl*** :the command line tool that is used to communicate with the cluster.
  
    ` $ kubectl run "podname" "image" `
  
3. ***API server*** : The request goes here.
- When the request reaches the API server, Authentication, Authorization and Administration takes place.
  - With the help of the headers that are passed, the request is Authenticated.
  - Using RBAC rules, it is Authorised, wether the user is allowed to make such requests.

4. ***Schedular***: has all the information about the nodes and the resources.
- Using all the information it has about the nodes, it will find the best pod to do the specific job.
  - On finding the suited node, it gives its information to the API server.

5. ***Worker node*** : responsible for managing the containers and pods.


6. ***Kubelet*** : main component on the worker side.
- Keeps on asking the API server if there is a work load to run. (lets say node 1 is ready)
  - Schedular tells API server to run workload on node 1.
  - API server tells the Kubelet of node 1 to run the workload.
  - On getting the instruction from the API server, Container runtime gets into play.

7. ***Container runtime*** : container runtime like docker, containerd etc.
- performs the task of:
  - pulling the images to run.
  - setting up the right configuration, needed.

8. ***Kubeproxy*** : The core networking component of Kubernetes .
- What it does ?
  - Manages networking and communication across the cluster.
  - Manages all the pod to pod communication and the rules.

9. ***Controller manager*** : Set of controllers that are running. example: replication controller manager etc.
- What it does ?
  - Let's say you want to run 4 instances/nodes always, the controller manager will make sure it happens, in case some of them collapse, it will restart a new one to satisfy the instructions.
  - There are many more such controller managers.

10. ***ETCD*** : High value, key value database for kubernetes and stores all the info about the Kubernetes objects.
- What it does ?
  - Stores the cluster state.
  - Any pod spec or configuration is changing, API server stores it in ETCD.
