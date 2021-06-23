package getalluc

import "gopspa/pkg/gopspa"

type EmprestimosRepo interface {
	FindAll() ([]gopspa.Emprestimo, error)
}
