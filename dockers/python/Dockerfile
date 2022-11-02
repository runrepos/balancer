#FROM python:3.8
#FROM python:3.10-slim
FROM python:3.10
#FROM python:3.11
ENV PYTHONUNBUFFERED 1

WORKDIR /app

RUN apt-get update
RUN apt-get -y install gcc libffi-dev libssl-dev
#libxml2-dev libxslt-dev swig

ADD /req.txt /app/req.txt
RUN python3 -m pip install -r /app/req.txt --no-cache-dir 
#--no-binary=:all: --no-deps --ignore-installed

# RUN adduser --disabled-password --gecos '' user
# USER user

EXPOSE 3000

#ENTRYPOINT ["python3", "/app/api_sanic/api.py"]
#CMD ["python3", "api_sanic/api.py"]

#ENTRYPOINT ["python"]
#CMD ["main.py"]
