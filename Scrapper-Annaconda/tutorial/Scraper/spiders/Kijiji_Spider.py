import scrapy
import logging


class Kijiji_Spider(scrapy.Spider):
    name = "kijiji"

    def start_requests(self):
        apple = ['iPhone SE', 'iPhone 7', 'iPhone 7 Plus', 'iPhone 5S', 'iPhone 6S', 'iPhone 6S Plus']
        samsung = ['Galaxy S6', 'Galaxy S5', 'Galaxy S7', 'Galaxy J3', 'Galaxy S7 Edge']
        lg = ['G3', 'G4', 'G5', 'K4']
        motorola = ['Moto Z', 'Moto G']
        google = ['Pixel', 'Pixel XL', 'Nexus 6P', 'Nexus 6']
        htc = ['M9']
        huawei = ['GR5','Nova Plus', 'Y6']
        phones = {
            'Apple' : apple,
            'Samsung' : samsung,
            'Lg' : lg,
            'Motorola' : motorola,
            'Google' : google,
            'Htc' : htc ,
            'Huawei' : huawei
        }

        for brand in phones:
            for model in phones[brand]:
                safe_brand = brand.replace(' ','-')
                safe_model = model.replace(' ','-')
                url = 'http://www.kijiji.ca/b-cellulaire/grand-montreal/' + safe_brand + '-' + safe_model + '/k0c760l80002?ad=offering'
                request =  scrapy.Request(url=url, callback=self.parse, meta={'brand':brand,'model':model})
                yield request



    def parse(self, response):
        brand = response.meta.get('brand')
        model = response.meta.get('model')
        if brand is None or model is None:
            self.log(response, logging.ERROR)
            return
        objects = response.xpath("//div[@class='info-container']")
        for obj in objects:
            price = obj.xpath(".//div[@class='price']/text()").extract_first()
            title = obj.xpath(".//a[@class='title enable-search-navigation-flag ']/text()").extract_first()
            url = obj.xpath(".//a[@class='title enable-search-navigation-flag ']/@href").extract_first()
            location_coarse = obj.xpath(".//div[@class='location']/text()").extract_first()
            date_posted = obj.xpath(".//span[@class='date-posted']/text()").extract_first()
            description = obj.xpath(".//div[@class='description']/text()").extract_first()
            details = obj.xpath(".//div[@class='details']/text()").extract_first()
            if url is not None:
                url = response.urljoin(url)
            yield {
                    'price': price,
                    'title': title,
                    'url': url,
                    'location_coarse': location_coarse,
                    'date_posted': date_posted,
                    'description': description,
                    'details': details,
                    'brand': brand,
                    'model': model,
                    'brandmodel': brand + ' ' + model,
            }
        pagenum = response.xpath("//span[@class='selected']/text()").extract_first()

        if(pagenum is None):
            self.log('ERROR FINDING NEXT PAGE', logging.ERROR)
            return
        if int(pagenum) < 100:
            # follow pagination links
            next_page_en = response.xpath("//a[@title='Next']/@href").extract_first()
            next_page_fr = response.xpath("//a[@title='Suivante']/@href").extract_first()
            if next_page_en is not None:
                next_page = next_page_en
            else:
                next_page = next_page_fr
            if next_page is not None:
                next_page = response.urljoin(next_page)
                yield scrapy.Request(next_page, callback=self.parse, meta={'brand':brand,'model':model})
            else:
                self.log('Last Page was: ' + pagenum, logging.DEBUG)