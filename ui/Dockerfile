FROM node:8-alpine
WORKDIR ./ui_js
COPY ./package.json .
RUN npm install --quiet
COPY . .
EXPOSE 8080
