# Labels and Selectors:

**Labels** in kubernetes are set of key-value pairs attached to kubernetes objects like pods. They are used to provide attrubutes to create meaningful objects.

`kubectl get 'object name' --show-labels`

### considering pods as objects:

- `kubectl get pods --show-labels`
- the above command will show the label attached to the pod.
- The labels can be attached during the pod creation or even after the creation.

- `kubectl label pod nginx demo=true`

- The above command will add the label to the object pod.

Couldnt understand what was ahead, will look for other resources, or maybe things will be explained better later on !

I think it is something like CSS selectors, that select the objects with the label they are assigned to.
