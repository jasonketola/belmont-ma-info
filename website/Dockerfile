# Use a nginx Docker image
FROM nginx
# Copy the static HTMLs to the nginx directory
COPY hugo/public /usr/share/nginx/html
# Copy the nginx configuration template to the nginx config directory
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
