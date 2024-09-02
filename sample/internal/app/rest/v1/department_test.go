package v1

import (
	"testing"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDepartmentRoute_GetDepartments(t *testing.T) {

	type (
		args struct {
			offset int
			limit  int
		}
	)

	testCases := []struct {
		desc     string
		args     args
		mockFunc func(args args, deptSvc *mocks.IDepartmentService)
		response string
	}{
		{
			desc: "success get departments with content",
			args: args{
				offset: 0,
				limit:  10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				depts := []domain.Department{
					{Id: 1, Name: "IT"},
					{Id: 2, Name: "HR"},
				}

				deptSvc.On("GetDepartments", args.offset, args.limit).Return(
					depts, nil)
			},
			response: `{"message":"success","data":{"departments":[{"id":1,"name":"IT"},{"id":2,"name":"HR"}]},"version":"-"}`,
		}, {
			desc: "success get departments with empty content",
			args: args{
				offset: 0,
				limit:  10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("GetDepartments", args.offset, args.limit).Return(
					[]domain.Department{}, nil)
			},
			response: `{"message":"success","data":{"departments":[]},"version":"-"}`,
		}, {
			desc: "success get departments with nil",
			args: args{
				offset: 0,
				limit:  10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("GetDepartments", args.offset, args.limit).Return(
					nil, nil)
			},
			response: `{"message":"success","data":{"departments":[]},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			c, rec := requestTestHelper[any]("GET", nil, "")
			deptSvc := new(mocks.IDepartmentService)

			tc.mockFunc(tc.args, deptSvc)

			route := NewDepartmentRoute(deptSvc)
			err := route.getDepartments(c)
			assert.NoError(t, err)

			assert.Equal(t, 200, rec.Code)
			assert.JSONEq(t, tc.response, rec.Body.String())

			deptSvc.AssertExpectations(t)
		})
	}

}
