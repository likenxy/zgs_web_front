import os
from flask import Flask
from flask import send_from_directory
from flask import request
from flask import make_response
from waitress import serve
import hashlib
import json
import src.db.zgs_mysql as zgs_mysql
import src.util.hash as zgs_hash
import logging
import datetime

logging.getLogger('sqlalchemy').setLevel(logging.ERROR)
logging.basicConfig(
        level = logging.DEBUG,
        format='[%(asctime)s][%(filename)s][%(levelname)s][%(message)s]',
        datefmt='%Y-%m-%d %H:%M:%S',
        )

app = Flask(__name__)

@app.route("/", methods=['GET'])
@app.route("/index.html", methods=['GET'])
def main_index():
    return send_report("index.html")

@app.route("/login", methods=['POST'])
def login():
    info = request.get_data(as_text=True)
    info = json.loads(info)
    ok = zgs_mysql.CheckUserPassword(info['name'], info['password'])
    logging.debug("login status : [{}]".format(ok))
    if ok:
        # set cookie
        user_cookie = zgs_hash.PrefixHash(info['name'])
        resp = make_response("ok")
        e = datetime.datetime.now() + datetime.timedelta(days=1)
        resp.set_cookie('lcookie', user_cookie, expires=e)
        resp.set_cookie('username', info['name'], expires=e)
        return resp, 200
    return "", 400

def _login_check(request):
    c = request.cookies.get('lcookie')
    u = request.cookies.get('username')
    if c is None or u is None:
        return False
    if zgs_hash.PrefixHashCheck(u, c):
        return True
    return False


@app.route("/login_check", methods=['GET'])
def login_check():
    if _login_check(request):
        return 'True', 200
    return 'False', 400

@app.route('/<path:path>')
def send_report(path):
    return send_from_directory('resource', path)


def Init():
    dbConfig = json.loads(open('config/db.json').read())
    zgs_mysql.InitMysql(dbConfig)

    config = json.loads(open('config/config.json').read())
    zgs_hash.Init(config)

if __name__ == "__main__":
    Init()
    logging.info('starting...')
    serve(app, host="0.0.0.0", port=8989)
