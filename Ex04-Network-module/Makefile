CC=gcc

LIB=/usr/local/lib/libnanomsg.so

OBJS=pubsub.o

TARGET=pubsub

all:$(TARGET)

$(TARGET): $(OBJS)
	$(CC) $^ -o $@ $(LIB) 

%.o: %.c
	$(CC) -c $< -o $@

clean:
	rm -rf pubsub *.o *~

run:all
	./$(TARGET)

.PHONY: all clean
