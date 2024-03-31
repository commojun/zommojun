build:
	docker build \
	-t commojun/zommojun \
	-f ./docker/Dockerfile \
	.

secret:
	-kubectl delete secret zommojun-secret
	kubectl create secret generic \
		--save-config zommojun-secret \
		--namespace zommojun \
		--from-env-file ./envfile
