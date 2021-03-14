FROM node:15.11.0-alpine3.10
WORKDIR /usr/src/app
COPY package*.json ./
ADD package.json /usr/src/app/package.json
RUN npm install
RUN npm install react-scripts@4.0.3 -g
COPY . .
EXPOSE 80
CMD ["npm ","start"];