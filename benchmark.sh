#!/bin/bash

url="https://localhost:8443/index"

ab -f TLS1.2 -kc 300 -n 10000 -m POST $url
