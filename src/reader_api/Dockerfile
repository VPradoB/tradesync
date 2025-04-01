FROM python:3.13-slim

RUN apt update && apt install -y python3-distutils
COPY . .

RUN pip install -r requirements.txt

CMD ["fastapi", "run"]
