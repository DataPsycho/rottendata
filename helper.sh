# Rebuild New Image
docker build -t demoapp .

# Runt the New Image with Linked Port
docker run -p 8080:8080 demoapp:latest

# Module Path: github.com/DataPsycho/portfolio/src