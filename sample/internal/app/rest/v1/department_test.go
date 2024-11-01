package v1

import (
	"fmt"
	"testing"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDepartmentRoute_GetDepartment(t *testing.T) {

	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc             string
		args             args
		mockFunc         func(args args, deptSvc *mocks.IDepartmentService)
		expectedCode     int
		expectedResponse string
	}{
		{
			desc: "success get department",
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				dept := domain.Department{Id: 1, Name: "IT"}

				deptSvc.On("GetDepartment", args.id).Return(dept, nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"department":{"id":1,"name":"IT"}},"version":"-"}`,
		}, {
			desc: "failed get department",
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("GetDepartment", args.id).Return(domain.Department{}, fmt.Errorf("not found"))
			},
			expectedCode:     500,
			expectedResponse: `{"message":"not found","data":{"department":{}},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			c, rec := requestTestHelper[any]("GET", nil, "")
			c.SetPath("/department/:" + DEPT_ID)
			c.SetParamNames(DEPT_ID)
			c.SetParamValues(fmt.Sprintf("%d", tc.args.id))

			deptSvc := new(mocks.IDepartmentService)

			tc.mockFunc(tc.args, deptSvc)

			route := NewDepartmentRoute(deptSvc)
			err := route.getDepartment(c)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			deptSvc.AssertExpectations(t)
		})
	}
}

func TestDepartmentRoute_GetDepartments(t *testing.T) {

	type (
		args struct {
			page  int
			limit int
		}
	)

	testCases := []struct {
		desc             string
		args             args
		mockFunc         func(args args, deptSvc *mocks.IDepartmentService)
		expectedCode     int
		expectedResponse string
	}{
		{
			desc: "success get departments with content",
			args: args{
				page:  1,
				limit: 10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				depts := []domain.Department{
					{Id: 1, Name: "IT"},
					{Id: 2, Name: "HR"},
				}

				deptSvc.On("GetDepartments", 0, args.limit).Return(
					depts, nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"departments":[{"id":1,"name":"IT"},{"id":2,"name":"HR"}]},"version":"-"}`,
		}, {
			desc: "success get departments with empty content",
			args: args{
				page:  3,
				limit: 10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("GetDepartments", 20, args.limit).Return(
					[]domain.Department{}, nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"departments":[]},"version":"-"}`,
		}, {
			desc: "success get departments with nil",
			args: args{
				page:  1,
				limit: 10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("GetDepartments", 0, args.limit).Return(
					nil, nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"departments":[]},"version":"-"}`,
		}, {
			desc: "failed get departments",
			args: args{
				page:  1,
				limit: 10,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("GetDepartments", 0, args.limit).Return(
					nil, fmt.Errorf("internal server error"))
			},
			expectedCode:     500,
			expectedResponse: `{"message":"internal server error","data":{"departments":[]},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			path := fmt.Sprintf("/department?page=%d&limit=%d", tc.args.page, tc.args.limit)
			c, rec := requestTestHelper[any]("GET", nil, path)
			deptSvc := new(mocks.IDepartmentService)

			tc.mockFunc(tc.args, deptSvc)

			route := NewDepartmentRoute(deptSvc)
			err := route.getDepartments(c)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			deptSvc.AssertExpectations(t)
		})
	}
}

