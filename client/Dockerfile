# Development stage
FROM node:18-alpine AS development
WORKDIR /app
# Copy package files
COPY package*.json ./
# Update Vite to latest version first
RUN npm install vite@latest
# Install other dependencies
RUN npm install
RUN npm install tailwindcss @tailwindcss/vite @tailwindcss/postcss postcss autoprefixer@latest
RUN npm install axios @types/axios
# Copy source code
COPY . .
# Expose port
# EXPOSE 3000
EXPOSE 80
# Start development server
# CMD ["npm", "run", "dev", "--host"]
CMD ["npm", "run", "dev", "--", "--host", "--port", "80"]


# Production build stage
FROM node:18-alpine AS builder
WORKDIR /app
# Copy package files
COPY package*.json ./
# Update Vite to latest version first
RUN npm install vite@latest
# Install dependencies
RUN npm install
RUN npm install tailwindcss @tailwindcss/vite @tailwindcss/postcss postcss
RUN npm install axios @types/axios
# Copy source code
COPY . .
# Build the application
RUN npm run build


# Production stage
FROM nginx:alpine AS production
# Copy built assets from builder
COPY --from=builder /app/dist /usr/share/nginx/html
# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf
# Expose port
EXPOSE 80
# Start nginx
CMD ["nginx", "-g", "daemon off;"]

