from playwright.sync_api import sync_playwright
import os

playwright = sync_playwright().start()
browser = playwright.firefox.launch()
page = browser.new_page()

# Open page
page.goto("https://www.craiyon.com/", wait_until="domcontentloaded")
page.wait_for_selector("//*[@id=\"prompt\"]", state="visible")


#Type prompt
input = page.query_selector('//*[@id="prompt"]')
page.wait_for_timeout(500)
input.fill(os.environ.get('PROMPT'))

#Select style
select = page.query_selector('//*[@id="app"]/div[2]/div/div[2]/div[1]/button[' + str(os.environ.get('STYLE')) + ']')
select.click()

#Click to generate output
button = page.query_selector('//*[@id="generateButton"]')
button.click()
page.wait_for_selector("//*[@id=\"app\"]/div[2]/div/div[3]/div/div/div/div[1]/div[1]/img", state="visible", timeout=90000)


#Grab result
result = []
images = ["//*[@id=\"app\"]/div[2]/div/div[3]/div/div/div/div[1]/div[" + str(i) + "]/img" for i in range(1, 10)]
for i in images:
    img = page.query_selector(i)
    result.append(img.get_attribute('src'))

browser.close()
playwright.stop()

#MOCK FOR TEST
"""
result =['https://img.craiyon.com/2023-08-10/6860e1b3b5b6441f8fb0bf3798c364bd697d88a3.webp',
         'https://img.craiyon.com/2023-08-10/09353a9746314d11a072d35d2f6fc2a7697d88a3.webp',
         'https://img.craiyon.com/2023-08-10/7d306e2a343048819ec873324f79ad78697d88a2.webp', 
         'https://img.craiyon.com/2023-08-10/83b6f80c3b87461ca2733861dbdf8f71697d88a0.webp', 
         'https://img.craiyon.com/2023-08-10/6082bbb414974e179d21176f60233cc7697d88a2.webp',
         'https://img.craiyon.com/2023-08-10/fe40cd3a0a834e8e8231d08ef0722dcd697d88a2.webp',
         'https://img.craiyon.com/2023-08-10/47582ceaee1b4f799d4f94d1f01dd700697d88a3.webp', 
         'https://img.craiyon.com/2023-08-10/ac0ddf2d31c348cdbc2ec8817cc21bcb697d88a0.webp', 
         'https://img.craiyon.com/2023-08-10/4b22ba4afa8747f9a3290dd9cf22eaf9697d88a1.webp']
"""


print(result)