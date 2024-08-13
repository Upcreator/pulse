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

# Install serve globally
RUN npm install -g serve

# Expose port 5000 for serve
EXPOSE 5000

# Command to run the app using serve
CMD ["serve", "-s", "dist", "-l", "5000"]
