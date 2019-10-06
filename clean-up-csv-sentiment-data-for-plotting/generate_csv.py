import pandas as pd
from pymongo import MongoClient

# https://stackoverflow.com/questions/16249736/how-to-import-data-from-mongodb-to-pandas
def _connect_mongo(host, port, username, password, db):
    """ A util for making a connection to mongo """

    if username and password:
        mongo_uri = 'mongodb://%s:%s@%s:%s/%s' % (username, password, host, port, db)
        conn = MongoClient(mongo_uri)
    else:
        conn = MongoClient(host, port)


    return conn[db]


def read_mongo(db, collection, query={}, host='localhost', port=27017, username=None, password=None, no_id=True):
    """ Read from Mongo and Store into DataFrame """

    # Connect to MongoDB
    db = _connect_mongo(host=host, port=port, username=username, password=password, db=db)

    # Make a query to the specific DB and Collection
    cursor = db[collection].find(query)

    # Expand the cursor and construct the DataFrame
    df =  pd.DataFrame(list(cursor))

    # Delete the _id
    if no_id:
        del df['_id']

    return df

def main():
    query = {}
    collection = 'NewsContentScraped'
    db = 'news-scraped-round2'
    host = 'localhost'
    port = 27017
    username = None
    password = None
    no_id = True

    results = read_mongo(db, collection, query, host, port, username, password, no_id)
    results['month'] = results.date.map(lambda x: x.strftime('%Y-%m-01'))
    count = results.groupby(['month'])['month'].agg('count').to_frame('count').reset_index()

    print(count)

if __name__== "__main__":
  main()