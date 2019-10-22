package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	TABLE_ENGINE = "INNODB"
)

/*基本用户表 */
type UserBase struct {
	Id                  int                                           //自增主键
	Nickname            string                 `orm:"size(50)"`       //昵称
	Avatar              string                 `orm:"size(255)"`      //头像
	Phone               string                 `orm:"size(20);null"`  //手机号
	Email               string                 `orm:"size(64);null"`  //邮箱
	Country             string                 `orm:"size(64);null"`  //国家
	City               string                 `orm:"size(64);null"`  //城市
	Province               string                 `orm:"size(64);null"`  //城市
	CreateTime          time.Time              `orm:"Type(datetime)"` //创建时间
	LoginTime           time.Time              `orm:"Type(datetime)"` //登录时间
	LastLoginTime       time.Time              `orm:"Type(datetime)"` //上一次登录时间
	AboutMe             string                 `orm:"size(500);null"` //关于我
	Interests           string                 `orm:"size(500);null"` //兴趣爱好
	EmotionalAttitude   string                 `orm:"size(500);null"` //感情观
	FavoriteTa          string                 `orm:"size(500);null"` //心仪的Ta
	UserPic             []*UserPic             `orm:"reverse(many)"`  //用户发布的照片组
	//UserAuths           *UserAuths             `orm:"rel(one)"`       //用户授权表
	UserFollow          []*UserFollow          `orm:"reverse(many)"`  //用户关注表
	UserFans            []*UserFans            `orm:"reverse(many)"`  //用户粉丝表
	UserDefriend        []*UserDefriend        `orm:"reverse(many)"`  //用户拉黑表
	FriendCircleMessage []*FriendCircleMessage `orm:"reverse(many)"`  //朋友圈
}

func (t *UserBase) TableEngine() string {
	return TABLE_ENGINE
}

/* 用户发布的照片组*/
type UserPic struct {
	Id         int
	Img        string `orm:"size(255)"`                  //图片地址
	Pos        int                                       //排序号
	Status     int8                                      //状态 0删除，1非删除
	CreateTime time.Time `orm:"Type(datetime)"`          //创建时间
	UserBase   *UserBase `orm:"column(user_id);rel(fk)"` //用户表，一对多
}

func (t *UserPic) TableEngine() string {
	return TABLE_ENGINE
}

/**
用户授权表
 */
type UserAuths struct {
	Id           int
	UserBase     *UserBase `orm:"column(user_id);rel(fk);reverse(one)"`                        //用户表,一对一关系
	IdentityType string    `orm:"size(20);nul;description(登录类型 (手机号/邮箱/用户名) 或第三方应用名称 (微信 , 微博等))"` //登录类型
	Identifier   string    `orm:"size(100);nul;description(标识 (手机号/邮箱/用户名或第三方应用的唯一标识,如openid))"`   //标识
	Credential   string    `orm:"size(100);nul;description(密码凭证 (站内的保存密码 , 站外的不保存或保存 token))"`
}

func (t *UserAuths) TableEngine() string {
	return TABLE_ENGINE
}

/**
用户关注信息表
 */
type UserFollow struct {
	Id         int
	UserBase   *UserBase `orm:"column(user_id);rel(fk)"` //用户表，一对多
	FollowedId int                                       //关注用户id
	Status     int8                                      //关注状态，1关注，0取消关注。默认1
	CreateTime time.Time `orm:"Type(datetime)"`          //关注时间
	UpadteTime time.Time `orm:"Type(datetime)"`          //更新关注时间
}
func (t *UserFollow) TableEngine() string {
	return TABLE_ENGINE
}

/**
用户粉丝信息表
 */
type UserFans struct {
	Id         int
	UserBase   *UserBase `orm:"column(user_id);rel(fk)"` //用户表，一对多
	FansId     int                                       //粉丝用户id
	Status     int8                                      //粉丝状态，1粉丝，0取消粉丝。默认1
	CreateTime time.Time `orm:"Type(datetime)"`          //粉丝时间
	UpadteTime time.Time `orm:"Type(datetime)"`          //更新粉丝时间
}
func (t *UserFans) TableEngine() string {
	return TABLE_ENGINE
}
/**
用户关注信息表D
 */
