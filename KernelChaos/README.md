# KernelChaos: Linux Kernel Fault Injection

## Bug

The last update of KernelChaos changed the Pod selector to the container selector, 
but one thing was not updated and seems to have been missed.

It will be called through reflection and cause error in function 'ParseNamespacedNameContainer'. 
The error message is 'too few parts of namespacedname'.

The last update incorrectly used the container name as the container ID for selecting the container.


The update is:

use ContainerSelector in kernel chaos #3395

https://github.com/chaos-mesh/chaos-mesh/pull/3395

https://github.com/chaos-mesh/chaos-mesh/commit/b80f0ca42cea242047940455e46913fb7e200f26

The last right version is 2.2.3

## Fix

* Fix inconsistent last update of KernelChaos 
  Update selector from PodSelector to ContainerSelector in file 'api/v1alpha1/kernelchaos_types.go'
* Select container by name
  Users should obviously use the name of the container instead of the ID (e.g. containerd://7dc3e8617cd......) to select the target container.
* Use widely used ContainerRecordDecoder:
  Use the ContainerRecordDecoder to parse the Pod and Container's Name and ID from record and build the ChaosDaemonClient.


## Verify

```shell
# Build in windows Command Prompt
SET CGO_ENABLED=0&&   SET GOOS=linux&&  SET GOARCH=amd64&&   go build -o demo
# Build Container
docker build -t localhost:5000/chaos-mesh-verify:KernelChaos-demo .

# Send to VPS and push to local Docker Registry
docker save localhost:5000/chaos-mesh-verify:KernelChaos-demo -o KernelChaos.tar

docker load -i KernelChaos.tar
docker push localhost:5000/chaos-mesh-verify:KernelChaos-demo

# create namespace
kubectl create ns cmv
# apply Pod in k8s
kubectl apply -f k8s.yaml
# apply KernelChaos in k8s
kubectl apply -f bio-pod.yaml
kubectl apply -f bio-container.yaml
# recover KernelChaos
kubectl delete -f bio-pod.yaml
kubectl delete -f bio-container.yaml

# get the result between apply and recover
kubectl -n cmv logs demo --all-containers=true
kubectl -n cmv logs demo -c container1
kubectl -n cmv logs demo -c container2

kubectl -n cmv describe KernelChaos bio

```

## Result (Example)

```shell
root@vmi1144888:~# kubectl -n cmv logs demo -c container1
2023/05/17 06:39:50 bio begin: 2023-05-17T06:39:50.329071919Z
2023/05/17 06:40:25 bio end: 2023-05-17T06:40:25.253606052Z
2023/05/17 06:40:25 sleep 10s
2023/05/17 06:40:35 sleep end
2023/05/17 06:40:35 bio begin: 2023-05-17T06:40:35.26172352Z
2023/05/17 06:41:09 bio end: 2023-05-17T06:41:09.858477351Z
2023/05/17 06:41:09 sleep 10s
2023/05/17 06:41:19 sleep end
# apply bio-pod, inject into pod actually only inject into the first container of the pod
2023/05/17 06:41:19 bio begin: 2023-05-17T06:41:19.864641523Z
2023/05/17 06:41:24 sync file error: sync /data.txt: input/output error
2023/05/17 06:41:24 bio end: 2023-05-17T06:41:24.668644581Z
2023/05/17 06:41:24 sleep 10s
2023/05/17 06:41:34 sleep end
2023/05/17 06:41:34 bio begin: 2023-05-17T06:41:34.669095149Z
2023/05/17 06:41:34 sync file error: sync /data.txt: input/output error
2023/05/17 06:41:34 bio end: 2023-05-17T06:41:34.669814449Z
2023/05/17 06:41:34 sleep 10s
2023/05/17 06:41:44 sleep end
2023/05/17 06:41:44 bio begin: 2023-05-17T06:41:44.678917483Z
2023/05/17 06:41:44 sync file error: sync /data.txt: input/output error
2023/05/17 06:41:44 bio end: 2023-05-17T06:41:44.679298172Z
2023/05/17 06:41:44 sleep 10s
2023/05/17 06:41:54 sleep end
2023/05/17 06:41:54 bio begin: 2023-05-17T06:41:54.685443469Z
2023/05/17 06:41:54 sync file error: sync /data.txt: input/output error
2023/05/17 06:41:54 bio end: 2023-05-17T06:41:54.685823306Z
2023/05/17 06:41:54 sleep 10s
2023/05/17 06:42:04 sleep end
2023/05/17 06:42:04 bio begin: 2023-05-17T06:42:04.69478221Z
2023/05/17 06:42:04 sync file error: sync /data.txt: input/output error
2023/05/17 06:42:04 bio end: 2023-05-17T06:42:04.696423009Z
2023/05/17 06:42:04 sleep 10s
2023/05/17 06:42:14 sleep end
2023/05/17 06:42:14 bio begin: 2023-05-17T06:42:14.705259366Z
2023/05/17 06:42:14 sync file error: sync /data.txt: input/output error
2023/05/17 06:42:14 bio end: 2023-05-17T06:42:14.707282858Z
2023/05/17 06:42:14 sleep 10s
2023/05/17 06:42:24 sleep end
2023/05/17 06:42:24 bio begin: 2023-05-17T06:42:24.713088797Z
2023/05/17 06:42:24 sync file error: sync /data.txt: input/output error
2023/05/17 06:42:24 bio end: 2023-05-17T06:42:24.713523617Z
2023/05/17 06:42:24 sleep 10s
2023/05/17 06:42:34 sleep end
2023/05/17 06:42:34 bio begin: 2023-05-17T06:42:34.721349359Z
2023/05/17 06:42:34 sync file error: sync /data.txt: input/output error
2023/05/17 06:42:34 bio end: 2023-05-17T06:42:34.721874468Z
2023/05/17 06:42:34 sleep 10s
2023/05/17 06:42:44 sleep end
2023/05/17 06:42:44 bio begin: 2023-05-17T06:42:44.729334847Z
2023/05/17 06:42:44 sync file error: sync /data.txt: input/output error
2023/05/17 06:42:44 bio end: 2023-05-17T06:42:44.729869132Z
2023/05/17 06:42:44 sleep 10s
2023/05/17 06:42:54 sleep end
# recover the first container from bio-pod
2023/05/17 06:42:54 bio begin: 2023-05-17T06:42:54.738351912Z
2023/05/17 06:43:34 bio end: 2023-05-17T06:43:34.68909588Z
2023/05/17 06:43:34 sleep 10s
2023/05/17 06:43:44 sleep end
# apply bio-pod, inject into the container2 container, the container1 container is not injected
2023/05/17 06:43:44 bio begin: 2023-05-17T06:43:44.697598362Z
2023/05/17 06:44:15 bio end: 2023-05-17T06:44:15.705196962Z
2023/05/17 06:44:15 sleep 10s
2023/05/17 06:44:25 sleep end
2023/05/17 06:44:25 bio begin: 2023-05-17T06:44:25.713558149Z
2023/05/17 06:44:46 bio end: 2023-05-17T06:44:46.558084819Z
2023/05/17 06:44:46 sleep 10s
2023/05/17 06:44:56 sleep end
2023/05/17 06:44:56 bio begin: 2023-05-17T06:44:56.561593008Z
2023/05/17 06:45:36 bio end: 2023-05-17T06:45:36.231970237Z
2023/05/17 06:45:36 sleep 10s
2023/05/17 06:45:46 sleep end
2023/05/17 06:45:46 bio begin: 2023-05-17T06:45:46.240986258Z
```

```shell
root@vmi1144888:~# kubectl -n cmv logs demo -c container2
2023/05/17 06:39:50 bio begin: 2023-05-17T06:39:50.442120597Z
2023/05/17 06:40:25 bio end: 2023-05-17T06:40:25.853170024Z
2023/05/17 06:40:25 sleep 10s
2023/05/17 06:40:35 sleep end
2023/05/17 06:40:35 bio begin: 2023-05-17T06:40:35.860858451Z
2023/05/17 06:41:12 bio end: 2023-05-17T06:41:12.935803092Z
2023/05/17 06:41:12 sleep 10s
2023/05/17 06:41:22 sleep end
# apply bio-pod, inject into the first container, the second container is not injected
2023/05/17 06:41:22 bio begin: 2023-05-17T06:41:22.944864828Z
2023/05/17 06:41:54 bio end: 2023-05-17T06:41:54.149188578Z
2023/05/17 06:41:54 sleep 10s
2023/05/17 06:42:04 sleep end
2023/05/17 06:42:04 bio begin: 2023-05-17T06:42:04.160521782Z
2023/05/17 06:42:24 bio end: 2023-05-17T06:42:24.692334655Z
2023/05/17 06:42:24 sleep 10s
2023/05/17 06:42:34 sleep end
2023/05/17 06:42:34 bio begin: 2023-05-17T06:42:34.700905656Z
2023/05/17 06:42:55 bio end: 2023-05-17T06:42:55.523709405Z
2023/05/17 06:42:55 sleep 10s
2023/05/17 06:43:05 sleep end
2023/05/17 06:43:05 bio begin: 2023-05-17T06:43:05.532872723Z
2023/05/17 06:43:45 bio end: 2023-05-17T06:43:45.564530384Z
2023/05/17 06:43:45 sleep 10s
2023/05/17 06:43:55 sleep end
# apply bio-pod, inject into the second container
2023/05/17 06:43:55 bio begin: 2023-05-17T06:43:55.573690734Z
2023/05/17 06:43:55 sync file error: sync /data.txt: input/output error
2023/05/17 06:43:55 bio end: 2023-05-17T06:43:55.574313805Z
2023/05/17 06:43:55 sleep 10s
2023/05/17 06:44:05 sleep end
2023/05/17 06:44:05 bio begin: 2023-05-17T06:44:05.581097126Z
2023/05/17 06:44:05 sync file error: sync /data.txt: input/output error
2023/05/17 06:44:05 bio end: 2023-05-17T06:44:05.586810294Z
2023/05/17 06:44:05 sleep 10s
2023/05/17 06:44:15 sleep end
2023/05/17 06:44:15 bio begin: 2023-05-17T06:44:15.595497314Z
2023/05/17 06:44:15 sync file error: sync /data.txt: input/output error
2023/05/17 06:44:15 bio end: 2023-05-17T06:44:15.596290131Z
2023/05/17 06:44:15 sleep 10s
2023/05/17 06:44:25 sleep end
2023/05/17 06:44:25 bio begin: 2023-05-17T06:44:25.602937183Z
2023/05/17 06:44:25 sync file error: sync /data.txt: input/output error
2023/05/17 06:44:25 bio end: 2023-05-17T06:44:25.603910017Z
2023/05/17 06:44:25 sleep 10s
2023/05/17 06:44:35 sleep end
2023/05/17 06:44:35 bio begin: 2023-05-17T06:44:35.628120166Z
2023/05/17 06:44:35 sync file error: sync /data.txt: input/output error
2023/05/17 06:44:35 bio end: 2023-05-17T06:44:35.629295948Z
2023/05/17 06:44:35 sleep 10s
2023/05/17 06:44:45 sleep end
# recover the second container (container2) from bio-container
2023/05/17 06:44:45 bio begin: 2023-05-17T06:44:45.636969502Z
2023/05/17 06:45:25 bio end: 2023-05-17T06:45:25.313074695Z
2023/05/17 06:45:25 sleep 10s
2023/05/17 06:45:35 sleep end
2023/05/17 06:45:35 bio begin: 2023-05-17T06:45:35.321199428Z
2023/05/17 06:46:15 bio end: 2023-05-17T06:46:15.579687843Z
2023/05/17 06:46:15 sleep 10s
```
