FROM node:alpine as build-step

RUN mkdir /app

WORKDIR /app

COPY package* ./

RUN npm ci

COPY . /app

ARG REACT_APP_NAMESPACE
ENV REACT_APP_NAMESPACE $REACT_APP_NAMESPACE

RUN npm run build

FROM nginx:1.19.8-alpine

COPY --from=build-step /app/build /usr/share/nginx/html
COPY --from=build-step /app/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]