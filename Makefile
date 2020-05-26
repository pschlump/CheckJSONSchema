
all:
	go build

install: all
	mv CheckJSONSchema ~/bin

test: all
	./CheckJSONSchema -s t1.json -d d1.json

	
