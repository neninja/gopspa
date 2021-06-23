import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { api } from '../../services/api'

import { GenericTable } from '../../components/GenericTable'

interface Emprestimo {
  id: number
  aluno: Aluno
  livro: Livro
}

interface Aluno {
  id: number
  nome: string
}

interface Livro {
  id: number
  nome: string
  isbn: string
}

export function Emprestimos() {
  const [ emprestimos, setEmprestimos ] = useState<Emprestimo[]>([])
  const [ linhas, setLinhas ] = useState<string[][]>([])

  useEffect(() => {
    api.get('emprestimos')
    .then(response => setEmprestimos(response.data.emprestimos))
  }, [])

  useEffect(() => {
    const em = emprestimos.map(e => ([
      `${e.id}`, e.aluno.nome, e.livro.nome
    ]))
    setLinhas(em)
  }, [emprestimos])

  return (
    <>
      <Link to="/cadastro/emprestimos">Novo Emprestimo</Link>
      <GenericTable
        cabecalho={['Id', 'Aluno', 'Livro']}
        corpo={linhas}
      />
    </>
  )
}
