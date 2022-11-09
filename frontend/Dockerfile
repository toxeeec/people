FROM node:alpine
WORKDIR /app
COPY openapi.yaml ./
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
COPY . .
RUN npm install
ENTRYPOINT ["/entrypoint.sh"]
CMD ["npm", "run", "dev"]
