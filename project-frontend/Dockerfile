FROM node:alpine

RUN mkdir /app

WORKDIR /app

COPY package* ./

RUN npm ci

COPY . /app

RUN npm run build


FROM nginx:1.17.1-alpine

COPY --from=build-step /app/build /usr/share/nginx/html