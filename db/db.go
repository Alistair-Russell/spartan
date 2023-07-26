package db
import (
   "context"
   "errors"
   "github.com/go-redis/redis/v8"
   "reflect"
)
type Database struct {
   Client *redis.Client
}
var (
   ErrNil = errors.New("no matching record found in redis database")
   Ctx    = context.TODO()
)

func structToMap(issue *Issue) map[string]interface{} {
    values := make(map[string]interface{})
    issueValue := reflect.ValueOf(issue).Elem()

    for i := 0; i < issueValue.NumField(); i++ {
        if issueValue.Field(i).CanInterface() {
            values[issueValue.Type().Field(i).Name] = issueValue.Field(i).Interface()
        }
    }

    return values
}

func mapToIssue(issueMap map[string]string) *Issue {
    issue := &Issue{}
    issueValue := reflect.ValueOf(issue).Elem()

    for i := 0; i < issueValue.NumField(); i++ {
        fieldName := issueValue.Type().Field(i).Name
        if value, ok := issueMap[fieldName]; ok {
            field := issueValue.FieldByName(fieldName)
            if field.IsValid() && field.CanSet() {
                switch field.Kind() {
                case reflect.String:
                    field.SetString(value)
                // Add cases for other types as needed, e.g. int, float, etc.
                }
            }
        }
    }

    return issue
}

func NewDatabase(address string) (*Database, error) {
   client := redis.NewClient(&redis.Options{
      Addr: address,
      Password: "",
      DB: 0,
   })
   if err := client.Ping(Ctx).Err(); err != nil {
      return nil, err
   }
   return &Database{
      Client: client,
   }, nil
}

func (db *Database) CreateIssue(issue *Issue) string {
    // generate a unique id for the issue
    issueID := fmt.Sprintf("issue:%s", issue.ID)

    // convert the issue to a map
    issueMap := structToMap(issue)

    // save the issue to redis
    db.Client.HMSet(Ctx, issueID, issueMap)

    return issueID
}

func (db *Database) GetIssue(issueID string) *Issue {
    issueMap := db.Client.HGetAll(Ctx, issueID).Val()

    // convert the map to an issue
    issue := mapToIssue(issueMap)

    return issue
}

func (db *Database) UpdateIssue(issueID string, updatedFields map[string]interface{}) {
    db.Client.HMSet(Ctx, issueID, updatedFields)
}

func (db *Database) DeleteIssue(issueID string) {
    db.Client.Del(Ctx, issueID)
}

