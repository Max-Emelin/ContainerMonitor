FROM node:18-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

EXPOSE 80

ARG SERVER_PORT
ENV REACT_APP_SERVER_PORT=$SERVER_PORT

CMD ["npm", "start"]