type UserDefriend struct {
	Id         int
	UserBase   *UserBase `orm:"column(user_id);rel(fk)"` //用户表，一对多
	DefriendId int                                       //拉黑用户id
	Status     int8                                      //拉黑状态，1拉黑，0取消拉黑。默认1
	CreateTime time.Time `orm:"Type(datetime)"`          //拉黑时间
	UpadteTime time.Time `orm:"Type(datetime)"`          //更新拉黑时间
}
func (t *UserDefriend) TableEngine() string {
	return TABLE_ENGINE
}
/**
朋友圈表
 */
type FriendCircleMessage struct {
	Id                   int
	UserBase             *UserBase `orm:"column(user_id);rel(fk)"`      //用户表，一对多
	Content              string    `orm:"size(500)"`                    //文字内容，支持emjoy
	Picture              string    `orm:"size(255)"`                    //图片内容。单图或多图
	Location             int                                            //位置
	CreateTime           time.Time               `orm:"Type(datetime)"` //创建时间
	FriendCircleTimeline []*FriendCircleTimeline `orm:"reverse(many)"`  //朋友圈时间线，一对多
	FriendCircleComment  []*FriendCircleComment  `orm:"reverse(many)"`  //朋友圈点赞数，一对多
}
func (t *FriendCircleMessage) TableEngine() string {
	return TABLE_ENGINE
}
/**
朋友圈时间线
 */
type FriendCircleTimeline struct {
	Id                  int64
	UserBase            *UserBase            `orm:"column(user_id);rel(fk)"` //用户表，一对多
	FriendCircleMessage *FriendCircleMessage `orm:"column(fcm_id);rel(fk)"`  //朋友圈id
	IsOwn               int8                                                 //是否是自己的
	CreateTime          time.Time `orm:"Type(datetime)"`                     //创建时间
}
func (t *FriendCircleTimeline) TableEngine() string {
	return TABLE_ENGINE
}
/**
朋友圈点赞评论表
 */
type FriendCircleComment struct {
	Id                  int64
	UserBase            *UserBase            `orm:"column(user_id);rel(fk)"` //用户表，一对多
	FriendCircleMessage *FriendCircleMessage `orm:"column(fcm_id);rel(fk)"`  //朋友圈id
	Content             string               `orm:"size(255)"`               //评论内容
	IsOwn               int8                                                 //是否是自己的
	CreateTime          time.Time `orm:"Type(datetime)"`                     //创建时间
	LikeCount           int                                                  //点赞数
}
func (t *FriendCircleComment) TableEngine() string {
	return TABLE_ENGINE
}
/**
推荐喜好设置表
 */
type RecommendedPreference struct {
	Id            int64
	UserBase      *UserBase `orm:"column(user_id);rel(fk)"` //用户表，一对多
	City          string    `orm:"size(50)"`                //城市，默认跟自己同城
	IsExtendScope int8      `orm:"default(0)"`              //是否扩大范围，及是否推荐城市之外，0.不扩大（默认），1.扩大
	Education     int8                                      //学历，1.大专，2.本科，3.硕士，4.博士，默认同自己的学历
	Age           int8                                      //年龄,默认自己的年龄上下浮动5岁
	Gender        int8                                      //性别，默认异性
}
func (t *RecommendedPreference) TableEngine() string {
	return TABLE_ENGINE
}


func init() {
	//初始化配置文件
	//iniconf, err := config.NewConfig("ini", "D:/myself_code/jie_ling/src/jie_ling_api/conf/config.conf")
	iniconf, err := config.NewConfig("ini", "C:/TNT/myself_code/jie_ling/src/jie_ling_api/conf/config.conf")
	if err != nil {
		beego.Error("iniconf error") //初始化配置文件出错
	}

	err = orm.RegisterDataBase(iniconf.String("db::aliasName"), iniconf.String("db::driverName"), iniconf.String("db::dataSource"), 30)
	if err != nil {
		beego.Error("RegisterDataBase error") //捕获数据库连接错误
	}
	orm.RegisterModelWithPrefix("jie_ling_", new(UserBase), new(UserPic), new(UserAuths), new(UserFollow), new(UserFans), new(UserDefriend),
		new(FriendCircleMessage), new(FriendCircleTimeline), new(FriendCircleComment), new(RecommendedPreference))

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		beego.Error(err)
	}
}
