from node:alpine

RUN mkdir /app

WORKDIR /app

COPY package* ./

RUN npm ci

COPY . /app

CMD ["npm", "start"]
