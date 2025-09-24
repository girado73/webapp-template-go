### STAGE Basis ###
FROM node:lts AS base

WORKDIR /web

### STAGE Development ###
FROM base AS dev

ENV PATH /web/node_modules/.bin:$PATH

EXPOSE 5173 
CMD ["sleep", "100000"]

### STAGE Build ###
FROM base AS build

COPY package.json ./

RUN npm install

COPY . .

RUN npm run build --omit=dev

### STAGE prod ###
FROM nginx:alpine AS prod

RUN apk add --no-cache curl
RUN apk add --no-cache bash
COPY --from=build /web/build /usr/share/nginx/html/
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
