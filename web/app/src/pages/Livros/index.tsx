import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { api } from '../../services/api'

import { GenericTable } from '../../components/GenericTable'

interface Livro {
  id: number
  nome: string
  isbn: string
}

export function Livros() {
  const [ livros, setLivros ] = useState<Livro[]>([])
  const [ linhas, setLinhas ] = useState<string[][]>([])

  useEffect(() => {
    api.get('livros')
    .then(response => setLivros(response.data.livros))
  }, [])

  useEffect(() => {
    const l = livros.map(livro => ([
      `${livro.id}`, livro.nome, livro.isbn
    ]))
    setLinhas(l)
  }, [livros])

  return (
    <>
      <Link to="/cadastro/livros">Novo Livro</Link>
      <GenericTable
        cabecalho={['Id', 'Nome', 'ISBN']}
        corpo={linhas}
      />
    </>
  )
}
