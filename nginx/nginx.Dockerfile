FROM nginx

COPY grpc_gateway.conf /etc/nginx/conf.d
COPY errors.grpc_conf /etc/nginx/conf.d