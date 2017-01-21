import scrapy


class Kijiji_Spider(scrapy.Spider):
    name = "kijiji"

    def start_requests(self):

        url_root = 'http://www.kijiji.ca/'

        cells = {
            'motorola' : ['moto G', 'moto Z', 'G4'],
            'apple' : ['iPhone 7', 'iPhone 6S', 'iPhone 6', 'iPhone 5S', 'iPhone 5C', 'iPhone 5', 'iPhone 4S', 'iPhone 4'],
            'blackberry' : ['Curve 9360' , '9720', 'Bold 9900', 'Q5', 'Q10', 'Q20', 'Z10', 'Z30'],
            'htc' : ['M7', 'M8', 'M9'],
            'lg' : ['G3', 'G4'],
        }

        locations = [
            'grand-montreal',
        ]

        for location in locations:
            for brand in cells:
                for model in cells[brand]:
                    url = url_root + '/b-cellulaire/' + location + '/' + brand + '?ad=offering'
                yield scrapy.Request(url=url, callback=self.parse)

    def parse(self, response):
        page = response.url.split("/")[-2]
        filename = 'quotes-%s.html' % page
        with open(filename, 'wb') as f:
            f.write(response.body)
        self.log('Saved file %s' % filename)