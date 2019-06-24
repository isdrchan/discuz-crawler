# discuz-crawler

### 简介

一个易配置，可扩展的discuz论坛系统的爬虫

- [x] 解析器、数据持久化、调度分离，方便扩展
- [x] 配置goquery(类似jQuery)选择器来获取网页内容
- [x] 配置请求头
- [ ] 关键词过滤
- [ ] 使用cookies爬取
- [ ] 并发爬取
- [ ] 失败重试

### 使用

1. 编译好的二进制文件和配置文件 **config.yaml** 和放在同一个目录下
2. 配置 **config.yaml**
    - **seed**
        - **url** 配置爬取的初始(种子)页面
        - **parser** 配置初始(种子)页面对应的解析器
            - 选项 **forum** 、 **section** 或 **article**，分别对应主页，板块页，文章页。方便对所有板块、单一板块或单一文章进行爬取
    - **selector** 配置选择器，语法几乎与 **jQuery** 一致，方便适配不同的 **discuz** 主题，用于定位爬取的HTML页面上相对应的DOM元素。**article** 定位td标签，其他页面为a标签。
        - **section** 定位主页的“板块”a标签
        - **sub_section** 定位主页的“子板块”a标签
        - **next_page** 定位板块页的“下一页”a标签
        - **title** 定位板块页的“文章标题”a标签
        - **article** 定位文章页的“文章内容”td标签
    - **header** 配置请求头