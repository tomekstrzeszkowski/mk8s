docker build -t gin_web:v1 -f Dockerfile .

docker run -it -p 2345:2345 -p 8080:8080 gin_web:v1
