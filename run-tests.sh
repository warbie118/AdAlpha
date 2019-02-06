#!/bin/bash

docker-compose -f docker-compose-test.yml up -d

echo "Waiting for docker containers to initialise"
sleep 5

cd api
if go test -v | grep -q 'FAIL'; then
    echo "API - FAILED (run 'go test' in package for details)"
else
    echo "API - PASS"
fi
cd ..
cd model
if go test -v | grep -q 'FAIL'; then
    echo "MODEL - FAILED (run 'go test' in package for details)"
else
    echo "MODEL - PASS"
fi
cd ..

docker-compose -f docker-compose-test.yml down

