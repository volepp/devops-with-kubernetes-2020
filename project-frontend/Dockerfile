from node:alpine

WORKDIR /usr/src/app

ENV PATH /user/src/app/node_modules/.bin:$PATH

COPY package* ./

RUN npm ci

COPY . .

CMD ["npm", "start"]
