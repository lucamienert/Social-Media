FROM node:18

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

ENV NEXT_PRIVATE_SWC_USE_WASM=true

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
