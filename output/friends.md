---
isarchive: true
comments: true
glightbox: false
hide:
  - footer
  - toc
  - edit
  - view
---

<div class="grid" style="display: grid;grid-template-columns: 32% 33% 32%;" markdown>

<div class="grid cards" style="display: grid; grid-template-columns: 1fr;" markdown>

-   :material-archive-plus:{ .lg .middle } __最近归档__

    ---

    待更新ww


-   :material-archive-star:{ .lg .middle } __完整归档__

    ---

    待更新ww



</div>

<div class="grid cards" markdown>

-   :material-star-face:{ .lg .middle } __社区推荐__

    ---

    待更新ww


</div>

<div class="grid cards" markdown>

-   :material-account-group:{ .lg .middle } __战队招新__

    ---

    待更新ww


</div>

</div>

<div class="grid cards" markdown>

-   :octicons-people-24:{ .lg .middle } __师傅们__

    ---
    - [ek1ng｜Hidden Gem](https://ek1ng.com/)
    - [曾哥｜弱小和无知不是生存的障碍，傲慢才是！](https://blog.zgsec.cn/)
    - [Y4tacker｜宁静致远，淡泊明志](https://y4tacker.github.io)
    - [4ra1n｜许少！](https://4ra1n.github.io/)
    - [crazymanarmy｜A Noob's Learning Record](https://crazymanarmy.github.io/)
    - [xia0ji233｜Nepnep team](https://xia0ji233.pro/)
    - [Steven Lynn's Blog｜Steven的个人博客](https://blog.stv.lol)

</div>
<div class="grid cards" markdown>

-   :fontawesome-solid-blog:{ .lg .middle } __最近更新__

    ---
    ### [某系统前台组合拳RCE](https://y4tacker.github.io/2025/02/23/year/2025/%E6%9F%90%E7%B3%BB%E7%BB%9F%E5%89%8D%E5%8F%B0%E7%BB%84%E5%90%88%E6%8B%B3RCE/)  
    >by [Y4tacker](https://y4tacker.github.io), 2025-02-23

    8133babd8f05e144d8e8c2c4c8ee0fb3d40c2c031c3925709b74cb30c7b27fca92bfd93cf7cf65a6834a805821993b6c9674428ad9fdcf275cf47eb5e07cfd2d776e01401bdf1c9d1052d24b573abb3ff41508c0d801b0496ab9e4257b11885ae0e4bf...
    ### [windowsAPC学习（2）——APC挂入与执行](https://xia0ji233.github.io/2025/02/17/WindowsAPC2/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-02-17

    今天详细了解一下 APC 到底是怎么工作的。APC挂入在 APC 挂入的时候，内核会准备一个 _KAPC 结构体，将该结构体挂入线程的 APC 队列中。KAPC结构体介绍先在 windbg 中查看一下：123456789101112131415161718kd> dt _KAPCntdll!_KAPC   +0x000 Type             : UChar   +0x001 Spare...
    ### [某二次元开放世界冒险游戏反作弊分析报告](https://xia0ji233.github.io/2025/02/14/Game6/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-02-14

    好久没碰某二次元开放世界冒险游戏了，听说新升级了反作弊，故来一探究竟，并尝试实现一些简单的功能。基本保护分析这种级别的游戏首先不考虑静态分析，直接跑起来。不出意外肯定不能直接内存读写，想附加调试器也是附加不上的，所以选择先从驱动入手，游戏加载时会加载驱动。先尝试简单的拦截，方法很多：注册 LoadImage 回调拦截，改驱动名等等等。后者比较好实现，但是运行游戏一段时间会弹窗强制退出。而如果说让保...
    ### [windowsAPC学习（1）——APC简介](https://xia0ji233.github.io/2025/02/09/WindowsAPC1/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-02-09

    来了解一下Windows的APC机制。APCAPC介绍APC 即 Asyncroneus Procedure Call，异步过程调用。学过之前的知识我们知道，线程是不能被杀掉、挂起和恢复的，线程在执行的时候自己占据着CPU，其他线程如何控制它呢？改变一个线程的行为，这就需要APC了。APC结构体APC的结构体如下所示123456789101112131415161718kd> dt _KAPCnt...
    ### [再见2024 | 幸运且美好](https://www.ek1ng.com/Goodbye2024.html)  
    >by [ek1ng](https://ek1ng.com/), 2025-01-28

    再见2024 | 幸运且美好记得23年末时，大四上的我刚结束秋招，在多个offer的抉择中选择来到字节。初到沪城，新的城市，新的公司，新的出租屋，新的工作内容，一切都是新的开始，那时，我写下了再见2023。而如今一年过去，我已回到再熟悉不过的杭州，从一个临近毕业的求职大学生转变成日复一日工作的社畜，我想，这一年，有太多变化。23年许下的期愿 谈恋爱在临近毕业时很幸运的认识了同校的女友，她很温柔，很...
    ### [windows句柄表学习（1）](https://xia0ji233.github.io/2025/01/26/WindowsObjectTable1/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-26

    来了解一下Windows内核的句柄表。句柄句柄就类似 Linux 的文件描述符，指示了某个进程在内核对象的偏移，内核可以通过这个下标找到对应的内核对象。1234HANDLE g_hMutex = ::CreateMutex( NULL , FALSE, "XYZ");HANDLE g_hMutex = ::OpenMutex( MUTEX_ALL_ACCESSFALSE, "XYZ");HANDL...
    ### [windows进程与线程学习——深入研究线程调度（2）](https://xia0ji233.github.io/2025/01/25/WindowsProcess5/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-25

    深入研究一下线程调度，由于篇幅较多，分章节分析，第二篇。进程挂靠一个进程可以包含多个线程，线程结构体中会指向自己所属的进程。切换到这个线程的时候，会将对应的 cr3 切换到该进程的页目录基址，那么这个线程就可以访问这个进程的所有资源了。前面逆向的时候看到过，在切换 cr3 的时候，是拿到了 KTHREAD.ApcState.Process，而并不是 KTHREAD.Process，这个因为没学 A...
    ### [windows进程与线程学习——深入研究线程调度（1）](https://xia0ji233.github.io/2025/01/25/WindowsProcess4/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-24

    深入研究一下线程调度，由于篇幅较多，分章节分析SwapContext首先研究一下 SwapContext 函数的实现。伪代码分析这里我们不去分析汇编代码，而是直接用 IDA + F5，把定义还原回去即可清晰地看出逻辑12345678910111213141516171819202122232425262728293031323334353637383940414243444546474849505...
    ### [windows进程与线程学习——调度实现的学习](https://xia0ji233.github.io/2025/01/24/WindowsProcess3/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-24

    来学习一下 Windows 线程调度的实现Windows中有一个函数SwapContext用来实现线程切换，我们先得了解一下什么情况下会引发线程切换。线程切换途径分主动切换和被动切换，很好理解，主动切换就是线程主动让出 CPU 执行，被动就是被打断而不得不让出 CPU。在做实验的时候，多核真的是很困扰的一个问题，想了半天想不明白多核怎么工作的，为了好理解线程切换，建议虚拟机都换成单核的，线程大多数...
    ### [windows进程与线程学习——调度相关结构学习](https://xia0ji233.github.io/2025/01/24/WindowsProcess2/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-24

    来学习一下调度相关的结构很早就听说过断链隐藏的操作，因为 Windows 都是使用链表去管理进程，线程等结构的，所以断链可以达到隐藏自身的目的。那么这里就引申出来一个问题，为什么断链可以隐身且不破坏大部分的功能呢，下面的线程调度会给出答案。线程调度操作系统的一些理论，线程有三种状态：就绪（ready）、等待（wait）、运行（running）。至于为什么进程/线程断链可以达到隐藏且继续执行的目的，...
    ### [windows进程与线程学习——基本结构](https://xia0ji233.github.io/2025/01/24/WindowsProcess1/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-24

    来学习一下进程与线程的结构EPROCESS定义描述先来看看结构体的描述123456789101112131415161718192021222324252627282930313233343536373839404142434445464748495051525354555657585960616263646566676869707172737475767778798081828384858687...
    ### [windows系统调用学习——系统描述符表](https://xia0ji233.github.io/2025/01/24/WindowsSyscall4/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-24

    来学习一下系统描述符表这个结构SSDTSSDT的全称是System Services Descriptor Table，意为系统服务描述符表。我们可以通过ETHREAD结构体加偏移的方式进行访问。在内核文件中，有一个变量是导出的：KeServiceDescriptorTable。通过它我们可以访问SSDT。可以看看在内核中看看 SSDT 是什么样的。123456789kd> dd KeServic...
    ### [windows系统调用学习——调用细节与系统服务表](https://xia0ji233.github.io/2025/01/22/WindowsSyscall3/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-22

    来深入挖掘一下Windows系统调用的过程KiSystemService分析这个函数是通过中断门进的，中断门本身保存了 CS 和 EIP，跨段提权后通过 TSS 拿到零环的 SS 和 ESP。此时为了维护三环的上下文状态，则会将各种寄存器保存到堆栈，也就是 Trap_Frame 结构体，中断门提权之后本身就会按顺序压入 SS，ESP，ELFAGS，CS，EIP。此时比较一下上一篇文章中提到的 Tr...
    ### [windows系统调用学习——调用相关结构体学习](https://xia0ji233.github.io/2025/01/22/WindowsSyscall2/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-22

    来深入挖掘一下Windows系统调用的过程相关结构体介绍Trap_Frame首先第一个要讲的是 Trap_Frame 结构，如下图所示。栈帧结构体，用于 Windows API 保存现场。经过提权进入0环的时候，Windows就会遵守这个结构体保存一系列的数据，最后四个成员用于虚拟8086模式下，不属于保护模式的范畴。中断发生时，若发生权限变换，则要保存旧堆栈，CPU压入的，由 HardwareE...
    ### [windows系统调用学习——R3到R0](https://xia0ji233.github.io/2025/01/21/WindowsSyscall1/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2025-01-21

    来学习一下windows的系统调用简介API，应用程序接口（Application Programming Interface）。Windows API 顾名思义就是 Windows 提供的应用程序接口，为了在 Windows 上实现一定的功能，我们需要学习 API 的用途、传参、返回值等，微软对这部分都提供了大量的文档说明，因此详细学习 API 的实现原理对我们而言是很有必要的。Windows...
    ### [2024不是年终总结的总结](https://y4tacker.github.io/2024/12/31/year/2024/12/2024%E4%B8%8D%E6%98%AF%E5%B9%B4%E7%BB%88%E6%80%BB%E7%BB%93%E7%9A%84%E6%80%BB%E7%BB%93/)  
    >by [Y4tacker](https://y4tacker.github.io), 2024-12-31

    2024不是年终总结的总结今年总体而言都和计划差不多走了下去，差不多做到了心若止水，心里基本没什么情绪波动了博客里不想分享过多的生活(娱乐活动多，玩得很爽🤪)让博客一直纯粹下去，简单总结下今年的技术进展吧博客博客上差不多算是“大满贯”吧，差不多每个月都坚持写了点东西Github与漏洞研究Github也罕见的做到了每月有更新🤪说到这当然也就离不开年初定下的flag🚩，不然没事谁天天往Github上P...
    ### [2024年终总结](https://xia0ji233.github.io/2024/12/31/Summary2024/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2024-12-31

    2024年也快过去了，过去的目标完成了多少，现在又如何呢？2024总结2024 01-05在队内师傅的带领下，开始写 LLVM 的插件，持续开发到六月份，也算给自己计算机生涯积累了一个大的项目开发经验。01-14时隔两年，再次走上 EC Finals 的舞台，感谢老师们坚持不懈的努力拿到了来之不易的一个名额，最后遗憾打铁，但是这趟上海之行依然很开心。附一段当时写的半退役感言：人生中最后一场以大学生...
    ### [Apache Struts2 文件上传逻辑绕过(CVE-2024-53677)(S2-067)](https://y4tacker.github.io/2024/12/16/year/2024/12/Apache-Struts2-%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E9%80%BB%E8%BE%91%E7%BB%95%E8%BF%87-CVE-2024-53677-S2-067/)  
    >by [Y4tacker](https://y4tacker.github.io), 2024-12-16

    Apache Struts2 文件上传逻辑绕过(CVE-2024-53677)(S2-067)前言​    Apache官方公告又更新了一个Struts2的漏洞，考虑到很久没有发无密码的博客了，再加上漏洞的影响并不严重，因此公开分享利用的思路。分析影响版本Struts 2.0.0 - Struts 2.3.37 (EOL), Struts 2.5.0 - Struts 2.5.33, Struts...
    ### [强网杯S8决赛Reverse writeup](https://xia0ji233.github.io/2024/12/11/qwb2024_final_reverse/)  
    >by [xia0ji233](https://xia0ji233.pro/), 2024-12-11

    复盘一下强网决赛的Reverse题。S1mpleVM附件下载题目名字已经很明显的告诉你了，就是 vm 逆向。基本分析入口其实没啥，就是输入 32 长度的 passcode 然后校验，启动方式是 ./secret_box.exe quest 命令行传参。可以找到最关键的函数 sub_140001D30 就是 VM 入口。这个函数里面很明显的 vm_handler1234567891011121314...
    ### [干货满满之2024广州补天城市沙龙有感](https://blog.zgsec.cn/archives/613.html)  
    >by [曾哥](https://blog.zgsec.cn/), 2024-12-08

    前言哈哈，好久没在博客上面更新文章了，许多朋友（不管是线上还是线下）都在问我是不是不更新博客了，也在催我赶紧更新一下博客的内容，在此也非常感谢各位师傅和朋友的关注和支持~近期不更新博客，原因如下：最近大家也知道emmmmm..我最近到深圳这边，这边项目的对抗程度很高，在这段时间学到很多知识，但工作强度也比较大，人也相比之前时间更少了，也累了不少；但能学到技术还是非常开心的，有时候工作确实太累了，导...

</div>
