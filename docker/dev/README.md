# Development image

## Build it yourself

| Base image | tag    |
| :-         | :-     |
| golang     | 1.16.1 |


Create a container and run the following command


```bash
$ apt update
$ apt install -y tshark default-jre
```


## Docker pull


Or pull the container image that you have uploaded to DockerHub


```bash
$ docekr pull monandkey/makeplantuml-dev:1.0
```
