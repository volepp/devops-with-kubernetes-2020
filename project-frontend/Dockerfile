FROM node:15
WORKDIR /app
COPY package*.json /app
RUN npm ci
COPY . /app
EXPOSE 80
CMD ["npm ","start"];