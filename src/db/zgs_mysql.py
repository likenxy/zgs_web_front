from sqlalchemy import create_engine
from sqlalchemy import text
from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base
import json

mysqlConn = None
session = None

def InitMysql(config):
    global mysqlConn
    global session
    mysqlConn = create_engine(
            config['db'],
            encoding='utf-8',
            # echo=True,
            pool_recycle=-1
            )
    session = sessionmaker(mysqlConn)

# user info table orm
Base = declarative_base()
class UserInfo(Base):
    __tablename__ = 'user_info'
    id = Column(Integer, primary_key=True, autoincrement=True)
    user_name = Column(String(100),nullable=False)
    password = Column(String(100),nullable=False)
    __table_args__ = {
            "mysql_charset":"utf8"
            }

def CheckUserPassword(user_name, password):
    s = session()
    result = s.query(UserInfo).filter(text("user_name=:user_name")).params(user_name=user_name).first()
    if result is not None:
        return result.password == password
    return False
