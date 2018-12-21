gen:
	cd services; docker run -v `pwd`:/defs gofunct/prototool:1.17_0 all

build-docker:
	cd hack; make build

clean-all-images:
	docker rmi -f $(docker images -a -q)

clean-all-containers:
	docker rm -vf $(docker ps -a -q)