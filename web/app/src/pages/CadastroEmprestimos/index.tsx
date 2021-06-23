import { FormEvent, useState, useEffect } from 'react'
import { toast } from 'react-toastify';

import { Form } from '../../components/Form'
import { SelectInput } from '../../components/SelectInput'

import { api } from '../../services/api'

interface Aluno {
  id: string
  nome: string
}

interface Livro {
  id: string
  nome: string
}

interface Dropdown {
  text: string
  value: string
}

export function CadastroEmprestimos() {
  const [ alunos, setAlunos ] = useState<Dropdown[]>([])
  const [ livros, setLivros ] = useState<Dropdown[]>([])

  const [ idAluno, setIdAluno ] = useState<number>(0)
  const [ idLivro, setIdLivro ] = useState<number>(0)

  function handleSubmit(e: FormEvent){
    e.preventDefault()
    api.post('emprestimos', { idAluno, idLivro })
    .then(_ => {
      toast.success("Sucesso");
      setIdAluno(0)
      setIdLivro(0)
    })
    .catch(_ => {
      toast.error("Erro");
    })
  }

  useEffect(() => {
    api.get('alunos')
    .then(r => {
      const aa = r.data.alunos.map((a: Aluno) => ({
        text: a.nome,
        value: `${a.id}`
      }))
      console.log(aa)
      setAlunos(aa)
    })

    api.get('livros')
    .then(r => {
      const ll = r.data.livros.map((l: Livro) => ({
        text: l.nome,
        value: `${l.id}`
      }))
      setLivros(ll)
    })
  }, [])

  return (
    <Form
      onSubmit={handleSubmit}
    >
      <SelectInput
        label="Aluno"
        options={alunos}
        onChange={e => setIdAluno(parseInt(e.target.value))}
      />
      <SelectInput
        label="Livro"
        options={livros}
      />
    </Form>
  )
}
