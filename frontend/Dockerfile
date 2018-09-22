FROM node:alpine

COPY ./assets /app/assets
COPY ./components /app/components
COPY ./layouts /app/layouts
COPY ./middleware /app/middeware
COPY ./pages /app/pages
COPY ./plugins /app/plugins
COPY ./static /app/static
COPY ./store /app/store
COPY ./nuxt.config.js /app/
COPY ./package.json /app/

WORKDIR /app/

RUN npm i && npm run build

EXPOSE 3000

ENTRYPOINT ["npm", "run", "start"]