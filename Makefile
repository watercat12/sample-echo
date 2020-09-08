build:
	docker build -t watercat12/arm-go .
push:
	docker push watercat12/arm-go:latest
kube:
	kubectl create deployment arm-go-2 --image watercat12/arm-go:latest
getPodName:
	kubectl get pods | awk '/arm/ {print $$1}'
log:
	$(shell kubectl get pods | awk '/arm/ {print $$1}' | while read -r line ; do \
	 kubectl logs $$line ; \
	 done;)
exp:
	kubectl expose deployment arm-go --type="NodePort" --port 1323
deleteKube:
	kubectl delete deployments.apps arm-go
getPods:
	kubectl get pods
getDep:
	kubectl get deployments.apps
clean:
	docker rmi -f watercat12/arm-go