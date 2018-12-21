gen:
	cd services; docker run -v `pwd`:/defs colemanword/prototool:1.17_0 all

build-docker:
	cd hack; make build