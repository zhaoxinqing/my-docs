# GIT

## 撤销上次commit

git reset --soft HEAD^

## 配置upstream

git remote -v
git remote add upstream git@gitlab.clouddeep.cn:lnfra/deepspa-admin.git
<!-- git branch --unset-upstream -->

git fetch upstream
git merge upstream/dev

## git config

git config --global user.name "zhaoxinqing"
git config --global user.email "xinqing.zhao@clouddeep.cn"
git config --global --list
<!-- user.name=zhaoxinqing -->
<!-- user.email=xinqing.zhao@clouddeep.cn -->

## rebase 和 merge 的区别

<https://blog.csdn.net/weixin_41231928/article/details/107040880>
2.回滚代码到某个commit
(1)Git log：获取历史提交信息；
(2)Git reset --hard commited_id（提交记录的唯一标识id）;

optimize: 优化
upgrade: 升级改造
chore: 小改动，一行两行的改动
style: 格式（不影响代码运行的变动）
perf: Performance的缩写，提升代码性能
refactor：重构（既不是新增功能，也不是修改bug的代码变动）
br: 此项特别针对bug号，用于向测试反馈bug列表的bug修改情况
feat：新功能（feature）
fix：修补
docs：文档（documentation）
test：增加测试
revert：feat(pencil): add 'graphiteWidth' option (撤销之前的commit)
bugfix：修补bug
test：新增测试用例或是更新现有测试

## Git 使用相关

git commit之后，撤销commit

git reset --soft HEAD^

- 这样就成功的撤销了你的commit
- 注意，仅仅是撤回commit操作，您写的代码仍然保留。
- HEAD^的意思是上一个版本，也可以写成HEAD~1
- 如果你进行了2次commit，想都撤回，可以使用HEAD~2

## 变基和合并

Git Merge 和 Git Rebase 目的相同，它们都是把不同分支的提交合并到一起。虽然最终目的是一致的，但是其过程却颇为不同。

rebase，合并的结果好看，一条线，但合并过程中出现冲突的话，比较麻烦

- rebase过程中，一个commit出现冲突，下一个commit也极有可能出现冲突，一次rebase可能要解决多次冲突；
- 合并的历史脉络（冲突）被物理消灭了
- merge，合并结果不好看，一堆线交错，但合并有冲突的话，只要解一次就行了；

 **推荐方案：**

- 尽量及时rebase上游分支，发现有冲突，及时merge；
- 在未提交到代码库的本地文件上使用rebase，这样还可以隐藏自己过多的分支细节；
- 随着团队增长，通过 merge 策略很难管理和追踪到每个提交。为了提交历史更清晰、更易于理解，使用 Rebase 是一个明智、高效的选择；

## 常见命令

- git fetch：可以git fetch [alias]取某一个远程repo,也可以git fetch --all取到全部repo；
- git checkout：切换分支；
- git branch：查看分支；
- git log：查看提交log信息；
- git status：查看仓库文件修改状态；

## Github库删除处理办法

故障描述：删除私人仓库误删项目总库，并且原始库不能直接恢复；

1、创建新仓库：

- 在原位置创建新仓库：<https://github.com/spatial-go/geoos>

2、本地代码推送到远程仓库：

- 将本地比较新的个人仓库推送到远程仓库：git push -u upstream main
- git checkout upstream/develop

3、推送标签到远程仓库：

- git push upstream --tags

4、github仓库中根据已有标签创建release；

5、github仓库添加钉钉机器人；

6、github仓库权限设置；

7、github仓库pull request规则设置；

## 统计代码提交行数

git log --since ==2022-05-10 --until=2022-07-22 | wc -l
