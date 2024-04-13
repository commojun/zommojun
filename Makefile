build:
	docker build \
	-t commojun/zommojun \
	-f ./docker/Dockerfile \
	.

push:
	docker push commojun/zommojun:latest

secret:
	-kubectl delete secret zommojun-secret
	kubectl create secret generic \
		--save-config zommojun-secret \
		--namespace zommojun \
		--from-env-file ./envfile
