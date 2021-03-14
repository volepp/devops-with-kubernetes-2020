from node:alpine

RUN mkdir /app

WORKDIR /app

COPY package* /app

RUN npm ci

COPY . .

CMD ["npm", "start"]
