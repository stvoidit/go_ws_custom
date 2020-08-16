build:
	mkdir build
	cd src/backend && go install && go build -o ../../build .
	cd src/frontend && yarn && yarn build