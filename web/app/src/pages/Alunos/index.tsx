import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { api } from '../../services/api'

import { GenericTable } from '../../components/GenericTable'

interface Aluno {
  id: number
  nome: string
}

export function Alunos() {
  const [ alunos, setAlunos ] = useState<Aluno[]>([])
  const [ linhas, setLinhas ] = useState<string[][]>([])

  useEffect(() => {
    api.get('alunos')
    .then(response => setAlunos(response.data.alunos))
  }, [])

  useEffect(() => {
    const a = alunos.map(aluno => ([
      `${aluno.id}`, aluno.nome
    ]))
    setLinhas(a)
  }, [alunos])

  return (
    <>
      <Link to="/cadastro/alunos">Novo Aluno</Link>
      <GenericTable
        cabecalho={['Id', 'Nome']}
        corpo={linhas}
      />
    </>
  )
}
