import hashlib
import sys
print(sys.argv[1], hashlib.md5(sys.argv[1].encode()).hexdigest())
