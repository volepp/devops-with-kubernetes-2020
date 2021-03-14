FROM node:15.11.0-alpine3.10
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
EXPOSE 80
CMD ["npm ","start"];