Commands:
```
docker build -t my-go-app .
$ docker run -p 7500:7500 -it my-go-app
```

container is configured to warsaw time zone

working endpoints:
/lublin
/sydney
/new_york

To build container from your local computer:
docker build https://github.com/IRusio/cloudsOne.git

To run container from your local computer without downloading:
docker run -p 7500:7500 -e github="https://github.com/IRusio/cloudsOne" -it my-go-app
