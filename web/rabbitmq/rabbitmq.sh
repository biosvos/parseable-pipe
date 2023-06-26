#!/bin/bash

podman build -t localhost:5000/rabbitmq:latest .
podman push --tls-verify=false localhost:5000/rabbitmq:latest