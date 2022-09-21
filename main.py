import os
from flask import Flask
from flask import send_from_directory
from flask import request
from flask import make_response
from waitress import serve
import logging
import json

logging.basicConfig(
        level = logging.DEBUG,
        format='[%(asctime)s][%(filename)s][%(levelname)s][%(message)s]',
        datefmt='%Y-%m-%d %H:%M:%S',
        )

app = Flask(__name__)

@app.route('/resource/<path:path>')
def send_report(path):
    return send_from_directory('resource', path)

@app.route("/", methods=['GET'])
@app.route("/index.html", methods=['GET'])
def main_index():
    return send_report("index.html")

@app.route("/login", methods=['POST'])
def login():
    info = request.get_data(as_text=True)
    logging.debug(info)
    resp = make_response("ok")
    resp.set_cookie('lcookie', "xxxxx")
    return resp, 200

if __name__ == "__main__":
    logging.info('starting...')
    serve(app, host="0.0.0.0", port=8989)