#!/bin/sh

go build

# Run 100 parallel processes
for i in $(seq 0 100) ; do
	./example &
done

wait

# Check the file, it should print "Helloworld\n"
cat file.txt
