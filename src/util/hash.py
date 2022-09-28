import hashlib

prefix = None

def Init(config):
    global prefix
    prefix = config['cookie_prefix']

def PrefixHash(key):
    k = prefix + key
    return hashlib.md5(k.encode()).hexdigest()

def PrefixHashCheck(key, value):
    v = hashlib.md5((prefix + key).encode()).hexdigest()
    return v == value
