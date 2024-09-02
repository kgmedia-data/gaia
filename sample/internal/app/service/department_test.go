package service

import (
	"fmt"
	"testing"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/mocks"
)

func TestDepartmentService_GetDepartment(t *testing.T) {

	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc     string
		args     args
		wantErr  bool
		mockFunc func(args args, deptRepo *mocks.IDepartmentRepo)
	}{
		{
			desc:    "success get department",
			wantErr: false,
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				temp := domain.Department{
					Id:   args.id,
					Name: "IT",
				}
				deptRepo.On("GetDepartment", args.id).Return(temp, nil)
			},
		}, {
			desc:    "failed get department not found",
			wantErr: true,
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				deptRepo.On("GetDepartment", args.id).Return(domain.Department{},
					fmt.Errorf("department not found"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			deptRepo := new(mocks.IDepartmentRepo)
			tc.mockFunc(tc.args, deptRepo)

			svc := NewDepartmentService(deptRepo)
			_, err := svc.GetDepartment(tc.args.id)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetDepartment() error = %v, wantErr %v", err, tc.wantErr)
			}

			deptRepo.AssertExpectations(t)
		})
	}
}

func TestDepartmentService_GetDepartments(t *testing.T) {

	type (
		args struct {
			offset int
			limit  int
		}
	)

	testCases := []struct {
		desc     string
		args     args
		wantErr  bool
		mockFunc func(args args, deptRepo *mocks.IDepartmentRepo)
	}{
		{
			desc:    "success get departments",
			wantErr: false,
			args: args{
				offset: 0,
				limit:  10,
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				temp := []domain.Department{
					{
						Id:   1,
						Name: "IT",
					},
				}
				deptRepo.On("GetDepartments", args.offset, args.limit).Return(temp, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			deptRepo := new(mocks.IDepartmentRepo)
			tc.mockFunc(tc.args, deptRepo)

			svc := NewDepartmentService(deptRepo)
			_, err := svc.GetDepartments(tc.args.offset, tc.args.limit)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetDepartments() error = %v, wantErr %v", err, tc.wantErr)
			}

			deptRepo.AssertExpectations(t)
		})
	}
}

func TestDepartmentService_InsertDepartment(t *testing.T) {

	type (
		args struct {
			dept domain.Department
		}
	)

	testCases := []struct {
		desc     string
		args     args
		wantErr  bool
		mockFunc func(args args, deptRepo *mocks.IDepartmentRepo)
	}{
		{
			desc:    "success insert department",
			wantErr: false,
			args: args{
				dept: domain.Department{
					Name: "IT",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				temp := domain.Department{
					Id:   1,
					Name: args.dept.Name,
				}
				deptRepo.On("InsertDepartment", args.dept).Return(temp, nil)
			},
		}, {
			desc:    "failed insert department empty name",
			wantErr: true,
			args: args{
				dept: domain.Department{
					Name: "",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {},
		}, {
			desc:    "failed insert department duplicate name",
			wantErr: true,
			args: args{
				dept: domain.Department{
					Name: "IT",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				deptRepo.On("InsertDepartment", args.dept).Return(domain.Department{},
					fmt.Errorf("duplicate department"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			deptRepo := new(mocks.IDepartmentRepo)
			tc.mockFunc(tc.args, deptRepo)

			svc := NewDepartmentService(deptRepo)
			_, err := svc.InsertDepartment(tc.args.dept)

			if (err != nil) != tc.wantErr {
				t.Errorf("InsertDepartment() error = %v, wantErr %v", err, tc.wantErr)
			}

			deptRepo.AssertExpectations(t)
		})
	}
}

func TestDepartmentService_UpdateDepartment(t *testing.T) {
	type (
		args struct {
			id   int
			dept domain.Department
		}
	)

	testCases := []struct {
		desc     string
		args     args
		wantErr  bool
		mockFunc func(args args, deptRepo *mocks.IDepartmentRepo)
	}{
		{
			desc:    "success update department",
			wantErr: false,
			args: args{
				id: 1,
				dept: domain.Department{
					Id:   1,
					Name: "Information Technology",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				existing := domain.Department{
					Id:   1,
					Name: "IT",
				}
				deptRepo.On("GetDepartment", args.id).Return(existing, nil)
				deptRepo.On("UpdateDepartment", args.dept).Return(args.dept, nil)
			},
		}, {
			desc:    "failed update department empty name",
			wantErr: true,
			args: args{
				id: 1,
				dept: domain.Department{
					Id:   1,
					Name: "",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {},
		}, {
			desc:    "failed update department not found",
			wantErr: true,
			args: args{
				id: 1,
				dept: domain.Department{
					Id:   1,
					Name: "Information Technology",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				deptRepo.On("GetDepartment", args.id).Return(domain.Department{},
					fmt.Errorf("department not found"))
			},
		}, {
			desc:    "failed update department duplicate name",
			wantErr: true,
			args: args{
				id: 1,
				dept: domain.Department{
					Id:   1,
					Name: "Information Technology",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				existing := domain.Department{
					Id:   1,
					Name: "IT",
				}
				deptRepo.On("GetDepartment", args.id).Return(existing, nil)
				deptRepo.On("UpdateDepartment", args.dept).Return(domain.Department{},
					fmt.Errorf("duplicate department"))
			},
		}, {
			desc:    "failed update id is 0",
			wantErr: true,
			args: args{
				id: 0,
				dept: domain.Department{
					Id:   0,
					Name: "Information Technology",
				},
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			deptRepo := new(mocks.IDepartmentRepo)
			tc.mockFunc(tc.args, deptRepo)

			svc := NewDepartmentService(deptRepo)
			_, err := svc.UpdateDepartment(tc.args.dept)

			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateDepartment() error = %v, wantErr %v", err, tc.wantErr)
			}

			deptRepo.AssertExpectations(t)
		})
	}
}

func TestDepartmentService_DeleteDepartment(t *testing.T) {

	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc     string
		args     args
		wantErr  bool
		mockFunc func(args args, deptRepo *mocks.IDepartmentRepo)
	}{
		{
			desc:    "success delete department",
			wantErr: false,
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				existing := domain.Department{
					Id:   1,
					Name: "IT",
				}
				deptRepo.On("GetDepartment", args.id).Return(existing, nil)
				deptRepo.On("DeleteDepartment", args.id).Return(nil)
			},
		}, {
			desc:    "failed delete department not found",
			wantErr: true,
			args: args{
				id: 1,
			},
			mockFunc: func(args args, deptRepo *mocks.IDepartmentRepo) {
				deptRepo.On("GetDepartment", args.id).Return(domain.Department{},
					fmt.Errorf("department not found"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			deptRepo := new(mocks.IDepartmentRepo)
			tc.mockFunc(tc.args, deptRepo)

			svc := NewDepartmentService(deptRepo)
			err := svc.DeleteDepartment(tc.args.id)

			if (err != nil) != tc.wantErr {
				t.Errorf("DeleteDepartment() error = %v, wantErr %v", err, tc.wantErr)
			}

			deptRepo.AssertExpectations(t)
		})
	}
}
