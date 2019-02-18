build:
	find -name "*.java" | xargs javac -d bin

run:
	java -cp Main

clean:
	rm -rf bin/*