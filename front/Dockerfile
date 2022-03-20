FROM node:16-alpine3.11

WORKDIR /app

COPY . /app

RUN sh -c "npm i && npm run build"

#EXPOSE $PORT

CMD ["npm", "run", "serve"]