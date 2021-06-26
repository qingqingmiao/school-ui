def print_hi():
    # --*--conding:utf-8 --*--
    # Author: Gonggong
    # 使用python爬取一个网页中表格的内容，并把抓取到的内容以json格式保存到文件中

    import requests
    from lxml import etree
    import json

    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36',
        'Referer': 'https: // www.cingta.com / school / ser'
    }
    # 获取网页源代码
    left = requests.get(url='https://www.cingta.com/school/api/area_list/', headers=headers)
    print(json.loads(left.content)["data"]["list"])
    
    r = requests.get(url='https://www.cingta.com/school/api/name_uni_list/', headers=headers)
    a = json.loads(r.content)["data"]["list"]
    for i in a:
        print(i)


    # 使用xpath对爬取的源代码进行处理
    # dom_tree = etree.HTML(r.content)
    # links = dom_tree.xpath('//*[@id="app_1"]/div/div[2]/div[2]/div/ul/li[1]/div/span/span')
    # print(links)
    # for i in links:
    #     print(i.text)
    # 取出links的单行、双行的数据
    # res1 = [i.text for i in links[::2]]
    # res2 = [i.text for i in links[1::2]]

    # 把两行数据组合成在一起
    # result = tuple(zip(res1, res2))
    # print(result)
    # 使用json格式保存到文件中
    # json.dump(result, open('./xpath_get.txt', 'w'), ensure_ascii=True)


if __name__ == '__main__':
    print_hi()
