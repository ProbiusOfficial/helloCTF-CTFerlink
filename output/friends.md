---
comments: true

---
# CTFer档案馆
## List
- [ek1ng｜Hidden Gem](https://ek1ng.com/)
- [曾哥｜弱小和无知不是生存的障碍，傲慢才是！](https://blog.zgsec.cn/)
- [Y4tacker｜宁静致远，淡泊明志](https://y4tacker.github.io)
- [4ra1n｜许少！](https://4ra1n.github.io/)
- [crazymanarmy｜A Noob's Learning Record](https://crazymanarmy.github.io/)
- [xia0ji233｜Nepnep team](https://xia0ji233.pro/)
- [Steven Lynn's Blog｜Steven的个人博客](https://blog.stv.lol)

## Recent Post
### [如何判断在IDEA中程序正在运行或正在Debug](https://y4tacker.github.io/2024/01/04/year/2024/1/%E5%A6%82%E4%BD%95%E5%88%A4%E6%96%AD%E5%9C%A8IDEA%E4%B8%AD%E7%A8%8B%E5%BA%8F%E6%AD%A3%E5%9C%A8%E8%BF%90%E8%A1%8C%E6%88%96%E6%AD%A3%E5%9C%A8Debug/)  
>by [Y4tacker](https://y4tacker.github.io), 2024-01-04

如何判断在IDEA中程序正在运行或正在Debug给大家分享一个有趣又无用的东西，如何判断在IDEA中程序正在运行或正在Debug在这个之前我们首先需要了解一个类ManagementFactory ，它是 Java 标准库中的一个类，它提供了访问运行时系统管理接口的工厂方法。通过 ManagementFactory 类，可以获取包括操作系统、内存、线程、类加载器等在内的多种系统管理信息。一些常用的用...
### [再见2023 | 捡起六便士也不忘心中的月亮](https://www.ek1ng.com/Goodbye2023.html)  
>by [ek1ng](https://ek1ng.com/), 2024-01-01

再见2023 | 捡起六便士也不忘心中的月亮写下这篇文章的时候，我刚来到上海入职新公司一周，回看过去一年的竞赛，工作，学习，锻炼等等生活，大体上还是差强人意，但总归有不少值得唏嘘的地方，还得从在北京的生活说起。当时为什么会在北京呢？是去年年末时，有协会学长在群里问有没有学弟想找实习，当时除了CTF以外的经验几乎为零，也从来没有考虑过业界的安全岗都在做什么，我自己想做什么，于是在学长的内推和两轮面试...
### [感谢，渊龙三周年与龙年展望](https://blog.zgsec.cn/archives/573.html)  
>by [曾哥](https://blog.zgsec.cn/), 2024-01-01

1# 概述今天是2024年的第一天，很高兴能再次和大家见面，我是渊龙Sec安全团队的创始人——曾哥 @AabyssZG。首先，在这个日子里面祝各位师傅元旦快乐，在新的一年里面：事业如虎添翼，财运如虹贯日，家庭和谐美满，幸福安康常伴！也很感谢各位师傅，平时给予团队和我的关注和支持~同时，也非常感谢各位团队成员的共同建设和鼎力相助，团队正因为有了你们，才能走到今天！2# 关于渊龙三周年今年是渊龙Sec...
### [2024的目标](https://xia0ji233.github.io/2023/12/31/Summary2023/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-12-30

2023年也快过去了，至此，大学生涯算是快结束了。2023总结2023 01-14经过了很久的申诉，ICPC2023 侥幸拿铜（check）02-01CSAPP看完03-29XCTF FINALS 二等（真的很感谢队里师傅的认可，能让我去参加）04-15ZJCPC 又双叒叕打铁了，哭了04-25自顶向下计算机网络看完05-20CISCN 直接没去（因为一些意外）11-13ZJCTF 二等（单挑省一...
### [又又又是一个属性覆盖带来的漏洞](https://y4tacker.github.io/2023/12/28/year/2023/12/%E5%8F%88%E5%8F%88%E5%8F%88%E6%98%AF%E4%B8%80%E4%B8%AA%E5%B1%9E%E6%80%A7%E8%A6%86%E7%9B%96%E5%B8%A6%E6%9D%A5%E7%9A%84%E6%BC%8F%E6%B4%9E/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-28

又又又是一个属性覆盖带来的漏洞想到最近出了好几个与属性覆盖有关的漏洞，突然想到有一个国产系统也曾经出过这类问题，比较有趣这里简单分享一下，希望把一些东西串起来分享方便学到一些东西前后端框架信息梳理首先简单从官网可以看出所使用的框架信息以及技术选型https://gitee.com/mingSoft/MCMS?_from=gitee_search我们主要关注几个点一个是shiro，一个是freema...
### [Apache OFBiz未授权命令执行浅析(CVE-2023-51467)](https://y4tacker.github.io/2023/12/27/year/2023/12/Apache-OFBiz%E6%9C%AA%E6%8E%88%E6%9D%83%E5%91%BD%E4%BB%A4%E6%89%A7%E8%A1%8C%E6%B5%85%E6%9E%90-CVE-2023-51467/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-27

Apache OFBiz未授权命令执行浅析(CVE-2023-51467)未修复的权限绕过还是之前那个遗留的问题，首先是权限绕过，首先还是做一个简单的回顾关于登录的校验在org.apache.ofbiz.webapp.control.LoginWorker#checkLogin做处理来判断用户是否登录，可以看到这里的判断逻辑非常简单，跳过前两个判断，在后面只需要login返回的不是error，则为...
### [Apusic权限绕过浅析](https://y4tacker.github.io/2023/12/26/year/2023/12/Apusic%E6%9D%83%E9%99%90%E7%BB%95%E8%BF%87%E6%B5%85%E6%9E%90/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-26

Apusic权限绕过浅析真的是浅析前几天去参加补天了，一直想写但是一直抽不出时间学习，由于漏洞比较简单这里也不过多篇幅的讲解，仅分享一些关键的点，在这里关于权限校验Apusic没有使用第三方框架(毕竟是迫真信创产品)而是使用了自定义实现的安全性约束(关于什么安全性约束百度搜很多文章了不作搬运工)去实现了访问控制1234567891011<security-constraint>    <displ...
### [Hacking FernFlower](https://y4tacker.github.io/2023/12/22/year/2023/12/Hacking-FernFlower/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-22

Hacking FernFlower前言​    今天很开心，第一次作为speaker参与了议题的分享，也很感谢补天白帽大会给了我这样的一次机会​    其实本该在去年来讲Java混淆的议题，不过当时赶上疫情爆发，学校出于安全的考虑没让出省。在当时我更想分享的是对抗所有反混淆的工具cfr、procyon，但今年在准备过程中发现主题太大了其实不太好讲，再考虑到受众都是做web安全的，因此我最终还是将...
### [Airpods Pro 2上手体验](https://blog.stv.lol/archives/79/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-12-16

Airpods Pro 2 大概算是我今年买的最后一个产品了，也是我馋了很久却一直没买的产品在此之前，我的主力耳机是 Airpods Pro 一代，考虑到一代已经用了将近三年时间，也是时候换一个了没考虑其他耳机的原因主要还是因为我的主要使用的生态产品还是苹果生态为主，Airpods的无缝连接在各设备之间互相流转很方便于是在tb的双十二活动中以1500出头的价格拿下（耳机本体+一年以换代修）开箱包装...
### [一周年小记&amp;&amp;那些快乐的技术时光](https://blog.zgsec.cn/archives/548.html)  
>by [曾哥](https://blog.zgsec.cn/), 2023-12-14

1# 概述不知不觉，个人博客已经开办了一年了回头看一年前的自己，仍有些感触，遂在闲暇时光提笔写下一些碎碎念数了数我在这一年发表过的博客文章，共计约二十余篇，其实我是真没想到能有那么多文章，比我原定的目标（每个月写一篇原创技术文）多出不少，也算是suprise吧[...]...
### [亿赛通电子文档安全管理系统远程代码执行漏洞浅析](https://y4tacker.github.io/2023/12/13/year/2023/12/%E4%BA%BF%E8%B5%9B%E9%80%9A%E7%94%B5%E5%AD%90%E6%96%87%E6%A1%A3%E5%AE%89%E5%85%A8%E7%AE%A1%E7%90%86%E7%B3%BB%E7%BB%9F%E8%BF%9C%E7%A8%8B%E4%BB%A3%E7%A0%81%E6%89%A7%E8%A1%8C%E6%BC%8F%E6%B4%9E%E6%B5%85%E6%9E%90/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-13

亿赛通电子文档安全管理系统远程代码执行漏洞浅析漏洞分析最近天天曝亿赛通的漏洞，又是这个新手向的项目，有点烦其实不是很想写的，本次原理也很简单熟悉的人可能知道这个系统在windows与linux下有点区别，在linux系统下多了一个8021端口相较于CDGServer3服务下又臭又长的代码，这个fileserver下的代码还是很短小的任意文件读取在com.esafenet.fileserver.co...
### [年轻人第一块电表：酷安人均一只酷态科10号移动电源简评&amp;杂谈](https://blog.stv.lol/archives/78/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-12-11

前言前段时间有幸去了酷科南京总部参观，被问到有没有酷安人均一只的酷态科10号时汗流浃背了，因为一直以来我都在用闪极的产品，唯一的酷态科产品还是酷态科前身紫米的紫米200W移动电源而对于紫米200W移动电源，我的评价是很优秀，但是太大太重不便于日常携带，并且上次出去旅游的时候外壳被摔坏导致只有一个口可以用了于是在两周前淘宝百亿补贴的一次机会以189元的价格赶紧补票了酷态科10号开箱酷态科10号的包装...
### [CrushFTP Unauthenticated Remote Code Execution(CVE-2023-43177)](https://y4tacker.github.io/2023/12/10/year/2023/12/CrushFTP-Unauthenticated-Remote-Code-Execution-CVE-2023-43177/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-10

CrushFTP  Unauthenticated Remote Code Execution路由分析不像传统套件，这里自己实现了协议的解析并做调用，写法比较死板，不够灵活，在crushftp.server.ServerSessionHTTP可以看到具体的处理过程，代码”依托答辩”，不过漏洞思路值得学习前台权限绕过简单来说，原理是因为程序实现存在匿名访问机制，并且可以通过header污染当前会话的...
### [Apache Struts2 文件上传分析(S2-066)](https://y4tacker.github.io/2023/12/09/year/2023/12/Apache-Struts2-%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E5%88%86%E6%9E%90-S2-066/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-09

Apache Struts2 文件上传分析(S2-066)struts2也很久没出过漏洞了吧，这次爆的是和文件上传相关相关的commit在https://github.com/apache/struts/commit/162e29fee9136f4bfd9b2376da2cbf590f9ea163首先从commit可以看出，漏洞和大小写参数有关，后面会具体谈及同时结合CVE描述我们可以知道，大概和...
### [某某通漏洞浅析](https://y4tacker.github.io/2023/12/08/year/2023/12/%E6%9F%90%E6%9F%90%E9%80%9A%E6%BC%8F%E6%B4%9E%E5%88%86%E6%9E%90/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-12-08

e1bd3e68fe0e95aad543147235f13aa8c07db101f08238b29c49348e526ca2c395ff16b5184fe4ee67b8dd1840f69936917a685390975ecb78405622c312487d5e6b45d97870e169a38f37bd7a2e95cc5b864202dba1787e366b00f3973f1aa3686ad8...
### [Linux ptrace](https://xia0ji233.github.io/2023/12/03/Ptrace/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-12-03

这次学习Linux进程调试相关的知识。调试对于二进制选手来说，调试的重要性不言而喻，对于Linux来说，基本就是 gdb 一家独大，其余插件只是给gdb起了锦上添花的一些作用罢了，那么下面就来学习一下 gdb 的内核。ptrace在Linux调试程序，离不开一个系统调用就是 ptrace（%rax=101,%eax=26），来看看这个函数原型：12long ptrace(enum __ptrace...
### [Apache ActiveMQ Jolokia远程代码执行不依赖JDK打法](https://y4tacker.github.io/2023/11/30/year/2023/11/%E6%9F%90%E7%B3%BB%E7%BB%9F%E6%9C%80%E6%96%B0%E5%89%8D%E5%8F%B0RCE%E5%88%86%E6%9E%90/Apache-ActiveMQ-Jolokia%E8%BF%9C%E7%A8%8B%E4%BB%A3%E7%A0%81%E6%89%A7%E8%A1%8C%E4%B8%8D%E4%BE%9D%E8%B5%96JDK%E6%89%93%E6%B3%95/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-11-30

Apache ActiveMQ Jolokia远程代码执行不依赖JDK打法想着最近连写了几篇加密博客有点对不起看我博客的粉丝了，今天抽空简单分享一个姿势影响版本大概测了一下Apache ActiveMQ 5.16.x系列无log4j2的mbeanApache ActiveMQ 5.17.x系列漏洞版本受影响初探从网上已公开的打法可以知道使用jdk.management.jfr:type=Fligh...
### [Apache ActiveMQ Jolokia远程代码执行(CVE-2022-41678)简析及绕Waf技法](https://y4tacker.github.io/2023/11/29/year/2023/11/Apache-ActiveMQ-Jolokia%E8%BF%9C%E7%A8%8B%E4%BB%A3%E7%A0%81%E6%89%A7%E8%A1%8C-CVE-2022-41678-%E7%AE%80%E6%9E%90%E5%8F%8A%E7%BB%95Waf%E6%8A%80%E6%B3%95/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-11-29

b911718fdd890810916fcf22cb8016ca11ae7d872d624941bd9b20d0cf7e4036a45e816cf5955164cc4e0d7bc4e06277d1f41e412de1243b7603f6b8db578947e40f8bad6233f702a2ae224c03f173d620410ea4abfdb0fba73dbe0f45d4327f2d262b...
### [HITCTF2023题解](https://xia0ji233.github.io/2023/11/27/HITCTF2023/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-11-27

重铸mininep荣光，我辈义不容辞战队信息战队名称：mininep本次比赛附件，WEB和Reverse就不放了，主要是 Pwn，MISC和Cry。因为超过了 Github 的上传限制，因此单独放 MISC 的音频题附件Pwnscanf会先读入一个序列，然后根据序列去调用功能。[ ：分配大小为 0x20 的堆块，并读入一个 int 数据。] ：输出堆块的 int 数据并 free 掉堆块并且马上将...
### [某系统最新前台权限绕过分析](https://y4tacker.github.io/2023/11/26/year/2023/11/%E6%9F%90%E7%B3%BB%E7%BB%9F%E6%9C%80%E6%96%B0%E5%89%8D%E5%8F%B0%E6%9D%83%E9%99%90%E7%BB%95%E8%BF%87%E5%88%86%E6%9E%90/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-11-26

f8699a5360fe9f42f41b98c5cc3203c7cbe9b0e7b973590e7fd52b27b474d14b936d11c308462091084b3defadb8d998d3381b897b97ebf2923245170a43a3ab58393c3b23741eb3526a32145ff738802bd1881ac66ab453888465f4f27015647ba7da...
### [计算机网络(谢希仁)](https://xia0ji233.github.io/2023/10/31/ComputerNetwork-Sum/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-10-31

谢书作为考纲，也是需要好好看看的，接下来时间把谢书再过一遍把。首先列一下考纲吧，然后把重点知识列出来。计算机网络体系结构计算机网络概述计算机网络的概念、组成与功能什么是计算机网络？计算机网络是由若干节点和链路组成的互连的网络。计算机网络的组成？硬件软件的视角：是由软件，硬件及其对应的协议所组成。从工作方式来看：可以由核心部分和边缘部分组成（核心部分负责路由互联，边缘部分负责提供服务）。从组成上看：...
### [域名解析系统](https://xia0ji233.github.io/2023/10/30/DNS/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-10-30

来学一学DNS吧。简介域名解析系统采用 客户端/服务器（Client/Server）模型，是一种应用层协议，它的作用是把我们所熟知的域名（domain ）翻译成 ip 地址。主要是人们通常乐意记住域名而不是记住ip，就好比学校里老师喜欢叫名字而不是叫你学号，但是学号又是唯一确定的一个学生，而名字是可以重复的。给出一个学生名称，查询对应学号的一个系统就叫学号解析系统，类似的，给出一个域名，查询一个对...
### [记一次抑郁诊断和开药用药记录](https://blog.stv.lol/archives/76/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-10-12

前因自我判定进入比较明显的焦虑抑郁状态已经至少两年多了，前因可看这篇文章[post cid="58" /]当然更主要的还是远因，一些童年成长时的家庭、环境因素，造成了长期表达能力受到压制，进而失去表达情感的能力，比如不会大笑也不会大哭受到这些因素的影响，开始不断地侵蚀我的自我意识，在高中那会就开始出现记忆力下降和无法专注等情况，但在当时由于并未了解过相关心理学知识并不知道这是长期焦虑抑郁所致症状进...
### [一种基于HSTS的防域名劫持的方法](https://blog.stv.lol/archives/75/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-10-04

事情是这样的：国庆期间国内网络对很多域名进行了污染，其中影响最大的是Minecraft正版验证api以及VSCode官网。当尝试登录或使用上述应用时，其HTTP请求会被重定向到国家反诈中心以及工信部反诈中心的提示页面也因为这波莫名其妙的风波，最近某位朋友的手头大量部署在vercel上的服务被运营商劫持，被301跳转到反诈中心页面对于用户而言，只能切换DNS或者使用代理来防止劫持昨天和一位大佬交流过...
### [香港借记卡：汇丰(HSBC)&amp;中银香港(BOCHK)开卡过程全记录](https://blog.stv.lol/archives/74/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-10-04

前段时间有幸抢到了去香港的免费机票，于是顺路去香港开了中国银行香港和汇丰香港的账户本文就仅作一个开卡全过程记录吧汇丰(HSBC)周六那天经过了三家中银都没有成功，抱着一丝希望来到了HSBC葵芳分行，试图碰碰运气。当时并没有进行预约。入门后，我先询问前台是否还有号码。因为我没有预约，所以实际上是碰运气。结果，前台小姐姐非常客气地把我带到了一个办公小隔间门口，并告诉我等前一位先生结束之后就可以进去（没...
### [Flask的部署](https://xia0ji233.github.io/2023/10/01/Flask-deploy/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-10-01

最近在做毕业设计的时候搞了一下，记录一下。环境在使用 app run 去跑 flask 项目的时候，会出现这么一句话：WARNING: This is a development server. Do not use it in a production deployment.大概意思就是，不要用这种方式部署到生产环境中，之前我倒是直接就是这么部署的，但是明显会感觉部分操作会发生卡顿，因此这里顺着...
### [Fourier-Serials](https://xia0ji233.github.io/2023/09/15/Forier-Serials/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-09-15

Today study something about Fourier serials Taylor SerialAs we all know,Taylor formula give us the way for expressing complicated function as simple polynomial function.We should clearly remember the...
### [Markdown editor CVE of Marktext](https://www.ek1ng.com/MarkdownEditorCVE1.html)  
>by [ek1ng](https://ek1ng.com/), 2023-09-13

无CVE编号 XSS2RCEhttps://github.com/marktext/marktext/issues/2601https://github.com/marktext/marktext/commit/0dd09cc6842d260528c98151c394c5f63d733b62影响 <= 0.16.3 的marktext版本，点击链接触发。POC：123...
### [Probability Theory（5）](https://xia0ji233.github.io/2023/09/11/Probability-Theory5/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-09-11

another question about confidence interval. ProblemToday I got a problem about the difference of two normal distribution variables of confidence interval. Obviously it is much harder than  the ratio b...
### [概率论（4）](https://xia0ji233.github.io/2023/08/31/Probability-Theory4/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-08-31

一道压轴题引发的思考~题目分析题目如下：这个题目就这么短，而且也非常的清晰，却是 22 年的选择压轴题，这题深入做过之后发现常规的很多做法都需要记住一些公式或者是较大的计算量（答案确实也如此）。而在这里我感觉到可以使用另外一种比较简单的思路做出这题。解题首先从题目描述提取信息之后得到以下两个比较重要的信息。$X \sim N(0,1),f_X(x)=\frac{1}{\sqrt{2\pi}}e^{...
### [渗透测试思路总结](https://www.ek1ng.com/Summary%20of%20penetration%20ideas.html)  
>by [ek1ng](https://ek1ng.com/), 2023-08-29

最近做了一阵子攻防相关的事，正好最近国护结束，做个总结，简单写一下渗透的基本思路（Check List）。不同的标题间内容并不完全独立，在实战中，比如先钓鱼获取到一台个人PC，但这台PC并不在办公网。而后通过收集个人PC的信息，能够登陆外网其他站点的后台，配合一个后台RCE进入办公网/生产网。这其中就有钓鱼，也有外网打点的部分。资产收集资产搜集通俗说就是“了解目标有什么东西”，讲究一个越全越好。路...
### [Telegram Instant View即时预览适配过程记录](https://blog.stv.lol/archives/72/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-08-27

Telegram有个即时预览的功能，可以以阅读格式的页面形式加载网页，减少加载不必要的内容Telegram自己的Blog，以及主流内容平台（比如wikipedia, Medium, 知乎都有人做了适配）但直到今天一番研究才知道适配工作要在instant view上面写规则，并且提交审核这个工作可以说基本上由用户维护，官方通过排名的方式来选择每个网站最好的Instant View模板，也有2017和...
### [英国铁路系统体验小结](https://blog.stv.lol/archives/67/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-08-27

今年暑假去了趟英国，在英国生活了半个月的时间打算分几篇文章简述一下一些感受这篇文章主要讲讲英国的铁路系统（包括地铁和火车）文章偏体验向，可能比较流水账，想到哪写到哪概述因为起步时间早，英国铁路系统相当发达，早在19世纪就已经开始建成世界第一条地铁线路，后续又建了许多线路并不断完善城市铁路网络在伦敦市区，火车与地铁共用线路，可以凭oyster进站上下车，速度也与地铁速度一致，离开共用线路之后恢复至正...
### [浅析Smartbi逻辑漏洞(2)](https://y4tacker.github.io/2023/08/23/year/2023/8/%E6%B5%85%E6%9E%90Smartbi%E9%80%BB%E8%BE%91%E6%BC%8F%E6%B4%9E-2/)  
>by [Y4tacker](https://y4tacker.github.io), 2023-08-22

浅析Smartbi逻辑漏洞(2)写在前面仅分享逻辑漏洞部分补丁绕过思路，不提供完整payload厂商已发布补丁：https://www.smartbi.com.cn/patchinfo正文简单提一下，补丁部分由smartbi.security.patch.PatchFilter(来源于SecurityPatchExt.ext)做加载并处理，这里我们主要关注补丁返回的状态码的具体含义即可，可以看到只...
### [概率论（3）](https://xia0ji233.github.io/2023/08/09/Probability-Theory3/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-08-09

参数估计和置信区间计算。引入很多老师都用这样的一个例子引入，假设我想知道全国身高的均值。这个值 $\mu$ 必然是客观存在的，但是给全国十四亿人全部测一遍显然这个工作量太大，对于学过数理统计的人来说，解决这个的最好的办法自然就是抽样。理论极限中心定理告诉我们，当样本量足够大时，所有的分布最终都趋于正态分布，下图就是标准正态分布的概率密度图。概率密度函数反映了概率在该点的变化率，变化率越大显然证明该...
### [游戏安全的学习（4）](https://xia0ji233.github.io/2023/08/01/Game4/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-08-01

游戏安全的学习（4）—— 猫里奥外挂编写。这也就是一个单机游戏，如果游戏作者觉得侵权了可以联系删除。游戏分析突然想起了之前一个很搞人心态的游戏，就是这个猫里奥，各种陷阱防不胜防，因此今天打算用点科技的力量去对抗这个东西。八关版的外挂已经有人做了，这里我用九关版的，下载地址。坐标读取和修改首先最简单的就是获取坐标，使用CE反复横跳搜索内存，最终确定了 X 和 Y 坐标的地址分别是在 +78DDC 和...
### [java-sec-code 靶场题解](https://www.ek1ng.com/java-sec-code.html)  
>by [ek1ng](https://ek1ng.com/), 2023-07-20

这个靶场包含了各类基本漏洞在java语言上的场景以及java安全特有的JNDI注入，反序列化，表达式注入等等，并且给出了相关的利用手段和修复方案。java-sec-code搭建环境可以用Docker搭建，不过想了想不太熟练java的包管理和web server部署这一套，并且本地起相比于容器也方便调试，于是决定本地起一份。由于我是archlinux，包管理安装的都是最新的jdk版本，靶场的jdk版...
### [计算机网络复习（3）](https://xia0ji233.github.io/2023/07/17/408-2-3/)  
>by [xia0ji233](https://xia0ji233.pro/), 2023-07-17

数据链路层数据链路层的功能在物理层的基础上向网络层提供服务，对网络层表现为一条无差错的逻辑链路。提供的服务数据链路层可能提供以下的服务：成帧：链路层封装IP数据报，帧的结构由链路层协议规定。链路接入：媒体访问控制协议（Medium Access Control，MAC）规定了帧在链路上传输的规则.可靠交付：如果链路层实现可靠交付，通常是通过确认和重传取得的。对于高比特差错的链路来说，链路层通常要实...
### [CrewCTF 2023 Web Writeup](https://www.ek1ng.com/2023CrewCTFWP.html)  
>by [ek1ng](https://ek1ng.com/), 2023-07-14

环境还在，赛后看看题，一共四道Web，都挺有意思的。sequence_galleryDo you like sequences?http://sequence-gallery.chal.crewc.tf:8080/ 123456789101112131415sequence = request.args.get('sequence', None)if sequence is None:    re...
### [旅行日志：无锡&amp;苏州小记](https://blog.stv.lol/archives/61/)  
>by [Steven Lynn's Blog](https://blog.stv.lol), 2023-07-09

在苏北那个鬼地方关了两年多，尤其之前一直在被各种考试折腾，终于放暑假了，于是临时起意去长三角几个城市玩两天一个人出行最大的好处在于自由度，可以想到哪里就去哪里以及，和朋友们面基！这篇文章想到哪写哪，因此全文可能看起来比较流水账多图警告Day1 无锡说起来以前经常从无锡的高速或者火车站经过，但从来没有真正到过无锡，所以这一次来算是第一次玩先说下无锡的印象吧，个人觉得无锡这个城市既现代化又不怎么现代化...
