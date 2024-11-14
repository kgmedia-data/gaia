## How to use GCP log

GCP Log using logrus

add configuration at `config.yaml` file as below
```yaml
gcp_log:
  project_id: "<projectId>"
  log_name: "<logName>"
  labels:
    project: "gaia-kgdata-aiml"
    env: "development"
    dst: "dst"
```

add `logrus.error` on every method `error`, as below
```go
func (r DepartmentGormRepo) error(err error, method string, params ...interface{}) error {
    message := fmt.Errorf("DepartmentGormRepo.(%v)(%v) %w", method, params, err)
    logrus.Error(message)
    return message
}
```

if you want to add outside error, fatal or panic, you can use code like below
```go
logrus.WithFields(logrus.Fields{
	"gcp": true,
}).Info("GetDepartments Repo is run")
```

for example
```go
// inside method example
func (r *DepartmentGormRepo) GetDepartments(offset, limit int) ([]domain.Department, error) {
    logrus.WithFields(logrus.Fields{
        "gcp": true,
    }).Info("GetDepartments Repo is run")
	
    tx := r.GormDB.
    Where("is_deleted = ?", false).
    Offset(offset).
    Limit(limit).
    Order("id ASC")
    ...
}
```