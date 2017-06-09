NAME = testtech

SRC_PATH = src

FILE = main.go\
		read_file.go\
		database.go\

SRC = $(addprefix $(SRC_PATH)/,$(FILE))

CC = go build

CR = go run

export GOPATH := $(shell pwd)

all : dependance $(NAME)

$(NAME):
	$(CC) -o $(NAME) $(SRC)

clean:
	rm $(NAME)

dependance:
	apt-get install golang
	apt-get install mysql-server
	sudo service mysql start
	go get github.com/go-sql-driver/mysql
