import { FormEvent, useState } from 'react'
import { toast } from 'react-toastify';

import { Form } from '../../components/Form'
import { Input } from '../../components/Input'

import { api } from '../../services/api'

export function CadastroLivros() {
  const [ nome, setNome ] = useState("")
  const [ isbn, setIsbn ] = useState("")

  function handleSubmit(e: FormEvent){
    e.preventDefault()
    api.post('livros', { nome, isbn })
    .then(_ => {
      toast.success("Sucesso");
      setNome("")
      setIsbn("")
    })
    .catch(_ => {
      toast.error("Erro");
    })
  }

  return (
    <Form
      onSubmit={handleSubmit}
    >
        <Input
          label="Nome"
          name="nome"
          value={nome}
          onChange={e => setNome(e.target.value)}
        />
        <Input
          name="isbn"
          label="ISBN"
          value={isbn}
          onChange={e => setIsbn(e.target.value)}
        />
    </Form>
  )
}
