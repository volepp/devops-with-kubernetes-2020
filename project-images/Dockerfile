from node:alpine

RUN mkdir /app

WORKDIR /app

COPY package* ./

RUN npm ci

COPY . .

CMD ["npm", "start"]
