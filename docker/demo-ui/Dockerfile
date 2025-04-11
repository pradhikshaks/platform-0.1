# FROM node:23-alpine

# WORKDIR /app
# COPY . .
# RUN yarn
# EXPOSE 8002
# ENTRYPOINT ["yarn","run","dev"]

FROM node:23-alpine

WORKDIR /app
COPY . .

RUN yarn install

EXPOSE 8002

CMD ["yarn", "dev"]
