## go训练营第二周作业
***
#### Question
>我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

#### Answer
>我认为在dao层将一些便于定位问题的信息（比如SQL语句）wrap一下，然后向server层抛出比较好。
