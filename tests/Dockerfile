FROM python:3.9-slim

WORKDIR app

COPY req.txt .

RUN pip install -r req.txt

COPY test.py .

CMD ["pytest", "test.py"]
