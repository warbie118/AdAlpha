#!/bin/bash

#RUNS BACKEND SERVICES
docker-compose build -d
docker-compose up -d

#RUNS UI
cd ui
npm install
npm run serve



