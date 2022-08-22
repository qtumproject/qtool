FROM node AS react-build
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine

# Nginx config
RUN rm -rf /etc/nginx/conf.d
COPY ./nginx/conf /etc/nginx

# Static build
COPY --from=react-build /app/build /usr/share/nginx/html

# Default port exposure
EXPOSE 80

# Copy .env file and shell script to container
WORKDIR /usr/share/nginx/html
COPY ./env.sh .
COPY .env .

# Add bash
RUN apk add --no-cache bash

# Make our shell script executable
RUN chmod +x env.sh

# Script 'env.sh' will inject environment vars defined within .env into index.html to be used by React
# CMD ["/bin/bash", "-c", "/usr/share/nginx/html/env.sh && nginx -g \"daemon off;\""]
# Start Nginx server
CMD ["nginx -g \"daemon off;\""]