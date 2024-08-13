FROM node:18 AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN npm install -g serve

FROM node:18

WORKDIR /app

COPY --from=build /app/dist /app/dist

RUN npm install -g serve

EXPOSE 5000

CMD ["serve", "-s", "dist", "-l", "5000"]
