k8s_version = v1.24


setup:
	asdf install
	@echo 'All runtimes installed'
	@echo ''
	ctlptl create cluster kind --registry=ctlptl-registry --kubernetes-version=${k8s_version}

teardown: clean-setup clean-docker

clean-setup:
	ctlptl delete cluster kind
	@echo ''
	docker rm -f ctlptl-registry
	@echo 'ctlptl-registry deleted'
	@echo ''

clean-docker:
	docker volume prune -f
	docker network prune -f
	docker builder prune -f