func TestDepartmentRoute_InsertDepartment(t *testing.T) {
	type (
		args struct {
			deptDto departmentRespDto
		}
	)

	testCases := []struct {
		desc             string
		args             args
		mockFunc         func(args args, deptSvc *mocks.IDepartmentService)
		expectedCode     int
		expectedResponse string
	}{
		{
			desc: "success insert department",
			args: args{
				deptDto: departmentRespDto{Name: "IT"},
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				dept := domain.Department{Id: 0, Name: args.deptDto.Name}
				result := domain.Department{Id: 1, Name: args.deptDto.Name}
				deptSvc.On("InsertDepartment", dept).Return(result, nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"department":{"id":1,"name":"IT"}},"version":"-"}`,
		}, {
			desc: "failed insert department",
			args: args{
				deptDto: departmentRespDto{Name: "IT"},
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				dept := domain.Department{Id: 0, Name: args.deptDto.Name}
				deptSvc.On("InsertDepartment", dept).Return(domain.Department{}, fmt.Errorf("internal server error"))
			},
			expectedCode:     500,
			expectedResponse: `{"message":"internal server error","data":{"department":{}},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			c, rec := requestTestHelper("POST", tc.args.deptDto, "/department")
			deptSvc := new(mocks.IDepartmentService)

			tc.mockFunc(tc.args, deptSvc)

			route := NewDepartmentRoute(deptSvc)
			err := route.insertDepartment(c)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			deptSvc.AssertExpectations(t)
		})
	}
}

func TestDepartmentRoute_UpdateDepartment(t *testing.T) {
	type (
		args struct {
			id      int
			deptDto departmentRespDto
		}
	)

	testCases := []struct {
		desc             string
		args             args
		mockFunc         func(args args, deptSvc *mocks.IDepartmentService)
		expectedCode     int
		expectedResponse string
	}{
		{
			desc: "success update department",
			args: args{
				id:      1,
				deptDto: departmentRespDto{Name: "IT"},
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				dept := domain.Department{Id: args.id, Name: args.deptDto.Name}
				result := domain.Department{Id: args.id, Name: args.deptDto.Name}
				deptSvc.On("UpdateDepartment", dept).Return(result, nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"department":{"id":1,"name":"IT"}},"version":"-"}`,
		}, {
			desc: "failed update department",
			args: args{
				id:      1,
				deptDto: departmentRespDto{Name: "IT"},
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				dept := domain.Department{Id: args.id, Name: args.deptDto.Name}
				deptSvc.On("UpdateDepartment", dept).Return(domain.Department{}, fmt.Errorf("not found"))
			},
			expectedCode:     500,
			expectedResponse: `{"message":"not found","data":{"department":{}},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			c, rec := requestTestHelper("PUT", tc.args.deptDto, "/department")
			c.SetPath("/department/:" + DEPT_ID)
			c.SetParamNames(DEPT_ID)
			c.SetParamValues(fmt.Sprintf("%d", tc.args.id))

			deptSvc := new(mocks.IDepartmentService)

			tc.mockFunc(tc.args, deptSvc)

			route := NewDepartmentRoute(deptSvc)
			err := route.updateDepartment(c)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			deptSvc.AssertExpectations(t)
		})
	}
}

func TestDepartmentRoute_DeleteDepartment(t *testing.T) {
	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc             string
		args             args
		mockFunc         func(args args, deptSvc *mocks.IDepartmentService)
		expectedCode     int
		expectedResponse string
	}{
		{
			desc: "success delete department",
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("DeleteDepartment", args.id).Return(nil)
			},
			expectedCode:     200,
			expectedResponse: `{"message":"success","data":{"department":{}},"version":"-"}`,
		}, {
			desc: "failed delete department",
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptSvc *mocks.IDepartmentService) {
				deptSvc.On("DeleteDepartment", args.id).Return(fmt.Errorf("not found"))
			},
			expectedCode:     500,
			expectedResponse: `{"message":"not found","data":{"department":{}},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			c, rec := requestTestHelper[any]("DELETE", nil, "")
			c.SetPath("/department/:" + DEPT_ID)
			c.SetParamNames(DEPT_ID)
			c.SetParamValues(fmt.Sprintf("%d", tc.args.id))

			deptSvc := new(mocks.IDepartmentService)

			tc.mockFunc(tc.args, deptSvc)

			route := NewDepartmentRoute(deptSvc)
			err := route.deleteDepartment(c)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			deptSvc.AssertExpectations(t)
		})
	}
}
