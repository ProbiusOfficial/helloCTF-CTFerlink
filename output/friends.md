# My Friends
- [戴兜｜Coding the world.](https://im.daidr.me)
- [Bird｜aFlyBird0](https://blog.aflybird.cn/)
- [NX｜这家伙真勤奋，什么都留下了](https://nickxu.me/)
- [柏喵Sakura｜过气 emo 师傅](https://baimeow.cn/)
- [Marlene｜<归零人生>](https://blog.marlene.top/)
- [ek1ng｜Hidden Gem](https://ek1ng.com/)
- [Atom｜aka asjdf](https://www.homeboyc.cn/)
- [xyxsw｜大相的问题](https://xyxsw.ltd/)

# Recent Blogs
## [『LeetCode-HOT-100』T21～T30](https://nickxu.me/post/leetcode-hot-100-t21-t30.html)  by [NX](https://nickxu.me/), 2023-08-03

全排列板子题，不解释123456789101112131415161718func permute(nums []int) (ans [][]int) {var dfs func(begain, end int)dfs = func(begain, end int) {if begain == end {// 切片是引用类型，需要深拷贝一下tmp := make([]int, len(nums))
## [『LeetCode-HOT-100』T11～T20](https://nickxu.me/post/leetcode-hot-100-t11-t20.html)  by [NX](https://nickxu.me/), 2023-07-30

有效的括号栈的经典题目了属于是12345678910111213141516171819202122232425262728293031323334353637383940func isValid(s string) bool {stack := ""for k := 0; k < len(s); k++ {i := s[k]switch i {case '(':stack = stack + s
## [LLM Agent之结构化输出](https://blog.marlene.top/index.php/develop/88.html)  by [Marlene](https://blog.marlene.top/), 2023-07-26

引言自去年年底以来，GPT的迅速发展诞生了一系列大模型。出现了更新、更大、更强的GPT-4。OpenAI不断推出GPT-4，ChatGPT Plugins，代码解释器，Function calling,图片处理等等。7月的WAIC上，笔者也有幸见到了国内一众企业相继展示自家的大模型。在这段时间里，LLM从最初的PE工程走向智能体交互。而笔者从最开始考虑LLM能不能多人协作，思考”一个专家完成所有任
## [Java RMI 攻击梳理总结](https://www.ek1ng.com/java-rmi.html)  by [ek1ng](https://ek1ng.com/), 2023-07-26

RMI 是什么定义RMI（Remote Method Invocation）是远程方法调用，类似RPC（Remote Procedure Calls）。RPC是打包和传送数据结构，而在Java中，通常传递一个完整的对象，包含数据和操作数据的方法。通过RMI，能够让客户端JVM上的对象，像调用本地对象一样调用服务端JVM上的对象。RMI引入了 Stubs（客户端存根）和 Skeletons（服务端骨
## [重学 Java 反射机制](https://www.ek1ng.com/java-reflect-learning.html)  by [ek1ng](https://ek1ng.com/), 2023-07-25

近期跟一些java的最新漏洞，发现自己的语言基础太差了，跟着p牛的java安全漫谈重新学一下反射，p牛的文章确实是讲复杂的东西讲的浅显易懂。反射的定义对象可以通过反射获取对应的类，类可以通过反射获取所有方法，拿到的方法可以调用，这种机制就是反射。反射机制在安全方面的意义例如我们要完成RCE，但代码中绝大多数时候并没有Runtime，ProcessBuilder等常见的用于命令执行的类来让我们调用。
## [go语言三个小项目 ｜ 青训营笔记](https://xyxsw.ltd/2023/07/25/go%E8%AF%AD%E8%A8%80%E4%B8%89%E4%B8%AA%E5%B0%8F%E9%A1%B9%E7%9B%AE%20%EF%BD%9C%20%E9%9D%92%E8%AE%AD%E8%90%A5%E7%AC%94%E8%AE%B0/)  by [xyxsw](https://xyxsw.ltd/), 2023-07-25

猜数字猜数字这个项目非常简单，它涉及到随机数的生成和用户输入操作。我们使用了bufio库来处理输入数据。reader := bufio.NewReader(os.Stdin)input, _ := reader.ReadString('\n')简单字典标准库strconv它主要用于字符和其他类型之间的转换。strconv.Atoi(s string) int 标准库stringsstrings.T
## [『LeetCode-HOT-100』T1～T10](https://nickxu.me/post/leetcode-hot-100-t1-t10.html)  by [NX](https://nickxu.me/), 2023-07-24

两数之和暴力枚举暴力枚举 i 和 j ，没什么好说的12345678910func twoSum(nums []int, target int) []int {for i, _ := range nums {for j := i + 1; j < len(nums); j++ {if nums[i]+nums[j] == target {return []int{i, j}}}}return ni
## [Golang 性能调优速查笔记](https://homeboyc.cn/blog/golang-%E6%80%A7%E8%83%BD%E8%B0%83%E4%BC%98%E9%80%9F%E6%9F%A5%E7%AC%94%E8%AE%B0/)  by [Atom](https://www.homeboyc.cn/), 2023-07-24

# 技巧 unsafe转换字符串/字节切片的技巧 字符串 -> 字节切片：*(*[]byte)(unsafe.Pointer(&s)) 缓冲区不能修改，否则 go 会panic！ 字节切片 -> 字符串：*(*string)(unsafe.Pointer(&buf)) 重用缓冲区 复位缓冲器 bytes.Buffer.Reset buf = buf[:0] 尽可能直接分配所需大小的数组 清空Map
## [把“用VSCode打开”按钮加入MacOS右键菜单](https://homeboyc.cn/blog/%E6%8A%8A%E7%94%A8vscode%E6%89%93%E5%BC%80%E6%8C%89%E9%92%AE%E5%8A%A0%E5%85%A5macos%E5%8F%B3%E9%94%AE%E8%8F%9C%E5%8D%95/)  by [Atom](https://www.homeboyc.cn/), 2023-07-24

打开自动操作新建文稿选择快速操作设置“工作流程收到当前”为文件或文件夹在左上方搜索栏搜索“运行” -》 找到“运行 Shell 脚本” 并将其拖入右侧 -》 将下方脚本填入文本框for f in "$@"; do open -a 'Visual Studio Code' "$f" done Copy 设置“传递输入”为作为自变量保存为Open in Visual Studio Code
## [java-sec-code 代码审计靶场题解](https://www.ek1ng.com/java-sec-code.html)  by [ek1ng](https://ek1ng.com/), 2023-07-20

这个靶场包含了各类基本漏洞在java语言上的场景以及java安全特有的JNDI注入，反序列化，表达式注入等等，并且给出了相关的利用手段和修复方案。java-sec-code搭建环境可以用Docker搭建，不过想了想不太熟练java的包管理和web server部署这一套，并且本地起相比于容器也方便调试，于是决定本地起一份。由于我是archlinux，包管理安装的都是最新的jdk版本，靶场的jdk版
## [分布式系统测试工具 muxy 初探](https://homeboyc.cn/blog/%E5%88%86%E5%B8%83%E5%BC%8F%E7%B3%BB%E7%BB%9F%E6%B5%8B%E8%AF%95%E5%B7%A5%E5%85%B7-muxy-%E5%88%9D%E6%8E%A2/)  by [Atom](https://www.homeboyc.cn/), 2023-07-17

# 前言 本篇文章主要简单介绍分布式系统测试工具 muxy 工具，在介绍前，我想先谈谈分布式系统测试。# Coding is easier than testing 编程比测试简单。我认为一个程序你把它写出来不是最难的，把它测好才是最难的。真的这么夸张吗？ 我想从一个简单的 HelloWorld 程序来谈谈一个系统的稳定运行所需要的条件。# 一个普通的 HelloWorld 我们需要考虑来自以
## [CrewCTF 2023 Web Writeup](https://www.ek1ng.com/2023CrewCTFWP.html)  by [ek1ng](https://ek1ng.com/), 2023-07-14

环境还在，赛后看看题，一共四道Web，都挺有意思的。sequence_galleryDo you like sequences?http://sequence-gallery.chal.crewc.tf:8080/ 123456789101112131415sequence = request.args.get('sequence', None)if sequence is None:    re
## [使用 GitHub Actions Cache 加快 Workflow](https://blog.aflybird.cn/2023/07/use-github-actions-cache-to-speed-up-workflow/)  by [Bird](https://blog.aflybird.cn/), 2023-07-08

简单讲讲 GitHub Actions 的 Cache 功能，使用方式，以及探讨浅层的设计思想。虽然本文是我博客中为数不多的「教程」类文章，但我还是会侧重逻辑与思考的角度来讲述。详细教程请查看，GitHub Actions Cache 官方文档。
## [『OSPP2023』我与 OSPP 的故事 —— 从听闻到中选](https://nickxu.me/2023/06/my-story-with-OSPP-from-hearing-to-being-selected.html)  by [NX](https://nickxu.me/), 2023-06-30

6月26日下午3点，OSPP2023 中选结果正式发布，全球共有 1486 人成功申请，最终中选人数为 504 人鄙人非常荣幸地成为了这 504 个幸运儿的一份子，特撰此文，记录 我与 OSPP 的故事 —— 从听闻到中选当然，我也希望这篇博客能够吸引更多人参与开源，为开源项目做出自己的贡献给「杭电助手」打点小广告本次开源之夏活动，杭州电子科技大学共中选 13 人，其中 8 人来自杭电助手杭助本次
## [云原生安全分享会材料](https://www.ek1ng.com/cloudsecurity.html)  by [ek1ng](https://ek1ng.com/), 2023-06-28

这是一篇用于给协会小学弟们分享的文章，粗略从各个角度讲了一讲，有任何问题都欢迎联系我交流，email：ek1ng@qq.com。基础知识🧀在开始之前，你需要能够基本掌握Docker和Kubernetes的使用。基本使用推荐看官方文档，配合一些教程动手尝试。https://www.docker.com/Docker 能区分镜像/容器，能基本使用命令，能写Dockerfile，粗略了解原理即可。htt
## [开源之旅——OSPP](https://blog.marlene.top/index.php/develop/85.html)  by [Marlene](https://blog.marlene.top/), 2023-06-27

声明：以下内容仅代表本人观点，具有固有局限性，请辩证看待。任何有问题的地方也恳请指出。开源是什么我一直都认为开源是未来的趋势。用户创作的时代已经到来，作为一名开发者，手里握着强大的武器，个性化的创造自然不在话下。用户创作的目的不是孤芳自赏，而是自媒体式的分享。那么业务、框架的开发与不是自我消化，开源分享或许能够获得更大收益。个人的力量始终有限，我们不是大公司，技术在内部使用也能够逐渐完善。很多内容
## [请还国内开源活动一片净土](https://blog.aflybird.cn/2023/06/please-stop-fucking-open-source-activities-in-china/)  by [Bird](https://blog.aflybird.cn/), 2023-06-26

置顶声明：这篇博客我只发在了自己的独立博客、朋友圈和空间，没有发到其他的平台，没想到会有这么多的阅读量，所以必须要写个声明，保护一下好人：我不想文中涉及的社区，尤其是背后的公司受到损失。我同样也不代表任何团体与公司，不受任何利益。如果真的要对这个社区盖棺定论的话，我觉得这个社区的氛围，是非常学生友好，欢迎开源的。真的非常非常没话说，很热情积极与开放。首先是，社区举办了 学生开发者活动，一
## [Go语言整体替换式Map的Lock-Free实现](https://blog.aflybird.cn/2023/06/go-lock-free-map-with-replace/)  by [Bird](https://blog.aflybird.cn/), 2023-06-16

今天来分享一个写业务过程中，一种特殊场景下的 Lock-Free Map 的实现。 来，我们先抛开看了三遍可能都还看不懂的题目，讲讲故事的背景。 背景/需求 假设我们需要一
## [『算法拾遗』重学主流排序算法](https://nickxu.me/posts/sorting-algorithm.html)  by [NX](https://nickxu.me/), 2023-06-13

衡量排序算法的好坏时间复杂度包含最好情况、最坏情况和平均情况数据有序度不同的影响空间复杂度是否是原地排序稳定性排序后，相同元素之间的顺序是否会改变O( n^2 )冒泡排序（Bubble Sort）依次比较相邻的元素，如果顺序错误，则交换它们。每轮排序将最大（或最小）的元素“冒泡”到正确的位置简单易懂，但效率较低，不适用于大规模数据排序过程初始化待排序数组，设为 arr ，数组长度为 n外层循环：重
## [开源之夏2023申请书——优化CubeFS容器化部署](https://blog.aflybird.cn/2023/06/ospp-2023-proposal-cubefs/)  by [Bird](https://blog.aflybird.cn/), 2023-06-03

因为没确定的细节太多了，以及主要的精力放在了和 Mentor 沟通上。所以申请书写得很粗糙和仓促，后续有机会尽量再写个新版。 以下申请书正文： 一、课题背景与
