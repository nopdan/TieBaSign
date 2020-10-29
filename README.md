# tiebasign
使用 github actions 的百度贴吧自动签到
## 使用
Fork 本项目，点击你 Fork 后的仓库右上角的 settings，点击其中的 secrets。

点击 New secret，Name 填 BDUSS，Value 按如下格式填写。

`["bduss"]` 多帐号使用 `["bduss1","bduss2"]` 的格式

bduss 可参照[这里](https://jingyan.baidu.com/article/5552ef47e4358c518ffbc90f.html)获取 

随后点击 Add secret 即可。

之后点击仓库上方的 Actions，点击“I understand my workflows, go ahead and enable them” ，再打开 `README.md` 任意编辑一次，提交。

然后大概会在每天的 0 点和 8 点尝试签到一次。

## 和其他项目的区别
虽然 github 上已经有使用 github actions 来自动签到的项目了，但是我自己觉得不太满意。

比如说有的项目如果签到失败了，就只是打印了错误信息，然后寄希望于第二次签到能成功签上，这样可能 bduss 失效了也无法得知，不就会造成断签了。

所以这个项目在签到发生错误后，会尝试重试几次，如果重试几次后依然无效，就会让 github actions 流程失败，从而发送邮件到你 github 上设置的邮箱中。

此外签到是多线程的，虽然用的不是自己的机器，但是 github 一个月 actions 只有 2000 分钟的免费时长，使用多线程对于关注了大量贴吧的人来说，就能有效的提高速度，从而减少消耗的时长。
