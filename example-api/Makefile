# from: https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/
APP?=ExampleAppName
PORT?=8080

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

run: build
	PORT=${PORT} ./${APP}
