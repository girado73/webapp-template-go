# Stage 1: Build the React application
FROM node:18-alpine AS builder

WORKDIR /web

# Copy package.json and package-lock.json
COPY package.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Serve the application with Nginx
FROM nginx:alpine

# Copy the built application from the builder stage
COPY --from=builder /web/dist /usr/share/nginx/html

# Copy the Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf

RUN nginx -t

# After copying files, ensure nginx has proper permissions
RUN chown -R nginx:nginx /usr/share/nginx/html && \
    chmod -R 755 /usr/share/nginx/html


# Expose port 80
EXPOSE 80


# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
