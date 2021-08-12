package mongo
import (
	"context"
	"sync"

	"awesomeProject3/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type mongoRepository struct {
	session *mgo.Session
}

var defaultDBName = "test"
var defaultCollection = "students"
var userCollection = "user"

var mutex sync.Mutex
var mongoRepo = &mongoRepository{}

func Init(host string) *mongoRepository {
	mutex.Lock()
	defer mutex.Unlock()

	if mongoRepo.session != nil {
		return mongoRepo
	}

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	mongoRepo.session = session
	return mongoRepo
}

// The implementation of Repository
//func UpdateFilter(m bson.M, filter map[string]string) error {
//	for k, v := range filter {
//		m[k] = interface{}(v)
//	}
//	return nil
//}
//
//func UpdateContextFilter(ctx context.Context, m bson.M) error {
//	// if context is admin, no need filter by tenantId.
//	md, ok := metadata.FromContext(ctx)
//	if !ok {
//		log.Error("get context failed")
//		return errors.New("get context failed")
//	}
//
//	isAdmin, _ := md[common.CTX_KEY_IS_ADMIN]
//	if isAdmin != common.CTX_VAL_TRUE {
//		tenantId, _ := md[common.CTX_REPRE_TENANT]
//		if tenantId == "" {
//			tenantId, ok = md[common.CTX_KEY_TENANT_ID]
//			if !ok {
//				log.Error("get tenantid failed")
//				return errors.New("get tenantid failed")
//			}
//		}
//
//		m["tenantId"] = tenantId
//	}
//
//	return nil
//}

func (repo *mongoRepository) CreateStudent(ctx context.Context, Student *model.StudentDetails) (*model.StudentDetails, error) {
	session := repo.session.Copy()
	defer session.Close()

	if Student.Id == "" {
		Student.Id = bson.NewObjectId()
	}

	err := session.DB(defaultDBName).C(defaultCollection).Insert(Student)
	if err != nil {
		return nil, err
	}
	return Student, nil
}

func (repo *mongoRepository) CreateUser(ctx context.Context, User *model.Credentials) (*model.Credentials, error) {
	session := repo.session.Copy()
	defer session.Close()

	if User.Id == "" {
		User.Id = bson.NewObjectId()
	}

	err := session.DB(defaultDBName).C(userCollection).Insert(User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *mongoRepository) DeleteStudent(ctx context.Context, id string) error {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": bson.ObjectIdHex(id)}
	//err := UpdateContextFilter(ctx, m)
	//if err != nil {
	//	return err
	//}

	return session.DB(defaultDBName).C(defaultCollection).Remove(m)
}

func (repo *mongoRepository) UpdateStudent(ctx context.Context,
	Student *model.StudentDetails) (*model.StudentDetails, error) {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": Student.Id}
	//err := UpdateContextFilter(ctx, m)
	//if err != nil {
	//	return nil, err
	//}

	err := session.DB(defaultDBName).C(defaultCollection).Update(m, Student)
	if err != nil {
		return nil, err
	}

	return Student, nil
}

func (repo *mongoRepository) GetStudent(ctx context.Context, id string) (*model.StudentDetails,
	error) {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": bson.ObjectIdHex(id)}
	//err := UpdateContextFilter(ctx, m)
	//if err != nil {
	//	return nil, err
	//}

	var Student = &model.StudentDetails{}
	collection := session.DB(defaultDBName).C(defaultCollection)
	err := collection.Find(m).One(Student)
	if err != nil {
		return nil, err
	}
	return Student, nil
}

func (repo *mongoRepository) GetUser(ctx context.Context, id string) (*model.Credentials,
	error) {
	session := repo.session.Copy()
	defer session.Close()

	m := bson.M{"_id": bson.ObjectIdHex(id)}
	//err := UpdateContextFilter(ctx, m)
	//if err != nil {
	//	return nil, err
	//}

	var User = &model.Credentials{}
	collection := session.DB(defaultDBName).C(userCollection)
	err := collection.Find(m).One(User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *mongoRepository) ListStudent(ctx context.Context) ([]*model.StudentDetails, error) {

	session := repo.session.Copy()
	defer session.Close()

	//if limit == 0 {
	//	limit = math.MinInt32
	//}
	var Students []*model.StudentDetails

	m := bson.M{}
	//UpdateFilter(m, query.(map[string]string))
	//err := UpdateContextFilter(ctx, m)
	//if err != nil {
	//	return nil, err
	//}
	//log.Infof("ListStudent, limit=%d, offset=%d, m=%+v\n", limit, offset, m)

	//err := session.DB(defaultDBName).C(defaultCollection).Find(m).Skip(offset).Limit(limit).All(&Students)
	err := session.DB(defaultDBName).C(defaultCollection).Find(m).All(&Students)

	if err != nil {
		return nil, err
	}

	return Students, nil
}

func (repo *mongoRepository) Close() {
	repo.session.Close()
}
