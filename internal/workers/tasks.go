package workers

import (
	pb "uDES/api/proto"
	"uDES/internal/services"
)

type TaskResult struct {
	data interface{}
	err  error
}

type GetTask struct {
	InParams interface{}
	Result   chan TaskResult
}

func (r TaskResult) GetData() interface{} {
	return r.data
}

func (r TaskResult) GetErr() error {
	return r.err
}

func (t *GetTask) Execute() error {
	switch params := t.InParams.(type) {
	case *pb.GetEmployeeParams:
		employeesList, err := services.HandleGetEmployee(params)
		t.Result <- TaskResult{data: employeesList, err: err}
	case *pb.GetAbsencesParams:
		absencesList, err := services.HandleGetAbsences(params)
		t.Result <- TaskResult{data: absencesList, err: err}
	}
	return nil
}

func (t *GetTask) OnFailure(err error) {
	t.Result <- TaskResult{data: nil, err: err}
}
