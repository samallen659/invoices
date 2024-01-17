FROM node:20-alpine3.19 AS BUILD_IMAGE
WORKDIR /app/react-app

COPY package.json .

RUN npm install

COPY . .

RUN npm run build

FROM node:20-alpine3.19 as PRODUCTION_IMAGE
WORKDIR /app/react-app

COPY --from=BUILD_IMAGE /app/react-app/dist /app/react-app/dist/
EXPOSE 80

COPY package.json .
COPY vite.config.ts .

RUN npm install typescript

EXPOSE 80
CMD ["npm","run","preview"]
