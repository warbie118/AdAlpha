FROM node:7.7.2-alpine
WORKDIR ./ui
COPY package.json .
RUN npm install --quiet
COPY . .
EXPOSE 8080
