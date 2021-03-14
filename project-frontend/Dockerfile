FROM node:alpine

WORKDIR /usr/src/app

COPY package* ./

RUN npm ci

COPY . .

RUN npm run build


FROM nginx:1.17.1-alpine

COPY --from=build-step /usr/src//app/build /usr/share/nginx/html