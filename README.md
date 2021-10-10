## 前言

在面向对象编程的理念里，应用程序是对现实世界的抽象，我们经常会将现实中的事物建模为编程语言中的类/对象（“**是什么**”），而事物的行为则建模为方法（“**做什么**”）。面向对象编程有**三大基本特性**（封装、继承/组合、多态）和**五大基本原则**（单一职责原则、开放封闭原则、里氏替换原则、依赖倒置原则、接口分离原则），但知道这些还并不足以让我们设计出好的程序，于是很多方法论就涌现了出来。

近来最火的当属领域驱动设计（DDD），其中战术建模提出的实体、值对象、聚合等建模方法，能够很好的指导我们设计出符合现实世界的领域模型。但DDD也不是万能的，在某些应用场景下，按照传统的战术建模/面向对象方法设计出来的程序，也会存在可维护性差、违反单一职责原则等问题。

本文介绍的**DCI建模方法**可以看成是战术建模的一种辅助，在某些场景下，它可以很好的弥补DDD战术建模的一些缺点。接下来，我们将会通过一个案例来介绍DCI是如何解决DDD战术建模的这些缺点的。

> 本文涉及的代码归档在github项目[DCI-Architecture-Implementation](https://github.com/ruanrunxue/DCI-Architecture-Implementation)：https://github.com/ruanrunxue/DCI-Architecture-Implementation

## 案例

考虑一个普通人的生活日常，他会在学校上课，也会趁着暑假去公司工作，在工作之余去公园游玩，也会像普通人一样在家吃喝玩乐。当然，一个人的生活还远不止这些，为了讲解方便，本文只针对这几个典型的场景进行建模示例。

![](https://tva1.sinaimg.cn/large/008i3skNgy1gulxnp2uryj61400l6q7202.jpg)

## 使用DDD建模

按照DDD战术建模的思路，首先，我们会列出该案例的**通用语言**：

> 人、身份证、银行卡、家、吃饭、睡觉、玩游戏、学校、学生卡、学习、考试、公司、工卡、上班、下班、公园、购票、游玩

接着，我们使用战术建模技术（**值对象**、**实体**、**聚合**、**领域服务**、**资源库**）对通用语言进行领域建模。

DDD建模后的代码目录结构如下：

```yaml
- aggregate: 聚合
  - company.go
  - home.go
  - park.go
  - school.go
- entity: 实体
  - people.go
- vo: 值对象
  - account.go
  - identity_card.go
  - student_card.go
  - work_card.go
```

我们将身份证、学生卡、工卡、银行卡这几个概念，建模为**值对象**（Value Object）：

```go
package vo

// 身份证
type IdentityCard struct {
	Id   uint32
	Name string
}

// 学生卡
type StudentCard struct {
	Id     uint32
	Name   string
	School string
}

// 工卡
type WorkCard struct {
	Id      uint32
	Name    string
	Company string
}

// 银行卡
type Account struct {
	Id      uint32
	Balance int
}

...
```

接着我们将人建模成**实体**（Entity），他包含了身份证、学生卡等值对象，也具备吃饭、睡觉等行为：

```go
package entity

// 人
type People struct {
	vo.IdentityCard
	vo.StudentCard
	vo.WorkCard
	vo.Account
}

// 学习
func (p *People) Study() {
	fmt.Printf("Student %+v studying\n", p.StudentCard)
}
// 考试
func (p *People) Exam() {
	fmt.Printf("Student %+v examing\n", p.StudentCard)
}
// 吃饭
func (p *People) Eat() {
	fmt.Printf("%+v eating\n", p.IdentityCard)
	p.Account.Balance--
}
// 睡觉
func (p *People) Sleep() {
	fmt.Printf("%+v sleeping\n", p.IdentityCard)
}
// 玩游戏
func (p *People) PlayGame() {
	fmt.Printf("%+v playing game\n", p.IdentityCard)
}
// 上班
func (p *People) Work() {
	fmt.Printf("%+v working\n", p.WorkCard)
	p.Account.Balance++
}
// 下班
func (p *People) OffWork() {
	fmt.Printf("%+v getting off work\n", p.WorkCard)
}
// 购票
func (p *People) BuyTicket() {
	fmt.Printf("%+v buying a ticket\n", p.IdentityCard)
	p.Account.Balance--
}
// 游玩
func (p *People) Enjoy() {
	fmt.Printf("%+v enjoying park scenery\n", p.IdentityCard)
}

```

最后，我们将学校、公司、公园、家建模成**聚合**（Aggregate），聚合由一个或多个实体、值对象组合而成，组织它们完成具体的业务逻辑：

```go
package aggregate

// 家
type Home struct {
	me *entity.People
}
func (h *Home) ComeBack(p *entity.People) {
	fmt.Printf("%+v come back home\n", p.IdentityCard)
	h.me = p
}
// 执行Home的业务逻辑
func (h *Home) Run() {
	h.me.Eat()
	h.me.PlayGame()
	h.me.Sleep()
}

// 学校
type School struct {
	Name     string
	students []*entity.People
}
func (s *School) Receive(student *entity.People) {
	student.StudentCard = vo.StudentCard{
		Id:     rand.Uint32(),
		Name:   student.IdentityCard.Name,
		School: s.Name,
	}
	s.students = append(s.students, student)
	fmt.Printf("%s Receive stduent %+v\n", s.Name, student.StudentCard)
}
// 执行School的业务逻辑
func (s *School) Run() {
	fmt.Printf("%s start class\n", s.Name)
	for _, student := range s.students {
		student.Study()
	}
	fmt.Println("students start to eating")
	for _, student := range s.students {
		student.Eat()
	}
	fmt.Println("students start to exam")
	for _, student := range s.students {
		student.Exam()
	}
	fmt.Printf("%s finish class\n", s.Name)
}

// 公司
type Company struct {
	Name    string
	workers []*entity.People
}
func (c *Company) Employ(worker *entity.People) {
	worker.WorkCard = vo.WorkCard{
		Id:      rand.Uint32(),
		Name:    worker.IdentityCard.Name,
		Company: c.Name,
	}
	c.workers = append(c.workers, worker)
	fmt.Printf("%s Employ worker %s\n", c.Name, worker.WorkCard.Name)
}
// 执行Company的业务逻辑
func (c *Company) Run() {
	fmt.Printf("%s start work\n", c.Name)
	for _, worker := range c.workers {
		worker.Work()
	}
	fmt.Println("worker start to eating")
	for _, worker := range c.workers {
		worker.Eat()
	}
	fmt.Println("worker get off work")
	for _, worker := range c.workers {
		worker.OffWork()
	}
	fmt.Printf("%s finish work\n", c.Name)
}

// 公园
type Park struct {
	Name     string
	enjoyers []*entity.People
}
func (p *Park) Welcome(enjoyer *entity.People) {
	fmt.Printf("%+v come to park %s\n", enjoyer.IdentityCard, p.Name)
	p.enjoyers = append(p.enjoyers, enjoyer)
}
// 执行Park的业务逻辑
func (p *Park) Run() {
	fmt.Printf("%s start to sell tickets\n", p.Name)
	for _, enjoyer := range p.enjoyers {
		enjoyer.BuyTicket()
	}
	fmt.Printf("%s start a show\n", p.Name)
	for _, enjoyer := range p.enjoyers {
		enjoyer.Enjoy()
	}
	fmt.Printf("show finish\n")
}
```

那么，根据上述方法建模出来的模型是这样的：

![](https://tva1.sinaimg.cn/large/008i3skNgy1gulxqehxg5j61ga0ts46v02.jpg)

模型的运行方法如下：

```go
paul := entity.NewPeople("Paul")
mit := aggregate.NewSchool("MIT")
google := aggregate.NewCompany("Google")
home := aggregate.NewHome()
summerPalace := aggregate.NewPark("Summer Palace")
// 上学
mit.Receive(paul)
mit.Run()
// 回家
home.ComeBack(paul)
home.Run()
// 工作
google.Employ(paul)
google.Run()
// 公园游玩
summerPalace.Welcome(paul)
summerPalace.Run()
```

## 贫血模型 VS 充血模型（工程派 VS 学院派）

上一节中，我们使用DDD的战术建模完成了该案例领域模型。模型的核心是`People`实体，它有`IdentityCard`、`StudentCard`等数据属性，也有`Eat()`、`Study()`、`Work()`等业务行为 ，非常符合现实世界中定义。这也是**学院派**所倡导的，同时拥有数据属性和业务行为的**充血模型**。

*然而，充血模型并非完美，它也有很多问题，比较典型的是这两个：*

**问题一：上帝类**

`People`这个实体包含了太多的职责，导致它变成了一个名副其实的上帝类。试想，这里还是裁剪了很多“人”所包含的属性和行为，如果要建模一个完整的模型，其属性和方法之多，无法想象。**上帝类违反了单一职责原则，会导致代码的可维护性变得极差**。

**问题二：模块间耦合**

`School`与`Company`本应该是相互独立的，`School`不必关注上班与否，`Company`也不必关注考试与否。但是现在因为它们都依赖了`People`这个实体，`School`可以调用与`Company`相关的`Work()`和`OffWork()`方法，反之亦然。这导致**模块间产生了不必要的耦合，违反了接口隔离原则**。

这些问题都是**工程派**不能接受的，从软件工程的角度，它们会使得代码难以维护。解决这类问题的方法，比较常见的是对实体进行拆分，比如将实体的行为建模成**领域服务**，像这样：

```go
type People struct {
	vo.IdentityCard
	vo.StudentCard
	vo.WorkCard
	vo.Account
}

type StudentService struct{}
func (s *StudentService) Study(p *entity.People) {
	fmt.Printf("Student %+v studying\n", p.StudentCard)
}
func (s *StudentService) Exam(p *entity.People) {
	fmt.Printf("Student %+v examing\n", p.StudentCard)
}

type WorkerService struct{}
func (w *WorkerService) Work(p *entity.People) {
	fmt.Printf("%+v working\n", p.WorkCard)
	p.Account.Balance++
}
func (w *WorkerService) OffWOrk(p *entity.People) {
	fmt.Printf("%+v getting off work\n", p.WorkCard)
}

// ...
```

![](https://tva1.sinaimg.cn/large/008i3skNgy1gv9h5miav4j617i0o8n2402.jpg)

这种建模方法，解决了上述两个问题，但也变成了所谓的**贫血模型**：`People`变成了一个纯粹的数据类，没有任何业务行为。在人的心理上，这样的模型并不能在建立起对现实世界的对应关系，不容易让人理解，因此被学院派所抵制。

到目前为止，贫血模型和充血模型都有各有优缺点，工程派和学院派谁都无法说服对方。接下来，轮到本文的主角出场了。

## DCI架构

**DCI**（Data，Context，Interactive）架构是一种面向对象的软件架构模式，在《[The DCI Architecture: A New Vision of Object-Oriented Programming](https://www.artima.com/articles/the-dci-architecture-a-new-vision-of-object-oriented-programming)》一文中被首次提出。与传统的面向对象相比，DCI能更好地对数据和行为之间的关系进行建模，从而更容易被人理解。

- **Data**，也即数据/领域对象，用来描述系统“是什么”，通常采用DDD中的战术建模来识别当前模型的领域对象，等同于DDD分层架构中的领域层
- **Context**，也即场景，可理解为是系统的Use Case，代表了系统的业务处理流程，等同于DDD分层架构中的应用层
- **Interactive**，也即交互，是DCI相对于传统面向对象的最大发展，它认为我们应该显式地对领域对象（**Object**）在每个业务场景（Context）中扮演（**Cast**）的角色（**Role**）进行建模。**Role代表了领域对象在业务场景中的业务行为（“做什么”），Role之间通过交互完成完整的义务流程**。

> 这种角色扮演的模型我们并不陌生，在现实的世界里也是随处可见，比如，一个演员可以在这部电影里扮演英雄的角色，也可以在另一部电影里扮演反派的角色。

DCI认为，对Role的建模应该是面向Context的，因为特定的业务行为只有在特定的业务场景下才会有意义。通过对Role的建模，我们就能够将领域对象的方法拆分出去，从而避免了上帝类的出现。最后，领域对象通过组合或继承的方式将Role集成起来，从而具备了扮演角色的能力。

![](https://tva1.sinaimg.cn/large/008i3skNgy1guqhelrwoej61ei0ss45e02.jpg)

DCI架构一方面通过角色扮演模型使得领域模型易于理解，另一方面通过“小类大对象”的手法避免了上帝类的问题，从而较好地解决了贫血模型和充血模型之争。另外，将领域对象的行为根据Role拆分之后，模块更加的高内聚、低耦合了。

## 使用DCI建模

回到前面的案例，使用DCI的建模思路，我们可以将“人”的几种行为按照不同的角色进行划分。吃完、睡觉、玩游戏，是作为**人类**角色的行为；学习、考试，是作为**学生**角色的行为；上班、下班，是作为**员工**角色的行为；购票、游玩，则是作为**游玩者**角色的行为。“人”在**家**这个场景中，充当的是人类的角色；在**学校**这个场景中，充当的是学生的角色；在**公司**这个场景中，充当的是员工的角色；在**公园**这个场景中，充当的是游玩者的角色。

![](https://tva1.sinaimg.cn/large/008i3skNgy1gus2bxbcbbj61220jq78c02.jpg)

> 需要注意的是，学生、员工、游玩者，这些角色都应该具备人类角色的行为，比如在学校里，学生也需要吃饭。

最后，根据DCI建模出来的模型，应该是这样的：

![](https://tva1.sinaimg.cn/large/008i3skNgy1gv9gzqe2xrj612q0om43502.jpg)

在DCI模型中，`People`不再是一个包含众多属性和方法的“上帝类”，这些属性和方法被拆分到多个role中实现，而`People`由这些role组合而成。

另外，`School`与`Company`也不再耦合，`School`只引用了`Student`，不能调用与`Company`相关的`Worker`的`Work()`和`OffWorker()`方法。

![](https://tva1.sinaimg.cn/large/008i3skNgy1gv9i5utlk1j619w0q0tep02.jpg)

## 代码实现DCI模型

DCI建模后的代码目录结构如下；

```yaml
- context: 场景
  - company.go
  - home.go
  - park.go
  - school.go
- object: 对象
  - people.go
- data: 数据
  - account.go
  - identity_card.go
  - student_card.go
  - work_card.go
- role: 角色
  - enjoyer.go
  - human.go
  - student.go
  - worker.go
```

从代码目录结构上看，DDD和DCI架构相差并不大，`aggregate`目录演变成了`context`目录；`vo`目录演变成了`data`目录；`entity`目录则演变成了`object`和`role`目录。

首先，我们实现基础角色`Human`，`Student`、`Worker`、`Enjoyer`都需要组合它：

```go
package role

// 人类角色
type Human struct {
	data.IdentityCard
	data.Account
}
func (h *Human) Eat() {
	fmt.Printf("%+v eating\n", h.IdentityCard)
	h.Account.Balance--
}
func (h *Human) Sleep() {
	fmt.Printf("%+v sleeping\n", h.IdentityCard)
}
func (h *Human) PlayGame() {
	fmt.Printf("%+v playing game\n", h.IdentityCard)
}
```

接着，我们再实现其他角色，需要注意的是，**`Student`、`Worker`、`Enjoyer`不能直接组合`Human`**，否则`People`对象将会有4个`Human`子对象，与模型不符：

```go
// 错误的实现
type Worker struct {
	Human
}
func (w *Worker) Work() {
	fmt.Printf("%+v working\n", w.WorkCard)
	w.Balance++
}
...
type People struct {
	Human
	Student
	Worker
	Enjoyer
}
func main() {
	people := People{}
  fmt.Printf("People: %+v", people)
}
// 结果输出, People中有4个Human：
// People: {Human:{} Student:{Human:{}} Worker:{Human:{}} Enjoyer:{Human:{}}}
```

为解决该问题，我们引入了`xxxTrait`接口：

```go
// 人类角色特征
type HumanTrait interface {
	CastHuman() *Human
}
// 学生角色特征
type StudentTrait interface {
	CastStudent() *Student
}
// 员工角色特征
type WorkerTrait interface {
	CastWorker() *Worker
}
// 游玩者角色特征
type EnjoyerTrait interface {
	CastEnjoyer() *Enjoyer
}
```

`Student`、`Worker`、`Enjoyer`组合`HumanTrait`，并通过`Compose(HumanTrait)`方法进行特征注入，只要在注入的时候保证`Human`是同一个，就可以解决该问题了。

```go
// 学生角色
type Student struct {
	// Student同时也是个普通人，因此组合了Human角色
	HumanTrait
	data.StudentCard
}
// 注入人类角色特征
func (s *Student) Compose(trait HumanTrait) {
	s.HumanTrait = trait
}
func (s *Student) Study() {
	fmt.Printf("Student %+v studying\n", s.StudentCard)
}
func (s *Student) Exam() {
	fmt.Printf("Student %+v examing\n", s.StudentCard)
}

// 员工角色
type Worker struct {
	// Worker同时也是个普通人，因此组合了Human角色
	HumanTrait
	data.WorkCard
}
// 注入人类角色特征
func (w *Worker) Compose(trait HumanTrait) {
	w.HumanTrait = trait
}
func (w *Worker) Work() {
	fmt.Printf("%+v working\n", w.WorkCard)
	w.CastHuman().Balance++
}
func (w *Worker) OffWork() {
	fmt.Printf("%+v getting off work\n", w.WorkCard)
}

// 游玩者角色
type Enjoyer struct {
	// Enjoyer同时也是个普通人，因此组合了Human角色
	HumanTrait
}
// 注入人类角色特征
func (e *Enjoyer) Compose(trait HumanTrait) {
	e.HumanTrait = trait
}
func (e *Enjoyer) BuyTicket() {
	fmt.Printf("%+v buying a ticket\n", e.CastHuman().IdentityCard)
	e.CastHuman().Balance--
}
func (e *Enjoyer) Enjoy() {
	fmt.Printf("%+v enjoying scenery\n", e.CastHuman().IdentityCard)
}

```

最后，实现`People`这一领域对象：

```go
package object

type People struct {
	// People对象扮演的角色
	role.Human
	role.Student
	role.Worker
	role.Enjoyer
}
// People实现了HumanTrait、StudentTrait、WorkerTrait、EnjoyerTrait等特征接口
func (p *People) CastHuman() *role.Human {
	return &p.Human
}
func (p *People) CastStudent() *role.Student {
	return &p.Student
}
func (p *People) CastWorker() *role.Worker {
	return &p.Worker
}
func (p *People) CastEnjoyer() *role.Enjoyer {
	return &p.Enjoyer
}
// People在初始化时，完成对角色特征的注入
func NewPeople(name string) *People {
  // 一些初始化的逻辑...
	people.Student.Compose(people)
	people.Worker.Compose(people)
	people.Enjoyer.Compose(people)
	return people
}
```

进行角色拆分之后，在实现`Home`、`School`、`Company`、`Park`等场景时，只需依赖相应的角色即可，不再需要依赖`People`这一领域对象：

```go
// 家
type Home struct {
	me *role.Human
}
func (h *Home) ComeBack(human *role.Human) {
	fmt.Printf("%+v come back home\n", human.IdentityCard)
	h.me = human
}
// 执行Home的业务逻辑
func (h *Home) Run() {
	h.me.Eat()
	h.me.PlayGame()
	h.me.Sleep()
}

// 学校
type School struct {
	Name     string
	students []*role.Student
}
func (s *School) Receive(student *role.Student) {
  // 初始化StduentCard逻辑 ...
	s.students = append(s.students, student)
	fmt.Printf("%s Receive stduent %+v\n", s.Name, student.StudentCard)
}
// 执行School的业务逻辑
func (s *School) Run() {
	fmt.Printf("%s start class\n", s.Name)
	for _, student := range s.students {
		student.Study()
	}
	fmt.Println("students start to eating")
	for _, student := range s.students {
		student.CastHuman().Eat()
	}
	fmt.Println("students start to exam")
	for _, student := range s.students {
		student.Exam()
	}
	fmt.Printf("%s finish class\n", s.Name)
}

// 公司
type Company struct {
	Name    string
	workers []*role.Worker
}
func (c *Company) Employ(worker *role.Worker) {
  // 初始化WorkCard逻辑 ...
  c.workers = append(c.workers, worker)
	fmt.Printf("%s Employ worker %s\n", c.Name, worker.WorkCard.Name)
}
// 执行Company的业务逻辑
func (c *Company) Run() {
	fmt.Printf("%s start work\n", c.Name)
	for _, worker := range c.workers {
		worker.Work()
	}
	fmt.Println("worker start to eating")
	for _, worker := range c.workers {
		worker.CastHuman().Eat()
	}
	fmt.Println("worker get off work")
	for _, worker := range c.workers {
		worker.OffWork()
	}
	fmt.Printf("%s finish work\n", c.Name)
}

// 公园
type Park struct {
	Name     string
	enjoyers []*role.Enjoyer
}
func (p *Park) Welcome(enjoyer *role.Enjoyer) {
	fmt.Printf("%+v come park %s\n", enjoyer.CastHuman().IdentityCard, p.Name)
	p.enjoyers = append(p.enjoyers, enjoyer)
}
// 执行Park的业务逻辑
func (p *Park) Run() {
	fmt.Printf("%s start to sell tickets\n", p.Name)
	for _, enjoyer := range p.enjoyers {
		enjoyer.BuyTicket()
	}
	fmt.Printf("%s start a show\n", p.Name)
	for _, enjoyer := range p.enjoyers {
		enjoyer.Enjoy()
	}
	fmt.Printf("show finish\n")
}
```

模型的运行方法如下：

```go
paul := object.NewPeople("Paul")
mit := context.NewSchool("MIT")
google := context.NewCompany("Google")
home := context.NewHome()
summerPalace := context.NewPark("Summer Palace")

// 上学
mit.Receive(paul.CastStudent())
mit.Run()
// 回家
home.ComeBack(paul.CastHuman())
home.Run()
// 工作
google.Employ(paul.CastWorker())
google.Run()
// 公园游玩
summerPalace.Welcome(paul.CastEnjoyer())
summerPalace.Run()
```

## 写在最后

从前文所描述的场景中，我们可以发现传统的DDD/面向对象设计方法在对行为进行建模方面存在着不足，进而导致了所谓的**贫血模型和充血模型之争**。DCI架构的出现很好的弥补了这一点，它通过引入**角色扮演**的思想，巧妙地解决了充血模型中上帝类和模块间耦合问题，而且不影响模型的正确性。当然，DCI架构也不是万能的，在行为较少的业务模型中，使用DCI来建模并不合适。

最后，将DCI架构总结成一句话就是：***领域对象（Object）在不同的场景（Context）中扮演（Cast）不同的角色（Role），角色之间通过交互（Interactive）来完成具体的业务逻辑***。

> 参考
>
> [The DCI Architecture: A New Vision of Object-Oriented Programming](https://www.artima.com/articles/the-dci-architecture-a-new-vision-of-object-oriented-programming), **Trygve Reenskaug**, **James O. Coplien**
>
> [软件设计的演变过程](https://www.jianshu.com/p/18d1d582f5c2), **\_张晓龙\_**
>
> [Implement Domain Object in Golang](https://www.jianshu.com/p/9fc3654b8165), **\_张晓龙\_**
>
> [DCI: 代码的可理解性](https://blog.csdn.net/chelsea/article/details/7093693), **chelsea**
>
> [DCI in C++](https://www.jianshu.com/p/bb9c35606d29), **MagicBowen**
>
> 更多文章请关注微信公众号：**元闰子的邀请**

