FROM node:15.11.0-alpine3.10
RUN mkdir /usr/src/app
WORKDIR /usr/src/app
COPY package*.json ./
ADD package.json /usr/src/app/package.json
RUN npm ci
COPY . .
EXPOSE 80
CMD ["npm ","start"];