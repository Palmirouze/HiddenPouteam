import scrapy


class Kijiji_Spider(scrapy.Spider):
    name = "kijiji"

    def start_requests(self):

        categories = [
            'b-cellulaire',
        ]

        cell_brands = [
            'motorola',
            'apple',
            'blackberry',
            'htc',
            'lg',
            'ericsson',
            'nokia',
            'samsung',
            'sony',
            'sony+ericsson',
        ]

        locations = [
            'grand-montreal',
        ]

        params = [
            '?ad=offering',
        ]

        for location in locations:
            for cell_brand in cell_brands:

        urls = [
            'http://quotes.toscrape.com/page/1/',
            'http://quotes.toscrape.com/page/2/',
        ]
        for url in urls:
            yield scrapy.Request(url=url, callback=self.parse)

    def parse(self, response):
        page = response.url.split("/")[-2]
        filename = 'quotes-%s.html' % page
        with open(filename, 'wb') as f:
            f.write(response.body)
        self.log('Saved file %s' % filename)