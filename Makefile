
all:
	go buld

install:
	mv CheckJSONSchema ~/bin

test:
	go build
	./CheckJSONSchema -s t1.json -d d1.json

	
