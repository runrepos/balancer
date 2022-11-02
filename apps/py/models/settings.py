from sqlalchemy import Column, Integer, String #, ForeignKey, create_engine
from sqlalchemy import select, update, bindparam
from sqlalchemy.sql import text
import json as jsonutils

# find_and_update_sql = text("""UPDATE settings SET value=:value WHERE tag=:tag LIMIT 1""")

def init_class(db):
    class Settings(db.Base):
        ######### Model:
        __tablename__ = "settings"
        id = Column(Integer, primary_key=True)
        tag = Column(String(128))
        value = Column(String(512))
        ###############

        # обновление или insert настроек
        async def edit_settings(body): #session: AsyncSession, 
            async with db.async_session() as session:
                #new_settings = Settings(**body)
                # find = await db.execute_to_dict(
                #     session, 
                #     select(Settings)
                #     .filter(Settings.tag == new_settings.tag)
                #     .limit(1)
                # ) ; # print('found', find)
                # if find:
                if (upd_count:=
                        await db.update_values_or_none(
                                session, 
                                text("""UPDATE settings 
                                        SET value=:value 
                                        WHERE tag=:tag 
                                        LIMIT 1"""), #find_and_update_sql, 
                                [body], 
                                commit=True)): 
                    return { 'updatedCount': upd_count }  #return { 'updatedId': find[0]['id'] }
                elif (res_dict:=
                        await db.insert_class_or_none(
                                session, 
                                Settings(**body), #new_settings, 
                                commit=True)):
                    return { 'insertedId': res_dict['id'] } #db.to_dict(new_settings)['id']}
                return {"error":1}
                #return db.to_dict(new_settings)

        async def get_settings(): # -> list[Settings]:
            async with db.async_session() as session:
                return await db.execute_to_dict(session, 
                                                select(Settings), 
                                                filter_row={"value": lambda x: jsonutils.loads(x) }
                                                )

        # возвращает json настроек для get_settings()
        async def get_formated_settings(filter=None):
            print("Loading settings ...")
            els_dic={}
            for el in await Settings.get_settings():
                els_dic[el['tag']] = el
            return els_dic

    return Settings


