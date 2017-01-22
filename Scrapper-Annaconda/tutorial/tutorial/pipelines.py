# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html

import re
from scrapy.exceptions import DropItem
import pymongo
import datetime


class TutorialPipeline(object):
    def process_item(self, item, spider):
        return item

class MongoPipeline(object):
    collection_name = 'kijiji_items'

    def __init__(self, mongo_uri, mongo_db, mongo_root, mongo_account, mongo_user, mongo_pass):
        self.mongo_uri = mongo_uri
        self.mongo_db = mongo_db
        self.mongo_root = mongo_root
        self.mongo_account = mongo_account
        self.mongo_user = mongo_user
        self.mongo_pass = mongo_pass

    @classmethod
    def from_crawler(cls, crawler):
        return cls(
            mongo_uri=crawler.settings.get('MONGO_URI'),
            mongo_root=crawler.settings.get('MONGO_ROOT'),
            mongo_account=crawler.settings.get('MONGO_ACCOUNT'),
            mongo_db=crawler.settings.get('MONGO_DATABASE'),
            mongo_user=crawler.settings.get('MONGO_USER'),
            mongo_pass=crawler.settings.get('MONGO_PASS')
        )

    def open_spider(self, spider):
        self.client = pymongo.MongoClient(self.mongo_root, self.mongo_account)
        self.db = self.client[self.mongo_db]
        self.db.authenticate(self.mongo_user, self.mongo_pass)

    def close_spider(self, spider):
        self.client.close()

    def process_item(self, item, spider):
        self.db[self.collection_name].insert(dict(item))
        return item

class KijijiPipeline(object):

    def process_item(self, item, spider):
        price_string = item['price']
        regex = r"(([0-9]+,*)+)"
        if type(price_string) is not str and type(price_string) is not int:
            raise DropItem("Missing price")
        if re.search(regex, price_string):
            match = re.search(regex, price_string)
            price = re.sub(',', '', match.group(0))
            price = int(price)
        else:
            raise DropItem("Missing price")

        title = item['title']
        url = item['url']
        date_posted = item['date_posted']

        if price < 50:
            raise DropItem("Price too low")
        elif int(price) <= 0 or title is None or url is None:
            raise DropItem("Missing Data")

        if 'Il y a moins de' in date_posted:
            date_posted = datetime.datetime.now()
        else:
            dateArray = date_posted.split('-')
            month = dateArray[1]
            months = dict(janvier=1, fevrier=2, mars=3, avril=4, mai=5, juin=6, juillet=7, aout=8, septembre=9, octobre=10, novembre=11, decembre=12)
            month_num = months[month]
            date_posted = datetime.date(dateArray[2], month_num, dateArray[0])

        return {
            'price': price,
            'title': title.strip(),
            'url': url.strip(),
            'location_coarse': item['location_coarse'].strip(),
            'date_posted': item['date_posted'].strip(),
            'description': item['description'].strip(),
            'details': item['details'].strip(),
            'source': 'kijiji',
        }


class MetaPipeline(object):

    def process_item(self, item, spider):
        title = item['title']
        description = item['description']