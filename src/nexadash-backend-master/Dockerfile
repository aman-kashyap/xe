FROM python:2.7

COPY . /nexadash-backend

RUN pip install -r nexadash-backend/requirements.txt

EXPOSE 5000

CMD python nexadash-backend/wsgi.py
