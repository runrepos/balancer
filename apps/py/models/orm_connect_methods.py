#import pandas as pd
import os, asyncmy #, pymysql
#from sqlalchemy import Column, Integer, String #, ForeignKey, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker # relationship , Session
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine

# процент соотношение float
# ссылка host ссылка
# файл сервера m3u8

DATABASE_URL = (f"mysql+asyncmy://root:{os.environ.get('MYSQL_ROOT_PASSWORD')}@"
                f"{os.environ.get('MYSQL_HOST')}:{os.environ.get('MYSQL_PORT')}"
                f"/{os.environ.get('MYSQL_DATABASE')}") # "mysql+asyncmy://root:mysqlpass@localhost:3306/test"

Base = declarative_base() ; print(DATABASE_URL)
engine = create_async_engine(DATABASE_URL, echo=True)
async_session = sessionmaker( engine, class_=AsyncSession, expire_on_commit=False )


class DB:

    def __init__(self):
        self.Base = Base
        self.async_session = async_session
        pass

    # функции форматирования данных
    @staticmethod
    def to_dict(a):
        a = a.__dict__
        if '_sa_instance_state' in a: del(a['_sa_instance_state'])
        return a
    
    def list_result_to_dict(self, result, filter_row=None):
        if filter_row is None: filter_row=[]
        return [ self.convert_row_by_filter(
                        filter = filter_row,
                        row = DB.to_dict([row for row in colmn ][0])
                    ) for colmn in result.all()]
    
    @staticmethod
    def convert_row_by_filter(row, filter):
        if filter == None: return row
        for key in row:
            if key in filter: 
                try: row[key]=filter[key](row[key])
                except: row[key]='error filter'
        return row
    
    # функции интерфейсы
    async def execute_to_dict(self, session, sql, values=None, filter_row=None):
        if values: res = session.execute(sql, values)
        else: res = session.execute(sql)
        return self.list_result_to_dict(await res, filter_row)

    async def update_values_or_none(self, session, sql, values, commit=False):
        status=await session.execute(sql, values) 
        if commit: await session.commit() ; return status.rowcount
    
    async def insert_class_or_none(self, session, new_obj_class, commit=False):
        session.add(new_obj_class)
        if commit: await session.commit()
        return DB.to_dict(new_obj_class)


# async def init_models():
#     async with engine.begin() as conn:
#         await conn.run_sync(Base.metadata.drop_all)
#         await conn.run_sync(Base.metadata.create_all)

