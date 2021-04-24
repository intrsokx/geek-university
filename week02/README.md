## go训练营第二周作业
***
#### Question
>我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

#### Answer
>我认为在dao层将一些便于定位问题的信息（比如SQL语句）wrap一下，然后向server层抛出比较好。

***
## 总结：
#### 在我的代码逻辑中，分为handler、dao、entity三层。各层及功能如下

* handler: 处理业务逻辑，做错误的具体捕获者(catcher)
* dao: dao层呢，作为查询数据库的实际操作者，如果有错误的话，这一层是最先感触到的，然后它只细分一下是什么错误，然后把错误封装，上抛。（wrapping && return）
* entity: entity层是数据库表到程序的一个映射(mapping)。